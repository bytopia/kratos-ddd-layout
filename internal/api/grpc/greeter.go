package grpc

import (
	"context"
	"github.com/bytopia/kratos-ddd-layout/internal/app/usecase"
	"github.com/bytopia/kratos-ddd-layout/internal/domain/entity"

	v1 "github.com/bytopia/kratos-ddd-layout/api/helloworld/v1"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc *usecase.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *usecase.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &entity.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}
