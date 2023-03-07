//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/bytopia/kratos-ddd-layout/internal/adapter"
	"github.com/bytopia/kratos-ddd-layout/internal/app"
	"github.com/bytopia/kratos-ddd-layout/internal/infra"
	"github.com/bytopia/kratos-ddd-layout/internal/infra/conf"
	"github.com/bytopia/kratos-ddd-layout/pkg/logging"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Bootstrap, log.Logger, logging.LoggerAdapter) (*kratos.App, func(), error) {
	panic(wire.Build(adapter.ProviderSet, infra.ProviderSet, app.ProviderSet, newApp))
}
