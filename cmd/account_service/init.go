package main

import (
	"github.com/Fiagram/account_service/internal/app"
	"github.com/Fiagram/account_service/internal/configs"
	"github.com/Fiagram/account_service/internal/dataaccess/database"
	"github.com/Fiagram/account_service/internal/handler/grpc"
	"github.com/Fiagram/account_service/internal/logic"
	"github.com/Fiagram/account_service/internal/utils"
)

func InitStandaloneServer(configFilePath string) (app.StandaloneServer, func(), error) {
	config, err := configs.NewConfig(configFilePath)
	if err != nil {
		return nil, nil, err
	}

	logger, loggerCleanup, err := utils.InitializeLogger(config.Log)
	if err != nil {
		return nil, nil, err
	}

	db, dbCleanup, err := database.InitAndMigrateUpDatabase(config.Database, logger)
	if err != nil {
		loggerCleanup()
		return nil, nil, err
	}

	hashLogic := logic.NewHash(config.Auth.Hash)
	accountLogic := logic.NewAccount(hashLogic)

	accountHandler := grpc.NewHandler(accountLogic)
	grpcServer := grpc.NewServer(accountHandler)

	standaloneServer := app.NewStandaloneServer(grpcServer, logger)

	return standaloneServer,
		func() {
			dbCleanup()
			loggerCleanup()
		}, nil
}
