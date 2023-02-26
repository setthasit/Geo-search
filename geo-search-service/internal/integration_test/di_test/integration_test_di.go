// +build integration
//go:build wireinject
// +build wireinject

package di_test

import (
	"geo-search/internal/di"
	"geo-search/internal/integration_test/fixtures"
	"geo-search/internal/integration_test/repositories"

	"github.com/google/wire"
)

func InitializeFixtures() *fixtures.Fixtures {
	wire.Build(IntegrationTestSet)
	return &fixtures.Fixtures{}
}

func InitializeRepositorySuite() *repositories.RepositoryTestSuite {
	wire.Build(IntegrationTestSet)
	return &repositories.RepositoryTestSuite{}
}

var IntegrationTestSet = wire.NewSet(
	// Suite
	repositories.WireRepositoryTestSuite,

	// Fixtures sets
	fixtures.WireFixtures,
	fixtures.WireEstateFixtures,

	// Application sets
	di.ConfigSet,
	di.ConnectorSet,
	di.HttpControllerSet,
	di.InteractorSet,
	di.RepositorySet,
)
