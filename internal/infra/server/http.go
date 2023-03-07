package server

import (
	v1 "github.com/bytopia/kratos-ddd-template/api/helloworld/v1"
	"github.com/bytopia/kratos-ddd-template/internal/adapter/grpc"
	"github.com/bytopia/kratos-ddd-template/internal/infra/conf"
	"github.com/bytopia/kratos-ddd-template/pkg/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(bc *conf.Bootstrap, greeter *grpc.GreeterService, log logging.LoggerAdapter) *http.Server {
	c := bc.Server
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterGreeterHTTPServer(srv, greeter)
	return srv
}
