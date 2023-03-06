package adapter

import (
	"github.com/bytopia/kratos-ddd-template/internal/adapter/grpc"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(grpc.NewGreeterService)
