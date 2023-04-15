package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	dataV1 "kratosblog/api/kratosblog/data/v1"
)

// CategoryRepo is a Greater repo.
type CategoryRepo interface {
	List(ctx context.Context) (list []*dataV1.Category, err error)
	MapPath(ctx context.Context) (cates map[string]string, err error)
}

// CategoryUsecase is a Category usecase.
type CategoryUsecase struct {
	repo CategoryRepo
	log  *log.Helper
}

// NewCategoryUsecase new a Category usecase.
func NewCategoryUsecase(repo CategoryRepo, logger log.Logger) *CategoryUsecase {
	return &CategoryUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *CategoryUsecase) List(ctx context.Context) (list []*dataV1.Category, err error) {
	return uc.repo.List(ctx)
}
