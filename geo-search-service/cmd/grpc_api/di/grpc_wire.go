//go:build wireinject
// +build wireinject

package di

import (
	"geo-search/cmd/grpc_api/server"
	"geo-search/internal/di"

	"github.com/google/wire"
)

func InitializeGrpc() *server.GrpcServer {
	wire.Build(grpcSet)
	return &server.GrpcServer{}
}

var grpcSet = wire.NewSet(
	server.WireGrpcServer,
	di.GrpcControllerSet,
	di.ConfigSet,
	di.ConnectorSet,
	di.InteractorSet,
	di.RepositorySet,
)
