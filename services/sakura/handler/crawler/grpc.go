package handler

import (
	"context"

	"github.com/Wexlersolk/bloodysakura/services/common/genproto/crawler"
	"github.com/Wexlersolk/bloodysakura/services/sakura/types"
	"google.golang.org/grpc"
)

type CrawlerGrpcHandler struct {
	crawlerService types.CrawlerService
	crawler.UnimplementedCrawlerServiceServer
}

func NewGrpcCrawlerService(grpc *grpc.Server, crawlerService types.CrawlerService) {
	gRPCHandler := &CrawlerGrpcHandler{
		crawlerService: crawlerService,
	}

	crawler.RegisterCrawlerServiceServer(grpc, gRPCHandler)
}

func (h *CrawlerGrpcHandler) GetCrowler(ctx context.Context, req *crawler.GetCrawlerRequest) (*crawler.GetCrawlerResponse, error) {
	o := h.crawlerService.GetCrawler(ctx)
	res := &crawler.GetCrawlerResponse{
		Crawlers: o,
	}

	return res, nil
}

func (h *CrawlerGrpcHandler) CreateCrawler(ctx context.Context, req *crawler.CreateCrawlerRequest) (*crawler.CreateCrawlerResponse, error) {
	crawler := &crawler.CrawlerData{
		CrawlerID:  42,
		VisitUrl:   "github.com",
		WantedText: "github.com",
		GeckoPort:  4444,
		GeckoPath:  "local",
	}

	err := h.crawlerService.CreateCrawler(ctx, crawler)
	if err != nil {
		return nil, err
	}

	res := &crawler.CreateCrawlerResponse{
		VisitUrl: "success",
	}

	return res, nil
}
