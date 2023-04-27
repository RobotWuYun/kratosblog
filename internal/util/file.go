package util

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io/ioutil"
	"os"
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

func GetMarkdownContent(ctx context.Context, source string) (content string, err error) {
	var data []byte
	data, err = ioutil.ReadFile(source)
	if err != nil {
		return
	}
	content = string(data)
	return
}

func GetFileInfo(ctx context.Context, source string) (name string, updateAt time.Time, err error) {
	var file os.FileInfo
	file, err = os.Stat(source)
	if err != nil {
		return
	}
	name = file.Name()
	updateAt = file.ModTime()
	return
}

func GetMarkdownInfo(ctx context.Context, source string) (name string, updateAt *timestamppb.Timestamp, content string, err error) {
	var updateAtTime time.Time
	name, updateAtTime, err = GetFileInfo(ctx, source)
	if err != nil {
		return
	}
	name = strings.TrimSuffix(name, ".md")
	updateAt = timestamppb.New(updateAtTime)
	content, err = GetMarkdownContent(ctx, source)
	if err != nil {
		return
	}
	return
}

func GetFileNameWithoutFileFormat(fileName string) (out string) {
	index := strings.Split(fileName, ".")
	if len(index) > 0 {
		out = index[0]
	}
	return
}
