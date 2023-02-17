package rpc

import (
	"context"
	"github.com/go-kratos/kratos-layout/internal/app"
	"github.com/go-kratos/kratos-layout/internal/domain"

	v1 "github.com/go-kratos/kratos-layout/contract/helloworld/v1"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc *app.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *app.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &domain.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}
