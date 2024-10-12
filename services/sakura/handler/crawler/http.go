package handler

import (
	"net/http"

	"github.com/Wexlersolk/bloodysakura/services/common/genproto/crawler"
	"github.com/Wexlersolk/bloodysakura/services/common/util"
	"github.com/Wexlersolk/bloodysakura/services/sakura/types"
)

type CrawlerHttpHandler struct {
	sakuraService types.SakuraService
}

func NewHttpCrawlerHandler(sakuraService types.SakuraService) *CrawlerHttpHandler {
	handler := &CrawlerHttpHandler{
		sakuraService: sakuraService,
	}

	return handler
}

func (h *CrawlerHttpHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /crawler", h.CreateCrawler)
}

func (h *CrawlerHttpHandler) CreateCrawler(w http.ResponseWriter, r *http.Request) {
	var req crawler.CreateCrawlerRequest
	err := util.ParseJSON(r, &req)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	crawlerInstance := &crawler.CrawlerData{
		CrawlerID:  42,
		VisitUrl:   "github.com",
		WantedText: "github.com",
		GeckoPort:  4444,
		GeckoPath:  "local",
	}

	err = h.sakuraService.CreateSakura(r.Context(), crawlerInstance)
	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	res := &crawler.CreateCrawlerResponse{VisitUrl: "success"}
	util.WriteJSON(w, http.StatusOK, res)
}
