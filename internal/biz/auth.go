package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

// AuthRepo is a Greater repo.
type AuthRepo interface {
	Check(context.Context, string) (bool, error)
	Make(context.Context) error
}

// AuthUsecase is a Auth usecase.
type AuthUsecase struct {
	repo AuthRepo
	log  *log.Helper
}

// NewAuthUsecase new a Auth usecase.
func NewAuthUsecase(repo AuthRepo, logger log.Logger) *AuthUsecase {
	return &AuthUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *AuthUsecase) Check(ctx context.Context, key string) (bool, error) {
	return uc.repo.Check(ctx, key)
}
