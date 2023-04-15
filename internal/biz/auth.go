package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

// AuthRepo is a Greater repo.
type AuthRepo interface {
	Make(context.Context) error
	Check(ctx context.Context) (bool, error)
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
