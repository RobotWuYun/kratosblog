package util

import (
	"io/ioutil"
)

func GetAllLeafCate(source string) (cates []string, err error) {
	paths, err := ioutil.ReadDir(source)
	if err != nil {
		return
	}
	childCates := make([]string, 0)
	for _, path := range paths {
		if path.IsDir() {
			childCates, err = GetAllLeafCate(source + "/" + path.Name())
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

func GetAllFileName(source string) (total int64, files []string, err error) {
	path, err := ioutil.ReadDir(source)
	if err != nil {
		return
	}
	for _, v := range path {
		if !v.IsDir() {
			files = append(files, v.Name())
		}
	}
	total = int64(len(files))
	return
}
