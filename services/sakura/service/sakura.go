package service

import (
	"context"

	"github.com/Wexlersolk/bloodysakura/services/common/genproto/crawler"
)

var crawlerDb = make([]*crawler.CrawlerData, 0)

type CrawlerService struct {
}

func NewCrawlerService() *CrawlerService {
	return &CrawlerService{}
}

func (s *CrawlerService) CreateCrawler(ctx context.Context, crawler *crawler.CrawlerData) error {
	crawlerDb = append(crawlerDb, crawler)
	return nil
}

func (s *CrawlerService) GetCrawler(ctx context.Context) []*crawler.CrawlerData {
	return crawlerDb
}
