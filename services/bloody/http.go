package main

import (
	"context"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/Wexlersolk/bloodysakura/services/common/genproto/crawler"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()

	conn := NewGRPCClient(":9000")
	defer conn.Close()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received request at root endpoint")

		c := crawler.NewCrawlerServiceClient(conn)
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*1000)
		defer cancel()

		req := &crawler.CreateCrawlerRequest{
			CrawlerID:  42,
			VisitUrl:   "https://fantasyaudiobook.com",
			WantedText: []string{"A Thousand Sons", "Pages", "Length"},
			GeckoPort:  4444,
			GeckoPath:  "/usr/local/bin/geckodriver",
		}

		_, err := c.CreateCrawler(ctx, req)
		if err != nil {
			log.Printf("client error: %v", err)
			http.Error(w, "Failed to create crawler", http.StatusInternalServerError)
			return
		}

		log.Println("Crawler created")

		res, err := c.GetCrawler(ctx, &crawler.GetCrawlerRequest{
			CrawlerID: 42,
		})
		if err != nil {
			log.Printf("client error: %v", err)
			http.Error(w, "Failed to get crawler", http.StatusInternalServerError)
			return
		}

		log.Println("Successfully received crawler response")

		// Load template from file
		t, err := template.ParseFiles("/home/wexlersolk/work/bloodysakura/services/bloody/crawler.html")
		if err != nil {
			log.Printf("template error: %v", err)
			http.Error(w, "Failed to load template", http.StatusInternalServerError)
			return
		}

		// Execute the template with the data
		if err := t.Execute(w, res.GetCrawlers()); err != nil {
			log.Printf("template error: %v", err)
			http.Error(w, "Failed to render template", http.StatusInternalServerError)
		}
	})

	log.Println("Starting server on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
