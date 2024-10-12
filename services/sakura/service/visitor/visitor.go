package visitor

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

type VisitFunc func(io.Reader) error

type VisitRequest struct {
	Links      []string
	VisitFunc  VisitFunc
	wantedText string
}

type ShutdownMessage struct {
	URL string
}

func NewVisitRequest(links []string, wantedText string) VisitRequest {
	return VisitRequest{
		Links:      links,
		wantedText: wantedText,
		VisitFunc: func(r io.Reader) error {
			fmt.Println("==========================")
			fmt.Println("==========================")
			b, err := io.ReadAll(r)
			if err != nil {
				return err
			}
			fmt.Println("reading a website data")

			pageContent := string(b)

			if strings.Contains(pageContent, wantedText) {
				fmt.Printf("Wanted text '%s' found!\n", wantedText)
				return fmt.Errorf("wanted text found")
			}

			fmt.Println("==========================")
			fmt.Println("==========================")
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
		links, err := visitor.DoVisit(visitor.URL.String(), visitor.visitFn)
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

func (visitor *Visitor) ExtractLinks(body io.Reader) ([]string, error) {
	links := make([]string, 0)
	tokenizer := html.NewTokenizer(body)
	baseDomain := visitor.URL.Host

	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			return links, nil
		}

		if tokenType == html.StartTagToken {
			token := tokenizer.Token()
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

func (visitor *Visitor) DoVisit(link string, visit VisitFunc) ([]string, error) {
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

	links, err := visitor.ExtractLinks(r)
	if err != nil {
		return []string{}, err
	}

	if err := visit(w); err != nil {
		return links, err
	}

	return links, nil
}
