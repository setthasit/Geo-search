package di

import (
	"geo-search/internal/transports/http"

	"github.com/google/wire"
)

var HttpControllerSet = wire.NewSet(
	http.WireEstateController,
)
