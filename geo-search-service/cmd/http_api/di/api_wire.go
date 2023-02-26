//go:build wireinject
// +build wireinject

package di

import (
	"geo-search/cmd/http_api/server"
	"geo-search/internal/di"

	"github.com/google/wire"
)

func InitializeAPI() *server.ApiServer {
	wire.Build(ApiSet)
	return &server.ApiServer{}
}

var ApiSet = wire.NewSet(
	server.ProvideApiRoutes,
	server.ProvideApiServer,
	di.ConfigSet,
	di.ConnectorSet,
	di.HttpControllerSet,
	di.InteractorSet,
	di.RepositorySet,
)
