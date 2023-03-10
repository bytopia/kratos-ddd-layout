package api

import (
	"github.com/bytopia/kratos-ddd-layout/internal/api/grpc"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(grpc.NewGreeterService)
