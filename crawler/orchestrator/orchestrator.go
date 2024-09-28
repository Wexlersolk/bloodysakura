package orchestrator

import (
	"bloodysakura/crawler/visitor"
	"log/slog"
	"net/url"

	"github.com/anthdm/hollywood/actor"
)

type Orchestrator struct {
	visited    map[string]bool
	visitors   map[*actor.PID]bool
	wantedText string
	baseDomain string
}

func NewOrchestrator(wantedText, baseDomain string) actor.Producer {
	return func() actor.Receiver {
		return &Orchestrator{
			visitors:   make(map[*actor.PID]bool),
			visited:    make(map[string]bool),
			wantedText: wantedText,
			baseDomain: baseDomain,
		}
	}
}

func (orchestrator *Orchestrator) Receive(context *actor.Context) {
	switch msg := context.Message().(type) {
	case visitor.VisitRequest:
		orchestrator.HandleVisitRequest(context, msg)
	case visitor.ShutdownMessage:
		slog.Info("wanted text found, shutting down orchestrator", "url", msg.URL)
		context.Engine().Poison(context.PID())
	case actor.Started:
		slog.Info("orchestrator started")
		orchestrator.handleSearchBar(context)
	case actor.Stopped:
		slog.Info("orchestrator stopped")
	}
}

func (orchestrator *Orchestrator) HandleVisitRequest(context *actor.Context, msg visitor.VisitRequest) error {
	for _, link := range msg.Links {
		parsedLink, err := url.Parse(link)
		if err != nil {
			return err
		}

		if parsedLink.Host == orchestrator.baseDomain {
			if _, ok := orchestrator.visited[link]; !ok {
				slog.Info("visiting url", "url", link)
				context.SpawnChild(visitor.NewVisitor(parsedLink, context.PID(), msg.VisitFunc, orchestrator.wantedText), "visitor/"+link)
				orchestrator.visited[link] = true
			}
		}
	}
	return nil
}

func (orchestrator *Orchestrator) handleSearchBar(context *actor.Context) error {
	return nil
}
