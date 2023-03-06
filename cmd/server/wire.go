//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/bytopia/kratos-ddd-template/internal/adapter"
	"github.com/bytopia/kratos-ddd-template/internal/app"
	"github.com/bytopia/kratos-ddd-template/internal/infra"
	"github.com/bytopia/kratos-ddd-template/internal/infra/conf"
	"github.com/bytopia/kratos-ddd-template/internal/pkg/logging"
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Bootstrap, logging.Adapter) (*kratos.App, func(), error) {
	panic(wire.Build(adapter.ProviderSet, infra.ProviderSet, app.ProviderSet, newApp))
}
