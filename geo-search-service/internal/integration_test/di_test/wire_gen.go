// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di_test

import (
	"geo-search/internal/config"
	"geo-search/internal/connectors"
	"geo-search/internal/datasources"
	"geo-search/internal/di"
	"geo-search/internal/integration_test/fixtures"
	"geo-search/internal/integration_test/repositories"
	"github.com/google/wire"
)

// Injectors from integration_test_di.go:

func InitializeFixtures() *fixtures.Fixtures {
	commandConfig := config.WireCommandConfig()
	configLoaderImpl := config.WireConfigLoader(commandConfig)
	mongoConfig := config.WireMongoConfig(configLoaderImpl)
	mongoConnector := connectors.WireMongoConnector(mongoConfig)
	elasticsearchConfig := config.WireElasticsearchConfig(configLoaderImpl)
	elasticSearchConnector := connectors.WireElasticsearchConnector(elasticsearchConfig)
	estateFixtures := fixtures.WireEstateFixtures(mongoConnector, elasticSearchConnector)
	fixturesFixtures := fixtures.WireFixtures(estateFixtures)
	return fixturesFixtures
}

func InitializeRepositorySuite() *repositories.RepositoryTestSuite {
	commandConfig := config.WireCommandConfig()
	configLoaderImpl := config.WireConfigLoader(commandConfig)
	elasticsearchConfig := config.WireElasticsearchConfig(configLoaderImpl)
	elasticSearchConnector := connectors.WireElasticsearchConnector(elasticsearchConfig)
	mongoConfig := config.WireMongoConfig(configLoaderImpl)
	mongoConnector := connectors.WireMongoConnector(mongoConfig)
	estateDatasource := datasources.WireEstateDatasource(elasticSearchConnector, mongoConnector)
	repositoryTestSuite := repositories.WireRepositoryTestSuite(estateDatasource)
	return repositoryTestSuite
}

// integration_test_di.go:

var IntegrationTestSet = wire.NewSet(repositories.WireRepositoryTestSuite, fixtures.WireFixtures, fixtures.WireEstateFixtures, di.ConfigSet, di.ConnectorSet, di.HttpControllerSet, di.InteractorSet, di.RepositorySet)