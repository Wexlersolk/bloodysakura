package orchestrator

import (
	"bloodysakura/config"
	"bloodysakura/crawler/visitor"
	"log/slog"
	"net/url"

	"github.com/anthdm/hollywood/actor"
)

type Orchestrator struct {
	visited  map[string]bool
	visitors map[*actor.PID]bool
	config   *config.Config
}

func NewOrchestrator(config *config.Config) actor.Producer {
	return func() actor.Receiver {
		return &Orchestrator{
			visitors: make(map[*actor.PID]bool),
			visited:  make(map[string]bool),
			config:   config,
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
		orchestrator.HandleSearchBar(context)
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

		if parsedLink.Host == orchestrator.config.VisitUrl.Host {
			if _, ok := orchestrator.visited[link]; !ok {
				slog.Info("visiting url", "url", link)
				context.SpawnChild(visitor.NewVisitor(parsedLink, context.PID(), msg.VisitFunc, orchestrator.config.WantedText), "visitor/"+link)
				orchestrator.visited[link] = true
			}
		}
	}
	return nil
}
