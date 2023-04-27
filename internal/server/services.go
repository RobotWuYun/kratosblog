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
	article  *service.ArticleService
}

func NewServices(
	greeter *service.GreeterService,
	category *service.CategoryService,
	article *service.ArticleService,
) *Services {
	return &Services{
		greeter:  greeter,
		category: category,
		article:  article,
	}
}

func (s *Services) registerGRPC(srv *grpc.Server) {
	v1.RegisterGreeterServer(srv, s.greeter)
	v1.RegisterCategoryServer(srv, s.category)
	v1.RegisterArticleServer(srv, s.article)
}

func (s *Services) registerHTTP(srv *http.Server) {
	v1.RegisterGreeterHTTPServer(srv, s.greeter)
	v1.RegisterCategoryHTTPServer(srv, s.category)
	v1.RegisterArticleHTTPServer(srv, s.article)
}
