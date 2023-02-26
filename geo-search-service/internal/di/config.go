package di

import (
	"geo-search/internal/config"

	"github.com/google/wire"
)

var ConfigSet = wire.NewSet(
	config.WireCommandConfig,
	config.WireConfigLoader,
	config.WireElasticsearchConfig,
	config.WireMongoConfig,
	config.WireAppConfig,

	wire.Bind(new(config.ConfigLoader), new(*config.ConfigLoaderImpl)),
)
