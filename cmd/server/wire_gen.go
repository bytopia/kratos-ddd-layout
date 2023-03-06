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
	"github.com/bytopia/kratos-ddd-template/internal/pkg/logging"
	"github.com/go-kratos/kratos/v2"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(bootstrap *conf.Bootstrap, adapter logging.Adapter) (*kratos.App, func(), error) {
	data, cleanup, err := client.NewData(adapter)
	if err != nil {
		return nil, nil, err
	}
	greeterRepo := database.NewGreeterRepo(data, adapter)
	greeterUsecase := usecase.NewGreeterUsecase(greeterRepo, adapter)
	greeterService := grpc.NewGreeterService(greeterUsecase)
	grpcServer := server.NewGRPCServer(bootstrap, greeterService, adapter)
	httpServer := server.NewHTTPServer(bootstrap, greeterService, adapter)
	app := newApp(adapter, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
