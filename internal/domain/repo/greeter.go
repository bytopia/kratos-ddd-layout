package repo

import (
	"context"
	"github.com/bytopia/kratos-ddd-layout/internal/domain/entity"
)

// GreeterRepo is a Greater repo.
type GreeterRepo interface {
	Save(context.Context, *entity.Greeter) (*entity.Greeter, error)
	Update(context.Context, *entity.Greeter) (*entity.Greeter, error)
	FindByID(context.Context, int64) (*entity.Greeter, error)
	ListByHello(context.Context, string) ([]*entity.Greeter, error)
	ListAll(context.Context) ([]*entity.Greeter, error)
}
