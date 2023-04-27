package util

import (
	"context"
	"github.com/spf13/cast"
	"sort"
)

func Sort(ctx context.Context, list []int64) (result []int64) {
	var strs []string
	for _, v := range list {
		strs = append(strs, cast.ToString(v))
	}
	sort.Sort(sort.Reverse(sort.StringSlice(strs)))
	for _, v := range strs {
		result = append(result, cast.ToInt64(v))
	}
	return
}
