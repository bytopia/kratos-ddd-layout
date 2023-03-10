package server

import (
	v1 "github.com/bytopia/kratos-ddd-layout/api/helloworld/v1"
	rpc "github.com/bytopia/kratos-ddd-layout/internal/api/grpc"
	"github.com/bytopia/kratos-ddd-layout/internal/infra/conf"
	"github.com/bytopia/kratos-ddd-layout/pkg/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(bc *conf.Bootstrap, greeter *rpc.GreeterService, log logging.LoggerAdapter) *grpc.Server {
	c := bc.Server
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterGreeterServer(srv, greeter)
	return srv
}
