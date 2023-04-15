package util

import (
	"context"
	"google.golang.org/grpc/metadata"
)

func GetPageQuery(page, pageSize int64) (offset, limit int64) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	offset = (page - 1) * pageSize
	limit = pageSize
	return
}

func GetKey(ctx context.Context) string {
	md, _ := metadata.FromIncomingContext(ctx)
	if len(md.Get("key")) > 0 {
		return md.Get("key")[0]
	}
	return ""
}
