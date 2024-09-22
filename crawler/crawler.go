package crawler

import (
	"fmt"
	"log/slog"

	"github.com/anthdm/hollywood/actor"
)

type VisitRequest struct {
	links []string
}

type Orchestrator struct{}

func NewOrchestrator() actor.Producer {
	return func() actor.Receiver {
		return &Orchestrator{}
	}
}

func (orchestrator *Orchestrator) handleVisitRequest(msg VisitRequest) error {
	for _, link := range msg.links {
		slog.Info("visiting url", "url", link)
		fmt.Println("fu")
	}
	return nil
}

func (orchestrator *Orchestrator) Receive(context *actor.Context) {
	switch msg := context.Message().(type) {
	case VisitRequest:
		orchestrator.handleVisitRequest(msg)
	case actor.Initialized:
		fmt.Println("actor initialized")
	case actor.Started:
		fmt.Println("actor started")
	case actor.Stopped:
		fmt.Println("actor stopped")
	}
}
