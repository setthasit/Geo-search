package di

import (
	"geo-search/internal/transports/grpc"

	"github.com/google/wire"
)

var GrpcControllerSet = wire.NewSet(
	grpc.WireEstateController,
)
