package di

import (
	"geo-search/internal/interactors"

	"github.com/google/wire"
)

var InteractorSet = wire.NewSet(
	interactors.WireEstateInteractor,
)
