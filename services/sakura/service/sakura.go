package service

import (
	"context"

	"github.com/Wexlersolk/bloodysakura/services/common/genproto/crawler"
)

var crawlerDb = make([]*crawler.CrawlerData, 0)

type CrawlerService struct {
	// store
}

func NewOrderService() *CrawlerService {
	return &CrawlerService{}
}

func (s *CrawlerService) CreateOrder(ctx context.Context, crawler *crawler.CrawlerData) error {
	crawlerDb = append(crawlerDb, crawler)
	return nil
}

func (s *CrawlerService) GetOrders(ctx context.Context) []*crawler.CrawlerData {
	return crawlerDb
}
