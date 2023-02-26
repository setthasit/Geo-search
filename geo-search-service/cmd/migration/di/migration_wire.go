//go:build wireinject
// +build wireinject

package di

import (
	"geo-search/cmd/migration/client"
	"geo-search/internal/di"

	"github.com/google/wire"
)

func InitializeMigration() *client.DatabaseClient {
	wire.Build(MigrationSet)
	return &client.DatabaseClient{}
}

var MigrationSet = wire.NewSet(
	client.WireDatabaseClient,
	di.ConfigSet,
	di.ConnectorSet,
)
