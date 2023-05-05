package data

import (
	"go.uber.org/zap"
	"kratosblog/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"github.com/redis/go-redis/v9"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewGreeterRepo,
	NewCategoryRepo,
	NewAuthRepo,
	NewArticleRepo,
)

// Data .
type Data struct {
	conf *conf.Bootstrap
	Log  *log.Helper

	Redis *redis.Client
}

type initFn func(*conf.Bootstrap) (cleanup func(), err error)

// NewData .
func NewData(c *conf.Bootstrap, logger log.Logger) (data *Data, cleanup func(), err error) {
	data = &Data{
		conf: c,
		Log:  log.NewHelper(logger),
	}

	closeFns := make([]func(), 0)

	initFns := []initFn{
		data.initCache,
		//data.initMongo,
		//data.initMQ,
		//data.initMQJob,
		//data.initES, // 放在初始化缓存之后

	}

	for _, initFn := range initFns {
		var closeFn func()
		closeFn, err = initFn(c)
		if err != nil {
			return
		}
		closeFns = append(closeFns, closeFn)
	}

	// 初始化清理函数
	cleanup = func() {
		log.Info("closing the data resources")
		for _, fn := range closeFns {
			if fn != nil {
				fn()
			}
		}
	}
	return
}

func (d *Data) initCache(c *conf.Bootstrap) (cleanup func(), err error) {
	redisConf := c.GetData().GetRedis()
	d.Redis = redis.NewClient(&redis.Options{
		Addr:         redisConf.GetAddr(),
		Password:     redisConf.GetPassword(), // 没有密码，默认值
		DB:           int(redisConf.GetDb()),  // 默认DB 0
		ReadTimeout:  redisConf.GetReadTimeout().AsDuration(),
		WriteTimeout: redisConf.GetWriteTimeout().AsDuration(),
	})

	cleanup = func() {
		d.Log.Info("closing the redis resources")
		err = d.Redis.Close()
		if err != nil {
			d.Log.Info("closing the redis resources err.", zap.Error(err))
		}
	}
	return
}
