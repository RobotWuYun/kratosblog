package util

import (
	"context"
	"io/ioutil"
	"strings"
	"time"
)

func GetAllLeafCate(ctx context.Context, source string) (cates []string, err error) {
	paths, err := ioutil.ReadDir(source)
	if err != nil {
		return
	}
	childCates := make([]string, 0)
	for _, path := range paths {
		if path.IsDir() {
			childCates, err = GetAllLeafCate(ctx, source+"/"+path.Name())
			if err != nil {
				continue
			}
			cates = append(cates, childCates...)
		}
	}
	if len(childCates) == 0 {
		cates = append(cates, source)
	}
	return
}

func GetAllFileName(ctx context.Context, source string) (total int64, fileMap map[string]time.Time, err error) {
	fileMap = make(map[string]time.Time)
	path, err := ioutil.ReadDir(source)
	if err != nil {
		return
	}
	for _, v := range path {
		if !v.IsDir() {
			fileMap[v.Name()] = v.ModTime()
		}
	}
	total = int64(len(fileMap))
	return
}

func GetAllMarkdownMap(ctx context.Context, source string) (total int64, fileMap map[string]time.Time, err error) {
	fileMap = make(map[string]time.Time)

	var allFileMap map[string]time.Time
	_, allFileMap, err = GetAllFileName(ctx, source)
	if err != nil {
		return
	}

	for name, data := range allFileMap {
		if strings.HasSuffix(name, ".md") {
			fileMap[name] = data
		}
	}
	total = int64(len(fileMap))
	return
}
