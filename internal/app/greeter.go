package app

import (
	"context"
	"github.com/go-kratos/kratos-layout/internal/domain"

	v1 "github.com/go-kratos/kratos-layout/proto/helloworld/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// GreeterRepo is a Greater repo.
type GreeterRepo interface {
	Save(context.Context, *domain.Greeter) (*domain.Greeter, error)
	Update(context.Context, *domain.Greeter) (*domain.Greeter, error)
	FindByID(context.Context, int64) (*domain.Greeter, error)
	ListByHello(context.Context, string) ([]*domain.Greeter, error)
	ListAll(context.Context) ([]*domain.Greeter, error)
}

// GreeterUsecase is a Greeter usecase.
type GreeterUsecase struct {
	repo GreeterRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewGreeterUsecase(repo GreeterRepo, logger log.Logger) *GreeterUsecase {
	return &GreeterUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *GreeterUsecase) CreateGreeter(ctx context.Context, g *domain.Greeter) (*domain.Greeter, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
