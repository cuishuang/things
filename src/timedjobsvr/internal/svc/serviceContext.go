package svc

import (
	"github.com/i-Things/things/shared/stores"
	"github.com/i-Things/things/src/timedjobsvr/internal/config"
	"github.com/i-Things/things/src/timedjobsvr/internal/repo/event/publish/pubJob"
	"github.com/i-Things/things/src/timedjobsvr/internal/repo/relationDB"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/kv"
	"os"
)

type ServiceContext struct {
	Config config.Config
	Store  kv.Store
	PubJob *pubJob.PubJob
}

func NewServiceContext(c config.Config) *ServiceContext {
	pj, err := pubJob.NewPubJob(c.Event)
	if err != nil {
		logx.Error("初始化消息队列 err", err)
		os.Exit(-1)
	}
	stores.InitConn(c.Database)
	err = relationDB.Migrate(c.Database)
	if err != nil {
		logx.Error("timedjobsvr 数据库初始化失败 err", err)
		os.Exit(-1)
	}
	return &ServiceContext{
		Config: c,
		PubJob: pj,
		Store:  kv.NewStore(c.CacheRedis),
	}
}
