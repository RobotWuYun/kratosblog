package data

import (
	"context"
	"io/ioutil"
	dataV1 "kratosblog/api/kratosblog/data/v1"
	"kratosblog/internal/biz"
	"strings"

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

func (r *categoryRepo) List(ctx context.Context, getAll bool) (list []*dataV1.Category, err error) {
	paths, err := ioutil.ReadDir(r.data.conf.GetData().GetRepository().GetPath())
	if err != nil {
		return
	}
	for _, path := range paths {
		if path.IsDir() && (getAll || !strings.HasSuffix(path.Name(), r.data.conf.GetData().GetRepository().GetPrivateSuffix())) {
			list = append(list, &dataV1.Category{
				Name: path.Name(),
			})
		}
	}
	return
}
