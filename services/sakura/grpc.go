package main

import (
	"log"
	"net"

	"github.com/Wexlersolk/bloody/services/sakura/service"
	handler "github.com/Wexlersolk/services/sakura/handler/crawler"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// register our grpc services
	crawlerService := service.NewCrawlerService()
	handler.NewGrpcCrawlerService(grpcServer, crawlerService)

	log.Println("Starting gRPC server on", s.addr)

	return grpcServer.Serve(lis)
}
