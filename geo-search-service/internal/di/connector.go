package di

import (
	"geo-search/internal/connectors"

	"github.com/google/wire"
)

var ConnectorSet = wire.NewSet(
	connectors.WireElasticsearchConnector,
	connectors.WireMongoConnector,
)
