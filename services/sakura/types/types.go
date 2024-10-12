package types

import (
	"context"

	"github.com/Wexlersolk/bloodysakura/services/common/genproto/crawler"
)

type SakuraService interface {
	CreateSakura(context.Context, *crawler.CrawlerData) error
	GetSakura(context.Context) []*crawler.CrawlerData
}
