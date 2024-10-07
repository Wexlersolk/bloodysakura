package types

import (
	"context"

	"github.com/Wexlersolk/bloodysakura/services/common/genproto/crawler"
)

type CrawlerService interface {
	CreateCrawler(context.Context, *crawler.CrawlerData) error
	GetCrawler(context.Context) []*crawler.CrawlerData
}
