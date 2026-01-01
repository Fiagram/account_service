VERSION := $(shell cat VERSION)
COMMIT_HASH := $(shell git rev-parse HEAD)
PROJECT_NAME := account_service

all: generate build-all

.PHONY: init
init:
	go install github.com/rubenv/sql-migrate/...@v1.8.1
	go install github.com/go-delve/delve/cmd/dlv@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.11
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.6.0
	sudo apt install protobuf-compiler

.PHONY: generate
generate:
	protoc -I=. \
	--go_out=internal/generated \
	--go-grpc_out=internal/generated \
	api/account_service/*.proto

.PHONY: build-linux-amd64
build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build \
		-ldflags "-X main.version=$(VERSION) -X main.commitHash=$(COMMIT_HASH)" \
		-o build/$(PROJECT_NAME)_linux_amd64 cmd/$(PROJECT_NAME)/*.go

.PHONY: build-linux-arm64
build-linux-arm64:
	GOOS=linux GOARCH=arm64 go build \
		-ldflags "-X main.version=$(VERSION) -X main.commitHash=$(COMMIT_HASH)" \
		-o build/$(PROJECT_NAME)_linux_arm64 cmd/$(PROJECT_NAME)/*.go

.PHONY: build-macos-amd64
build-macos-amd64:
	GOOS=darwin GOARCH=amd64 go build \
		-ldflags "-X main.version=$(VERSION) -X main.commitHash=$(COMMIT_HASH)" \
		-o build/$(PROJECT_NAME)_macos_amd64 cmd/$(PROJECT_NAME)/*.go

.PHONY: build-macos-arm64
build-macos-arm64:
	GOOS=darwin GOARCH=arm64 go build \
		-ldflags "-X main.version=$(VERSION) -X main.commitHash=$(COMMIT_HASH)" \
		-o build/$(PROJECT_NAME)_macos_arm64 cmd/$(PROJECT_NAME)/*.go

.PHONY: build-windows-amd64
build-windows-amd64:
	GOOS=windows GOARCH=amd64 go build \
		-ldflags "-X main.version=$(VERSION) -X main.commitHash=$(COMMIT_HASH)" \
		-o build/$(PROJECT_NAME)_windows_amd64.exe cmd/$(PROJECT_NAME)/*.go

.PHONY: build-windows-arm64
build-windows-arm64:
	GOOS=windows GOARCH=amd64 go build \
		-ldflags "-X main.version=$(VERSION) -X main.commitHash=$(COMMIT_HASH)" \
		-o build/$(PROJECT_NAME)_windows_arm64.exe cmd/$(PROJECT_NAME)/*.go

.PHONY: build-all
build-all:
	make build-linux-amd64
	make build-linux-arm64
	make build-macos-amd64
	make build-macos-arm64
	make build-windows-amd64
	make build-windows-arm64

.PHONY: build
build:
	go build \
		-ldflags "-X main.version=$(VERSION) -X main.commitHash=$(COMMIT_HASH)" \
		-o build/$(PROJECT_NAME) \
		cmd/$(PROJECT_NAME)/*.go

.PHONY: clean
clean:
	rm -rf build/

.PHONY: docker-compose-up-dev
docker-compose-up-dev:
	docker-compose -f deployments/docker-compose.dev.yml up -d

.PHONY: docker-compose-down-dev
docker-compose-down-dev:
	docker-compose -f deployments/docker-compose.dev.yml down

.PHONY: run-standalone-server
run-standalone-server:
	go run cmd/$(PROJECT_NAME)/*.go
 

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: vendor
vendor:
	go mod vendor

.PHONY: lint
lint:
	golangci-lint run ./... 

.PHONY: migrate-up-dev
migrate-up-dev:
	sql-migrate up -config=internal/dataaccess/database/migrations/sql-migrate-config.yaml

.PHONY: migrate-down-dev
migrate-down-dev:
	sql-migrate down -config=internal/dataaccess/database/migrations/sql-migrate-config.yaml 

.PHONY: migrate-new
migrate-new:
	sql-migrate new -config=internal/dataaccess/database/migrations/sql-migrate-config.yaml "$(word 2,$(MAKECMDGOALS))" 
%:
	@true

.PHONY: migrate-status
migrate-status:
	sql-migrate status -config=internal/dataaccess/database/migrations/sql-migrate-config.yaml