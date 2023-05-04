package logic

import (
	"context"
	"example.com/m/v2/model"
	"github.com/jinzhu/copier"
	"time"

	"example.com/m/v2/todoList/internal/svc"
	"example.com/m/v2/todoList/todoList"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveTodoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveTodoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveTodoListLogic {
	return &SaveTodoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveTodoListLogic) SaveTodoList(in *todoList.SaveTodoListReq) (*todoList.SaveTodoListResp, error) {
	// todo: add your logic here and delete this line
	resp := &todoList.SaveTodoListResp{Ok: true}
	for _, todo := range in.TodoList {
		modelTodo := &model.TodoList{}
		err := copier.Copy(modelTodo, todo)
		modelTodo.Time = time.Unix(todo.Time, 0)
		if err != nil {
			logx.WithContext(l.ctx).Errorf("copier.Copy error: %v", err)
			resp.Ok = false
			continue
		}
		_, err = l.svcCtx.Model.Insert(l.ctx, modelTodo)
		if err != nil {
			logx.WithContext(l.ctx).Errorf("l.svcCtx.Model.Insert error: %v", err)
			resp.Ok = false
			continue
		}
	}
	return resp, nil
}
