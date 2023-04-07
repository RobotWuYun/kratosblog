package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "kratosblog/api/kratosblog/server/v1"
	"kratosblog/internal/service"
)

type Services struct {
	greeter  *service.GreeterService
	category *service.CategoryService
}

func NewServices(
	greeter *service.GreeterService,
	category *service.CategoryService,
) *Services {
	return &Services{
		greeter:  greeter,
		category: category,
	}
}

func (s *Services) registerGRPC(srv *grpc.Server) {
	v1.RegisterGreeterServer(srv, s.greeter)
	v1.RegisterCategoryServer(srv, s.category)
}

func (s *Services) registerHTTP(srv *http.Server) {
	v1.RegisterGreeterHTTPServer(srv, s.greeter)
	v1.RegisterCategoryHTTPServer(srv, s.category)
}
