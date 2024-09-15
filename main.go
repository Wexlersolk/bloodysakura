package main

import (
	"github.com/anthdm/hollywood/actor"
	"log"
	"log/slog"
	"time"
)

type VisitRequest struct {
	links []string
}

type Orchestrator struct {
}

func NewOrchestrator() actor.Producer {
	return func() actor.Receiver {
		return &Orchestrator{}
	}
}
func (m *Orchestrator) handleVisitRequest(msg VisitRequest) error {
	for _, link := range msg.links {
		slog.Info("visiting url", "url", link)
	}
	return nil
}
func (m *Orchestrator) Receive(c *actor.Context) {
	switch msg := c.Message().(type) {
	case VisitRequest:
		m.handleVisitRequest(msg)
	case actor.Started:
		slog.Info("Manager started")
	case actor.Stopped:
	}
}
func main() {
	e, err := actor.NewEngine(actor.NewEngineConfig())
	if err != nil {
		log.Fatal(err)
	}
	e.Spawn(NewOrchestrator(), "Orchestrator")

	time.Sleep(time.Second * 10)
}
