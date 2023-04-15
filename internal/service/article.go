package service

import (
	"context"
	"kratosblog/api/kratosblog/server/v1"
	"kratosblog/internal/util"

	"kratosblog/internal/biz"
)

// ArticleService is a article service.
type ArticleService struct {
	v1.UnimplementedArticleServer

	uc       *biz.ArticleUsecase
	category *biz.CategoryUsecase
}

// NewArticleService new a article service.
func NewArticleService(
	uc *biz.ArticleUsecase,
	category *biz.CategoryUsecase,
) *ArticleService {
	return &ArticleService{
		uc:       uc,
		category: category,
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

	offset, limit := util.GetPageQuery(req.GetPage(), req.GetPageSize())

	resp.Total, resp.List, err = s.uc.List(ctx, offset, limit)
	return
}
