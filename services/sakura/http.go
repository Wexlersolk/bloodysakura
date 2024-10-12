package main

import (
	"log"
	"net/http"

	handler "github.com/Wexlersolk/bloodysakura/services/sakura/handler/crawler"
	"github.com/Wexlersolk/bloodysakura/services/sakura/service"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()

	crawlerService := service.NewSakuraService()
	crawlerHandler := handler.NewHttpCrawlerHandler(crawlerService)
	crawlerHandler.RegisterRouter(router)

	log.Println("Starting server on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
