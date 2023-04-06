package service

import (
	"context"
	"kratosblog/api/kratosblog/server/v1"

	"kratosblog/internal/biz"
)

// CategoryService is a category service.
type CategoryService struct {
	v1.UnimplementedCategoryServer

	uc *biz.CategoryUsecase
}

// NewCategoryService new a category service.
func NewCategoryService(uc *biz.CategoryUsecase) *CategoryService {
	return &CategoryService{uc: uc}
}

func (s *CategoryService) ListCategory(ctx context.Context, req *v1.ListCategoryRequest) (resp *v1.ListCategoryReply, err error) {
	resp = &v1.ListCategoryReply{}
	return
}
