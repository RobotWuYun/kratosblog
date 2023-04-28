package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/timestamppb"
	dataV1 "kratosblog/api/kratosblog/data/v1"
	"kratosblog/internal/biz"
	"kratosblog/internal/util"
	"strings"
)

type articleRepo struct {
	data *Data
	log  *log.Helper

	cate biz.CategoryRepo
}

type file struct {
	Name string
	Path string
	Cate struct {
		ID   string
		Name string
	}
}

// NewArticleRepo .
func NewArticleRepo(data *Data, logger log.Logger, cate biz.CategoryRepo) biz.ArticleRepo {
	return &articleRepo{
		data: data,
		log:  log.NewHelper(logger),
		cate: cate,
	}
}

func (r *articleRepo) getFilesFromDir(ctx context.Context) (total int64, list []file, err error) {
	// 获取文章
	var cateMap map[string]string
	cateMap, err = r.cate.MapPath(ctx)
	if err != nil {
		return
	}

	var timeFileMap = make(map[int64]map[string]string)
	var dates []int64

	for _, path := range cateMap {
		_, files, err := util.GetAllMarkdownMap(ctx, path)
		if err != nil {
			r.log.Error("getFilesFromDir", err)
			continue
		}
		for name, date := range files {
			dates = append(dates, timestamppb.New(date).Seconds)
			timeFileMap[timestamppb.New(date).Seconds] = map[string]string{
				util.GetFileNameWithoutFileFormat(name): path + "/" + name,
			}
		}
	}
	total = int64(len(dates))

	dates = util.Sort(ctx, dates)
	for _, date := range dates {
		for name, path := range timeFileMap[date] {
			data := file{
				Name: name,
				Path: path,
			}
			data.Cate.Name = strings.Split(path, "/")[len(strings.Split(path, "/"))-2]
			data.Cate.ID, _ = util.GetID(ctx, strings.Join(strings.Split(path, "/")[:len(strings.Split(path, "/"))-1], "/"))
			list = append(list, data)
		}
	}
	return
}

func (r *articleRepo) getArticle(ctx context.Context, path ...string) (article map[string]*dataV1.Article, err error) {

	return
}

func (r *articleRepo) List(ctx context.Context, offset, limit int64) (total int64, list []*dataV1.Article, err error) {
	var files []file
	total, files, err = r.getFilesFromDir(ctx)
	if err != nil {
		r.log.Error("articleRepo.List", err)
		return
	}
	if offset > total {
		return
	}
	if offset+limit > total {
		limit = total - offset
	}
	for i := offset; i < offset+limit; i++ {
		var id string
		id, err = util.GetID(ctx, files[i].Path)
		if err != nil {
			continue
		}
		list = append(list, &dataV1.Article{
			Id:    id,
			Title: files[i].Name,
			Category: &dataV1.Category{
				Name: files[i].Cate.Name,
				Id:   files[i].Cate.ID,
			},
		})
	}
	return
}

func (r *articleRepo) Detail(ctx context.Context, id string) (article *dataV1.Article, err error) {
	article = &dataV1.Article{Id: id}

	var path string
	path, err = util.GetPath(ctx, id)
	if err != nil {
		r.log.Error("articleRepo.Detail Decode err .", err)
		return
	}
	article.Title, article.UpdatedAt, article.Content, err = util.GetMarkdownInfo(ctx, path)
	if err != nil {
		r.log.Error("articleRepo.Detail get Detail err .", err)
		return
	}
	return
}
