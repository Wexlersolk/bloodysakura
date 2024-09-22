package crawler

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strings"

	"github.com/anthdm/hollywood/actor"
	"golang.org/x/net/html"
)

type ShutdownMessage struct {
	URL string
}

type VisitFunc func(io.Reader) error

type VisitRequest struct {
	links      []string
	visitFunc  VisitFunc
	wantedText string
}

func NewVisitRequest(links []string, wantedText string) VisitRequest {
	return VisitRequest{
		links:      links,
		wantedText: wantedText,
		visitFunc: func(r io.Reader) error {
			fmt.Println("==========================")
			b, err := io.ReadAll(r)
			if err != nil {
				return err
			}
			pageContent := string(b)
			fmt.Println(pageContent)
			fmt.Println("==========================")

			if strings.Contains(pageContent, wantedText) {
				fmt.Printf("Wanted text '%s' found!\n", wantedText)
				return fmt.Errorf("wanted text found")
			}
			return nil
		},
	}
}

type Visitor struct {
	managerPID *actor.PID
	URL        *url.URL
	visitFn    VisitFunc
	wantedText string
}

func NewVisitor(url *url.URL, mpid *actor.PID, visitFn VisitFunc, wantedText string) actor.Producer {
	return func() actor.Receiver {
		return &Visitor{
			URL:        url,
			managerPID: mpid,
			visitFn:    visitFn,
			wantedText: wantedText,
		}
	}
}

func (visitor *Visitor) Receive(context *actor.Context) {
	switch context.Message().(type) {
	case actor.Started:
		slog.Info("visitor started", "url", visitor.URL)
		links, err := visitor.doVisit(visitor.URL.String(), visitor.visitFn)
		if err != nil {
			if err.Error() == "wanted text found" {
				slog.Info("wanted text found, sending shutdown signal", "url", visitor.URL.String())
				context.Send(visitor.managerPID, ShutdownMessage{URL: visitor.URL.String()})
			} else {
				slog.Error("visit error", "err", err)
			}
			return
		}
		context.Send(visitor.managerPID, NewVisitRequest(links, visitor.wantedText))
		context.Engine().Poison(context.PID())
	case actor.Stopped:
		slog.Info("visitor stopped", "url", visitor.URL)
	}
}

func (visitor *Visitor) extractLinks(body io.Reader) ([]string, error) {
	links := make([]string, 0)
	tokenizer := html.NewTokenizer(body)
	baseDomain := visitor.URL.Host // Get the base domain

	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			return links, nil
		}

		if tokenType == html.StartTagToken {
			token := tokenizer.Token()
			if token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						lurl, err := url.Parse(attr.Val)
						if err != nil {
							return links, err
						}
						actualLink := visitor.URL.ResolveReference(lurl)

						if actualLink.Host == baseDomain {
							links = append(links, actualLink.String())
						}
					}
				}
			}
		}
	}
}

func (visitor *Visitor) doVisit(link string, visit VisitFunc) ([]string, error) {
	baseURL, err := url.Parse(link)
	if err != nil {
		return []string{}, err
	}
	resp, err := http.Get(baseURL.String())
	if err != nil {
		return []string{}, err
	}

	w := &bytes.Buffer{}
	r := io.TeeReader(resp.Body, w)

	links, err := visitor.extractLinks(r)
	if err != nil {
		return []string{}, err
	}

	if err := visit(w); err != nil {
		return links, err
	}

	return links, nil
}

type Orchestrator struct {
	visited    map[string]bool
	visitors   map[*actor.PID]bool
	wantedText string
	baseDomain string // Add base domain field
}

func NewOrchestrator(wantedText, baseDomain string) actor.Producer {
	return func() actor.Receiver {
		return &Orchestrator{
			visitors:   make(map[*actor.PID]bool),
			visited:    make(map[string]bool),
			wantedText: wantedText,
			baseDomain: baseDomain, // Set base domain
		}
	}
}

func (orchestrator *Orchestrator) Receive(context *actor.Context) {
	switch msg := context.Message().(type) {
	case VisitRequest:
		orchestrator.handleVisitRequest(context, msg)
	case ShutdownMessage:
		slog.Info("wanted text found, shutting down orchestrator", "url", msg.URL)
		context.Engine().Poison(context.PID())
	case actor.Started:
		slog.Info("orchestrator started")
	case actor.Stopped:
		slog.Info("orchestrator stopped")
	}
}

func (orchestrator *Orchestrator) handleVisitRequest(context *actor.Context, msg VisitRequest) error {
	for _, link := range msg.links {
		parsedLink, err := url.Parse(link)
		if err != nil {
			return err
		}

		// Ensure we only visit links from the same base domain
		if parsedLink.Host == orchestrator.baseDomain {
			if _, ok := orchestrator.visited[link]; !ok {
				slog.Info("visiting url", "url", link)
				context.SpawnChild(NewVisitor(parsedLink, context.PID(), msg.visitFunc, orchestrator.wantedText), "visitor/"+link)
				orchestrator.visited[link] = true
			}
		}
	}
	return nil
}
