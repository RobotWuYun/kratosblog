package data

import (
	"context"
	"kratosblog/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type authRepo struct {
	data *Data
	log  *log.Helper
}

// NewAuthRepo .
func NewAuthRepo(data *Data, logger log.Logger) biz.AuthRepo {
	return &authRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *authRepo) Make(ctx context.Context) (err error) {
	return
}

func (r *authRepo) Check(ctx context.Context, key string) (pass bool, err error) {
	return key == "wuyun", nil
}
