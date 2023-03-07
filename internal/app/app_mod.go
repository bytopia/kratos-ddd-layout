package app

import (
	"github.com/bytopia/kratos-ddd-layout/internal/app/usecase"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(usecase.NewGreeterUsecase)
