package service

import (
	"context"
	"kratosblog/api/kratosblog/server/v1"

	"kratosblog/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

// SayHello implements kratosblog.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, req *v1.HelloRequest) (resp *v1.HelloReply, err error) {
	return &v1.HelloReply{Message: "Hello " + req.GetName()}, nil
}
