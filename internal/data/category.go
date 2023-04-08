package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	dataV1 "kratosblog/api/kratosblog/data/v1"
	"kratosblog/internal/biz"
	"kratosblog/internal/util"
	"strings"
)

type categoryRepo struct {
	data *Data
	log  *log.Helper
}

// DisableDirList
var DisableDirList = []string{
	".git",
}

// NewCategoryRepo .
func NewCategoryRepo(data *Data, logger log.Logger) biz.CategoryRepo {
	return &categoryRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *categoryRepo) List(ctx context.Context, getAll bool) (list []*dataV1.Category, err error) {
	var cates []string
	cates, err = util.GetAllLeafCate(r.data.conf.GetData().GetRepository().GetPath())
	if err != nil {
		return
	}
	for _, cate := range cates {
		add := true
		for _, v := range DisableDirList {
			if strings.Contains(cate, v) {
				add = false
			}
		}
		if !add {
			continue
		}
		nameIndex := strings.Split(cate, "/")
		list = append(list, &dataV1.Category{
			Name: nameIndex[len(nameIndex)-1],
		})
	}
	return
}
