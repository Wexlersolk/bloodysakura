package orchestrator

import (
	"log/slog"
	"net/url"

	"github.com/Wexlersolk/bloodysakura/services/common/genproto/crawler"
	"github.com/Wexlersolk/bloodysakura/services/sakura/service/visitor"
	"github.com/anthdm/hollywood/actor"
)

type Orchestrator struct {
	visited     map[string]bool
	visitors    map[*actor.PID]bool
	crawlerData *crawler.CrawlerData
}

func NewOrchestrator(crawlerData *crawler.CrawlerData) actor.Producer {
	return func() actor.Receiver {
		return &Orchestrator{
			visitors:    make(map[*actor.PID]bool),
			visited:     make(map[string]bool),
			crawlerData: crawlerData,
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
		slog.Info("info:", orchestrator.crawlerData.VisitUrl, orchestrator.crawlerData.WantedText)
		visitRequest := visitor.NewVisitRequest([]string{orchestrator.crawlerData.VisitUrl}, orchestrator.crawlerData.WantedText)
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

		if parsedLink.Host == orchestrator.crawlerData.VisitUrl {
			if _, ok := orchestrator.visited[link]; !ok {
				slog.Info("visiting url", "url", link)
				context.SpawnChild(visitor.NewVisitor(parsedLink, context.PID(), msg.VisitFunc, orchestrator.crawlerData.WantedText), "visitor/"+link)
				orchestrator.visited[link] = true
			}
		}
	}
	return nil
}
