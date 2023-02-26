package di

import (
	"geo-search/internal/datasources"
	"geo-search/internal/repositories"

	"github.com/google/wire"
)

var RepositorySet = wire.NewSet(
	datasources.WireEstateDatasource,
	wire.Bind(new(repositories.EstateRepository), new(*datasources.EstateDatasource)),
)
