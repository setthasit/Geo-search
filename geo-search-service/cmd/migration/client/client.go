package client

import "geo-search/internal/connectors"

type DatabaseClient struct {
	Mongo         *connectors.MongoConnector
	Elasticsearch *connectors.ElasticSearchConnector
}

func WireDatabaseClient(
	mongo *connectors.MongoConnector,
	elasticsearch *connectors.ElasticSearchConnector,
) *DatabaseClient {
	return &DatabaseClient{
		Mongo:         mongo,
		Elasticsearch: elasticsearch,
	}
}
