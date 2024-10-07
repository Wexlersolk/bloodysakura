package service

import (
	"context"
)

var crawlerDb = make([]*orders.Order, 0)

type CrawlerService struct {
	// store
}

func NewOrderService() *CrawlerService {
	return &CrawlerService{}
}

func (s *CrawlerService) CreateOrder(ctx context.Context, order *orders.Order) error {
	crawlerDb = append(crawlerDb, order)
	return nil
}

func (s *CrawlerService) GetOrders(ctx context.Context) []*orders.Order {
	return crawlerDb
}
