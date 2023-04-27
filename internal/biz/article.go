package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	dataV1 "kratosblog/api/kratosblog/data/v1"
)

// ArticleRepo is a Greater repo.
type ArticleRepo interface {
	List(ctx context.Context, page, pageSize int64) (int64, []*dataV1.Article, error)
	Detail(ctx context.Context, id string) (*dataV1.Article, error)
}

// ArticleUsecase is a Article usecase.
type ArticleUsecase struct {
	repo ArticleRepo
	log  *log.Helper
}

// NewArticleUsecase new a Article usecase.
func NewArticleUsecase(repo ArticleRepo, logger log.Logger) *ArticleUsecase {
	return &ArticleUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ArticleUsecase) List(ctx context.Context, page, pageSize int64) (total int64, list []*dataV1.Article, err error) {
	return uc.repo.List(ctx, page, pageSize)
}

func (uc *ArticleUsecase) Detail(ctx context.Context, id string) (detail *dataV1.Article, err error) {
	return uc.repo.Detail(ctx, id)
}
