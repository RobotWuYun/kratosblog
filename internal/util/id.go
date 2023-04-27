package util

import "context"

func GetID(ctx context.Context, path string) (string, error) {
	return AES.EncryptByAes(path), nil
}

func GetPath(ctx context.Context, id string) (string, error) {
	return AES.DecryptByAes(id)
}
