package server

import (
	"fmt"
	"log"
	itempb "monolith-divided-to-microservices/app/sdk/proto/item/v1"
	"monolith-divided-to-microservices/app/services/item/internal/config"
	"net"
	"strings"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	srv         *grpc.Server
	cfg         *config.Config
	itemHandler itempb.ItemServiceServer
}

func NewGRPCServer(cfg *config.Config, userHandler itempb.ItemServiceServer) *GRPCServer {
	return &GRPCServer{
		srv:         grpc.NewServer(),
		cfg:         cfg,
		itemHandler: userHandler,
	}
}

func (s *GRPCServer) Start() error {
	addr := s.cfg.GrpcPort
	if !strings.Contains(addr, ":") {
		addr = ":" + addr
	}
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	log.Printf("gRPC server started on %s", addr)

	itempb.RegisterItemServiceServer(s.srv, s.itemHandler)
	return s.srv.Serve(lis)
}

func (s *GRPCServer) Stop() {
	s.srv.GracefulStop()
}
