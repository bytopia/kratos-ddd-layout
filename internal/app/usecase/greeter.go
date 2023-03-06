package usecase

import (
	"context"
	"github.com/bytopia/kratos-ddd-template/internal/domain/entity"
	"github.com/bytopia/kratos-ddd-template/internal/domain/repo"
	"github.com/bytopia/kratos-ddd-template/internal/pkg/logging"

	v1 "github.com/bytopia/kratos-ddd-template/api/helloworld/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// GreeterUsecase is a Greeter usecase.
type GreeterUsecase struct {
	repo repo.GreeterRepo
	log  *logging.Logger
}

// NewGreeterUsecase new a Greeter usecase.
func NewGreeterUsecase(repo repo.GreeterRepo, lad logging.Adapter) *GreeterUsecase {
	return &GreeterUsecase{repo: repo, log: logging.NewLogger(lad)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *GreeterUsecase) CreateGreeter(ctx context.Context, g *entity.Greeter) (*entity.Greeter, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
