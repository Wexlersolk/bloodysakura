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
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
		defer cancel()

		_, err := c.CreateCrawler(ctx, &crawler.CreateCrawlerRequest{
			VisitUrl:   "github.com",
			WantedText: "github.com",
			GeckoPort:  4444,
			GeckoPath:  "local",
		})
		if err != nil {
			log.Printf("client error: %v", err)
			http.Error(w, "Failed to create crawler", http.StatusInternalServerError)
			return
		}

		res, err := c.GetCrawler(ctx, &crawler.GetCrawlerRequest{
			CrawlerID: 42,
		})
		if err != nil {
			log.Printf("client error: %v", err)
			http.Error(w, "Failed to get crawler", http.StatusInternalServerError)
			return
		}

		log.Println("Successfully received crawler response")

		t := template.Must(template.New("orders").Parse(ordersTemplate))
		if err := t.Execute(w, res.GetCrawlers()); err != nil {
			log.Printf("template error: %v", err)
			http.Error(w, "Failed to render template", http.StatusInternalServerError)
		}
	})

	log.Println("Starting server on", s.addr)
	return http.ListenAndServe(s.addr, router)
}

var ordersTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Crawlers</title>
</head>
<body>
    <h1>Crawler Data</h1>
    <table border="1">
        <tr>
            <th>Crawler ID</th>
            <th>VisitUrl</th>
            <th>WantedText</th>
        </tr>
        {{range .}}
        <tr>
            <td>{{.CrawlerID}}</td>
            <td>{{.VisitUrl}}</td>
            <td>{{.WantedText}}</td>
        </tr>
        {{end}}
    </table>
</body>
</html>`
