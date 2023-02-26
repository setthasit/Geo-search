include .env

##########################################
#
# Preparation targets
#
##########################################
# install - For Prerequisites setup please see link below.
# https://grpc.io/docs/languages/go/quickstart/
install:
	brew install ariga/tap/atlas
	go install github.com/google/wire/cmd/wire@v0.5.0
	go install github.com/golang/mock/mockgen@v1.6.0
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	go install github.com/onsi/ginkgo/v2/ginkgo

##########################################
#
# Code generation targets
#
##########################################
gen.mock:
	go generate ./internal/...

gen.wire:
	wire ./cmd/...
	wire ./internal/integration_test/di

gen.proto:
	protoc -I=internal/transports/grpc/proto --go_out=internal/transports/grpc/pb --go-grpc_out=internal/transports/grpc/pb internal/transports/grpc/proto/*.proto

##########################################
#
# Build targets
#
##########################################
build.%: COMMAND_TARGET=$*
build.%:
	go build -o build/${COMMAND_TARGET} cmd/${COMMAND_TARGET}/main.go	

##########################################
#
# Development targets
#
##########################################
dev.%: COMMAND_TARGET=$*
dev.%:
	@go build -o build/${COMMAND_TARGET} cmd/${COMMAND_TARGET}/main.go
	@build/${COMMAND_TARGET}
# HTTP API
dev.http: dev.http_api
# gRPC API
dev.grpc: dev.grpc_api

##########################################
#
# Test targets
#
##########################################
test:
	ginkgo ./...

test.cov:
	ginkgo --cover ./...

test.integration: test.compose.up test.go.integration test.compose.down test.integration.clean

test.integration.clean:
	rm -rf certs/ittest
	mkdir certs/ittest

test.go.integration:
	IT_TEST=true go test -tags=integration ./internal/integration_test -v

test.compose.up:
	docker compose -f docker-compose.test.yaml --env-file=.env.test up -d

test.compose.down:
	docker compose -f docker-compose.test.yaml down -v
##########################################
#
# Migration targets
#
##########################################
migrate.mongo:
	@echo "Not implemented"

migrate.elasticsearch:
	go build -o build/migration cmd/migration/main.go
	build/migration migrate:elasticsearch

migrate.seed:
	go build -o build/migration cmd/migration/main.go
	build/migration seed

##########################################
#
# Docker compose targets
#
# *** Use for local development ***
#
##########################################
compose.up:
	docker compose up -d

compose.down:
	docker compose down -v

compose.certs:
	docker cp geo-search-es01-1:/usr/share/elasticsearch/config/certs/ca/ca.crt ./certs
