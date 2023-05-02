package logic

import (
	"context"
	"example.com/m/v2/todoList/internal/svc"
	"example.com/m/v2/todoList/todoList"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTodoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTodoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTodoListLogic {
	return &GetTodoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTodoListLogic) GetTodoList(in *todoList.GetTodoListReq) (*todoList.GetTodoListResp, error) {
	// todo: add your logic here and delete this line
	resp := &todoList.GetTodoListResp{Ok: true}
	todo, err := l.svcCtx.Model.GetTodoListByAccountIDAndTime(l.ctx, in.AccountID, in.Months)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("GetTodoListByAccountIDAndTime error: %v", err)
		resp.Ok = false
		return resp, err
	}

	resp.TodoList = make([]*todoList.TodoList, 0)
	err = copier.Copy(&resp.TodoList, &todo)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("copier.Copy error: %v", err)
		resp.Ok = false
		return resp, err
	}
	for i, t := range todo {
		resp.TodoList[i].Time = t.Time.Unix()
	}

	return resp, nil
}
