package api

import (
	"github.com/bytopia/kratos-ddd-template/internal/api/rpc"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(rpc.NewGreeterService)
