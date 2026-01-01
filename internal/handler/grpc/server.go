package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/Fiagram/account_service/internal/generated/grpc/account_service"
	"google.golang.org/grpc"
)

type Server interface {
	Start(ctx context.Context) error
}

type server struct {
	handler account_service.AccountServiceServer
}

func NewServer(
	handler account_service.AccountServiceServer,
) Server {
	return &server{
		handler: handler,
	}
}

func (s *server) Start(ctx context.Context) error {
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		return err
	}
	defer listener.Close()

	server := grpc.NewServer()
	account_service.RegisterAccountServiceServer(server, s.handler)
	fmt.Println("server listening at 0.0.0.0:8080")
	return server.Serve(listener)
}
