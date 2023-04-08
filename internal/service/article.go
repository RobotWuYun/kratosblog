package service

import (
	"context"
	"kratosblog/api/kratosblog/server/v1"

	"kratosblog/internal/biz"
)

// ArticleService is a article service.
type ArticleService struct {
	v1.UnimplementedArticleServer

	uc *biz.ArticleUsecase
}

// NewArticleService new a article service.
func NewArticleService(
	uc *biz.ArticleUsecase,
) *ArticleService {
	return &ArticleService{
		uc: uc,
	}
}

func (s *ArticleService) CategoryList(ctx context.Context, req *v1.CategoryListRequest) (resp *v1.CategoryListReply, err error) {
	resp = &v1.CategoryListReply{}
	return
}
func (s *ArticleService) TagList(ctx context.Context, req *v1.TagListRequest) (resp *v1.TagListReply, err error) {
	resp = &v1.TagListReply{}
	return
}
func (s *ArticleService) List(ctx context.Context, req *v1.ListRequest) (resp *v1.ListReply, err error) {
	resp = &v1.ListReply{}
	return
}
