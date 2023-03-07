package database

import (
	"context"
	"github.com/bytopia/kratos-ddd-template/internal/domain/entity"
	"github.com/bytopia/kratos-ddd-template/internal/domain/repo"
	"github.com/bytopia/kratos-ddd-template/internal/infra/client"
	"github.com/bytopia/kratos-ddd-template/pkg/logging"
)

type greeterRepo struct {
	data *client.Data
	log  logging.LoggerAdapter
}

// NewGreeterRepo .
func NewGreeterRepo(data *client.Data, log logging.LoggerAdapter) repo.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log,
	}
}

func (r *greeterRepo) Save(ctx context.Context, g *entity.Greeter) (*entity.Greeter, error) {
	return g, nil
}

func (r *greeterRepo) Update(ctx context.Context, g *entity.Greeter) (*entity.Greeter, error) {
	return g, nil
}

func (r *greeterRepo) FindByID(context.Context, int64) (*entity.Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListByHello(context.Context, string) ([]*entity.Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListAll(context.Context) ([]*entity.Greeter, error) {
	return nil, nil
}
