package grpc

import (
	"context"
	"net"

	"github.com/Fiagram/account_service/internal/configs"
	"github.com/Fiagram/account_service/internal/generated/grpc/account_service"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Server interface {
	Start(ctx context.Context) error
}

type server struct {
	logger  *zap.Logger
	config  configs.Grpc
	handler account_service.AccountServiceServer
}

func NewServer(
	config configs.Grpc,
	handler account_service.AccountServiceServer,
	logger *zap.Logger,
) Server {
	return &server{
		config:  config,
		handler: handler,
		logger:  logger,
	}
}

func (s *server) Start(ctx context.Context) error {
	listener, err := net.Listen("tcp", s.config.Address+":"+s.config.Port)
	if err != nil {
		return err
	}
	defer listener.Close()

	server := grpc.NewServer()
	account_service.RegisterAccountServiceServer(server, s.handler)
	s.logger.With(zap.String("address", s.config.Address)).
		With(zap.String("port", s.config.Port)).
		Info("the grpc server listening")
	return server.Serve(listener)
}
