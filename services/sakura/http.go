package main

import (
	"log"
	"net/http"

	"github.com/Wexlersolk/sakura/services/sakura/service"
	handler "github.com/Wexlersolk/services/sakura/handler/crawler"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()

	crawlerService := service.NewCrawlerService()
	crawlerHandler := handler.NewHttpCrawlerHandler(crawlerService)
	crawlerHandler.RegisterRouter(router)

	log.Println("Starting server on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
