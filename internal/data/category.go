package data

import (
	"context"
	dataV1 "kratosblog/api/kratosblog/data/v1"

	"kratosblog/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type categoryRepo struct {
	data *Data
	log  *log.Helper
}

// NewCategoryRepo .
func NewCategoryRepo(data *Data, logger log.Logger) biz.CategoryRepo {
	return &categoryRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *categoryRepo) List(ctx context.Context) (list []*dataV1.Category, err error) {
	return
}
