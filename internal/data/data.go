package data

import (
	"kratosblog/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewGreeterRepo,
	NewCategoryRepo,
	NewAuthRepo,
)

// Data .
type Data struct {
	conf *conf.Bootstrap
}

// NewData .
func NewData(c *conf.Bootstrap, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		conf: c,
	}, cleanup, nil
}
