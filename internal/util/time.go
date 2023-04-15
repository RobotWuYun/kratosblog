package util

import (
	"context"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/types/known/timestamppb"
	"sort"
	"time"
)

func Sort(ctx context.Context, times []time.Time) (result []time.Time) {
	var timeStamps []string
	for _, v := range times {
		timeStamps = append(timeStamps, string(timestamppb.New(v).Seconds))
	}
	sort.Sort(sort.Reverse(sort.StringSlice(timeStamps)))
	for _, v := range timeStamps {
		result = append(result, timestamppb.New(time.Unix(cast.ToInt64(v), 0)).AsTime())
	}
	return
}
