package handler

import (
	"context"

	"github.com/Wexlersolk/bloodysakura/services/common/genproto/crawler"
	"github.com/Wexlersolk/bloodysakura/services/sakura/types"
	"google.golang.org/grpc"
)

type CrawlerGrpcHandler struct {
	sakuraService types.SakuraService
	crawler.UnimplementedCrawlerServiceServer
}

func NewGrpcCrawlerService(grpc *grpc.Server, crawlerService types.SakuraService) {
	gRPCHandler := &CrawlerGrpcHandler{
		sakuraService: crawlerService,
	}

	crawler.RegisterCrawlerServiceServer(grpc, gRPCHandler)
}

func (h *CrawlerGrpcHandler) GetCrawler(ctx context.Context, req *crawler.GetCrawlerRequest) (*crawler.GetCrawlerResponse, error) {
	o := h.sakuraService.GetSakura(ctx)
	res := &crawler.GetCrawlerResponse{
		Crawlers: o,
	}

	return res, nil
}

func (h *CrawlerGrpcHandler) CreateCrawler(ctx context.Context, req *crawler.CreateCrawlerRequest) (*crawler.CreateCrawlerResponse, error) {
	crawlerInstance := &crawler.CrawlerData{
		CrawlerID:  42,
		VisitUrl:   "https://www.kobo.com/",
		WantedText: "The Thouthand Sons",
		GeckoPort:  4444,
		GeckoPath:  "/usr/local/bin/geckodriver",
	}

	err := h.sakuraService.CreateSakura(ctx, crawlerInstance)
	if err != nil {
		return nil, err
	}

	res := &crawler.CreateCrawlerResponse{
		VisitUrl: "success",
	}

	return res, nil
}
