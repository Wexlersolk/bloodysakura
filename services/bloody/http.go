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
			log.Fatalf("client error: %v", err)
		}

		res, err := c.GetCrawler(ctx, &crawler.GetCrawlerRequest{
			CrawlerID: 42,
		})
		if err != nil {
			log.Fatalf("client error: %v", err)
		}

		t := template.Must(template.New("orders").Parse(ordersTemplate))

		if err := t.Execute(w, res.GetCrawlers()); err != nil {
			log.Fatalf("template error: %v", err)
		}
	})

	log.Println("Starting server on", s.addr)
	return http.ListenAndServe(s.addr, router)
}

var ordersTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Kitchen Orders</title>
</head>
<body>
    <h1>Orders List</h1>
    <table border="1">
        <tr>
            <th>Order ID</th>
            <th>Customer ID</th>
            <th>Quantity</th>
        </tr>
        {{range .}}
        <tr>
            <td>{{.OrderID}}</td>
            <td>{{.CustomerID}}</td>
            <td>{{.Quantity}}</td>
        </tr>
        {{end}}
    </table>
</body>
</html>`
