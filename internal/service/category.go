package service

import (
	"context"
	"kratosblog/api/kratosblog/server/v1"

	"kratosblog/internal/biz"
)

// CategoryService is a category service.
type CategoryService struct {
	v1.UnimplementedCategoryServer

	uc   *biz.CategoryUsecase
	auth *biz.AuthUsecase
}

// NewCategoryService new a category service.
func NewCategoryService(
	uc *biz.CategoryUsecase,
	auth *biz.AuthUsecase,
) *CategoryService {
	return &CategoryService{
		uc:   uc,
		auth: auth,
	}
}

func (s *CategoryService) ListCategory(ctx context.Context, req *v1.ListCategoryRequest) (resp *v1.ListCategoryReply, err error) {
	resp = &v1.ListCategoryReply{}

	var getAll bool
	getAll, err = s.auth.Check(ctx, req.GetKey())
	if err != nil {
		return
	}

	resp.List, err = s.uc.List(ctx, getAll)
	return
}
