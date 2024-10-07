package orchestrator

import (
	"bloodysakura/crawler/visitor"
	"bloodysakura/data"
	"log/slog"
	"net/url"

	"github.com/anthdm/hollywood/actor"
)

type Orchestrator struct {
	visited  map[string]bool
	visitors map[*actor.PID]bool
	data     *data.Data
}

func NewOrchestrator(config *data.Data) actor.Producer {
	return func() actor.Receiver {
		return &Orchestrator{
			visitors: make(map[*actor.PID]bool),
			visited:  make(map[string]bool),
			data:     config,
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
		slog.Info("info:", orchestrator.data.VisitUrl.String(), orchestrator.data.WantedText)
		visitRequest := visitor.NewVisitRequest([]string{orchestrator.data.VisitUrl.String()}, orchestrator.data.WantedText)
		context.Send(context.PID(), visitRequest)
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

		if parsedLink.Host == orchestrator.data.VisitUrl.Host {
			if _, ok := orchestrator.visited[link]; !ok {
				slog.Info("visiting url", "url", link)
				context.SpawnChild(visitor.NewVisitor(parsedLink, context.PID(), msg.VisitFunc, orchestrator.data.WantedText), "visitor/"+link)
				orchestrator.visited[link] = true
			}
		}
	}
	return nil
}
