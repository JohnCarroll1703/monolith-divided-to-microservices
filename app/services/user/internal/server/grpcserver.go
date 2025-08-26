package server

import (
	"fmt"
	"log"
	userpb "monolith-divided-to-microservices/app/sdk/proto/user/v1"
	"monolith-divided-to-microservices/app/services/user/internal/config"
	"net"
	"strings"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	srv         *grpc.Server
	cfg         *config.Config
	userHandler userpb.UserServiceServer
}

func NewGRPCServer(cfg *config.Config, userHandler userpb.UserServiceServer) *GRPCServer {
	return &GRPCServer{
		srv:         grpc.NewServer(),
		cfg:         cfg,
		userHandler: userHandler,
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

	userpb.RegisterUserServiceServer(s.srv, s.userHandler)
	return s.srv.Serve(lis)
}

func (s *GRPCServer) Stop() {
	s.srv.GracefulStop()
}
