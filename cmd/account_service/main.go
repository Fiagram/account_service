package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var (
	version    string
	commitHash string
)

func main() {
	var configFilePath string

	rootCommand := &cobra.Command{
		Use:     "account_service",
		Short:   "Starts the Account Service in standalone server mode.",
		Long:    "Account Service is a microservice for managing accounts belongs to Fiagram project.",
		Version: fmt.Sprintf("%s \ncommit: %s", version, commitHash),
		RunE: func(cmd *cobra.Command, _ []string) error {
			app, cleanup, err := InitStandaloneServer(configFilePath)
			if err != nil {
				return err
			}
			defer cleanup()
			return app.Start()
		},
	}

	rootCommand.Flags().StringVarP(&configFilePath,
		"config-file-path", "c", "",
		"Use the provided config file, otherwise the default embedded config applied.")

	if err := rootCommand.Execute(); err != nil {
		log.Panic(err)
	}
}
