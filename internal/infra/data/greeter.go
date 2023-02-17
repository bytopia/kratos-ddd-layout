package data

import (
	"context"
	"github.com/bytopia/kratos-ddd-template/internal/app"
	"github.com/bytopia/kratos-ddd-template/internal/domain"

	"github.com/go-kratos/kratos/v2/log"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) app.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *greeterRepo) Save(ctx context.Context, g *domain.Greeter) (*domain.Greeter, error) {
	return g, nil
}

func (r *greeterRepo) Update(ctx context.Context, g *domain.Greeter) (*domain.Greeter, error) {
	return g, nil
}

func (r *greeterRepo) FindByID(context.Context, int64) (*domain.Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListByHello(context.Context, string) ([]*domain.Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListAll(context.Context) ([]*domain.Greeter, error) {
	return nil, nil
}
