package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	dataV1 "kratosblog/api/kratosblog/data/v1"
	"kratosblog/internal/biz"
	"kratosblog/internal/util"
	"time"
)

type articleRepo struct {
	data *Data
	log  *log.Helper

	cate biz.CategoryRepo
}

type file struct {
	Name string
	Path string
}

// NewArticleRepo .
func NewArticleRepo(data *Data, logger log.Logger, repo biz.CategoryRepo) biz.ArticleRepo {
	return &articleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *articleRepo) List(ctx context.Context, offset, limit int64) (total int64, list []*dataV1.Article, err error) {
	return
}

func (r *articleRepo) getFilesFromDir(ctx context.Context) (total int64, list []file, err error) {
	// 获取文章
	var cateMap map[string]string
	cateMap, err = r.cate.MapPath(ctx)
	if err != nil {
		return
	}

	var timeFileMap = make(map[time.Time]map[string]string)
	var dates []time.Time

	for cate, path := range cateMap {
		_, files, err := util.GetAllMarkdownMap(ctx, path)
		if err != nil {
			r.log.Error("getFilesFromDir", err)
			continue
		}
		for name, date := range files {
			dates = append(dates, date)
			timeFileMap[date] = map[string]string{
				name: cate + "/" + name,
			}
		}
	}

	dates = util.Sort(ctx, dates)
	for _, date := range dates {
		for name, path := range timeFileMap[date] {
			list = append(list, file{
				Name: name,
				Path: path,
			})
		}
	}
	return
}
