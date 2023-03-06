package infra

import (
	"github.com/bytopia/kratos-ddd-template/internal/infra/client"
	"github.com/bytopia/kratos-ddd-template/internal/infra/database"
	"github.com/bytopia/kratos-ddd-template/internal/infra/server"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(client.NewData, database.NewGreeterRepo, server.NewHTTPServer, server.NewGRPCServer)
