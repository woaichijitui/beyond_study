package svc

import (
	"beyond_study/application/applet/internal/config"
	"beyond_study/application/user/rpc/user"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	UserRpc  user.User
	BizRedis *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	bizRedis, err := redis.NewRedis(c.BizRedis)
	//todo log
	fmt.Println(err)
	return &ServiceContext{
		Config:   c,
		UserRpc:  user.NewUser(zrpc.MustNewClient(c.UserRPC)),
		BizRedis: bizRedis,
	}
}
