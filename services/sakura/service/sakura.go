package service

import (
	"context"
	"log"
	"time"

	"github.com/Wexlersolk/bloodysakura/services/common/genproto/crawler"
	"github.com/Wexlersolk/bloodysakura/services/sakura/service/orchestrator"
	"github.com/anthdm/hollywood/actor"
)

var crawlerDb = make([]*crawler.CrawlerData, 0)

type SakuraService struct {
}

func NewSakuraService() *SakuraService {
	return &SakuraService{}
}

func (s *SakuraService) CreateSakura(ctx context.Context, crawler *crawler.CrawlerData) error {
	crawlerDb = append(crawlerDb, crawler)

	engine, err := actor.NewEngine(actor.NewEngineConfig())
	if err != nil {
		log.Fatal(err)
	}

	engine.Spawn(orchestrator.NewOrchestrator(crawler), "manager")

	time.Sleep(time.Millisecond * 200)
	return nil
}

func (s *SakuraService) GetSakura(ctx context.Context) []*crawler.CrawlerData {
	return crawlerDb
}
