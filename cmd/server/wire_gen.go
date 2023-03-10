// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/bytopia/kratos-ddd-template/internal/adapter/grpc"
	"github.com/bytopia/kratos-ddd-template/internal/app/usecase"
	"github.com/bytopia/kratos-ddd-template/internal/infra/client"
	"github.com/bytopia/kratos-ddd-template/internal/infra/conf"
	"github.com/bytopia/kratos-ddd-template/internal/infra/database"
	"github.com/bytopia/kratos-ddd-template/internal/infra/server"
	"github.com/bytopia/kratos-ddd-template/pkg/logging"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(bootstrap *conf.Bootstrap, logger log.Logger, loggerAdapter logging.LoggerAdapter) (*kratos.App, func(), error) {
	data, cleanup, err := client.NewData(loggerAdapter)
	if err != nil {
		return nil, nil, err
	}
	greeterRepo := database.NewGreeterRepo(data, loggerAdapter)
	greeterUsecase := usecase.NewGreeterUsecase(greeterRepo, loggerAdapter)
	greeterService := grpc.NewGreeterService(greeterUsecase)
	grpcServer := server.NewGRPCServer(bootstrap, greeterService, loggerAdapter)
	httpServer := server.NewHTTPServer(bootstrap, greeterService, loggerAdapter)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
