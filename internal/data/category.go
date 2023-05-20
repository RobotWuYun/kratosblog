package data

import (
	"context"
	dataV1 "kratosblog/api/kratosblog/data/v1"
	"kratosblog/internal/biz"
	"kratosblog/internal/util"
	"strings"
)

type categoryRepo struct {
	data *Data

	auth biz.AuthRepo
}

// DisableDirList
var DisableDirList = []string{
	".git",
}

// NewCategoryRepo .
func NewCategoryRepo(
	data *Data,
	auth biz.AuthRepo,
) biz.CategoryRepo {
	return &categoryRepo{
		data: data,
		auth: auth,
	}
}

func (r *categoryRepo) List(ctx context.Context) (list []*dataV1.Category, err error) {
	cates, err := r.MapPath(ctx)
	if err != nil {
		return
	}
	for name, _ := range cates {
		list = append(list, &dataV1.Category{
			Name: name,
		})
	}
	return
}
func (r *categoryRepo) MapPath(ctx context.Context) (cateMap map[string]string, err error) {
	return
}

func (r *categoryRepo) mapPathForFS(ctx context.Context) (cateMap map[string]string, err error) {
	cateMap = make(map[string]string)
	var getAll bool
	if getAll, err = r.auth.Check(ctx); err != nil {
		return
	}

	var allCates []string
	allCates, err = util.GetAllLeafCate(ctx, r.data.conf.GetData().GetRepository().GetPath())
	if err != nil {
		r.data.Log.Error("GetAllLeafCate err .", err)
		return
	}
	for _, cate := range allCates {
		pathIndex := strings.Split(cate, "/")
		name := pathIndex[len(pathIndex)-1]
		if !getAll {
			if strings.HasSuffix(name, r.data.conf.GetData().GetRepository().GetPrivateSuffix()) {
				continue
			}
		}
		add := true
		for _, v := range DisableDirList {
			if strings.Contains(cate, v) {
				add = false
			}
		}
		if !add {
			continue
		}
		cateMap[name] = cate
	}
	return
}
