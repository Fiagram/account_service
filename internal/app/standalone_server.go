package app

import (
	"context"
	"syscall"

	"github.com/Fiagram/account_service/internal/handler/grpc"
	"github.com/Fiagram/account_service/internal/utils"

	"go.uber.org/zap"
)

type StandaloneServer interface {
	Start() error
}

type standaloneServer struct {
	grpcServer grpc.Server
	logger     *zap.Logger
}

func NewStandaloneServer(
	grpcServer grpc.Server,
	logger *zap.Logger,
) StandaloneServer {
	return &standaloneServer{
		grpcServer: grpcServer,
		logger:     logger,
	}
}

func (s standaloneServer) Start() error {
	go func() {
		err := s.grpcServer.Start(context.Background())
		s.logger.With(zap.Error(err)).Info("grpc server stopped")
	}()

	utils.BlockUntilSignal(syscall.SIGINT, syscall.SIGTERM)
	return nil
}
