package data

import (
	"context"
	"kratosblog/internal/biz"
	"kratosblog/internal/util"

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

func (r *authRepo) Check(ctx context.Context) (pass bool, err error) {
	return util.GetKey(ctx) == "wuyun", nil
}
