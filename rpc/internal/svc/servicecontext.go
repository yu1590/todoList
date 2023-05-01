package svc

import (
	"example.com/m/v2/model"
	"example.com/m/v2/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	Model  model.TodoListModel // 手动代码
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  model.NewTodoListModel(sqlx.NewMysql(c.DataSource)), // 手动代码
	}
}
