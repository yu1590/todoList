package svc

import (
	"example.com/m/v2/api/internal/config"
	"example.com/m/v2/rpc/adder"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	Adder  adder.Adder // 手动代码
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Adder:  adder.NewAdder(zrpc.MustNewClient(c.Add)), // 手动代码
	}
}
