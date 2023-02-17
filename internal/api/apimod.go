package api

import (
	"github.com/go-kratos/kratos-layout/internal/api/rpc"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(rpc.NewGreeterService)
