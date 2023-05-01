package logic

import (
	"context"

	"example.com/m/v2/rpc/add"
	"example.com/m/v2/rpc/internal/svc"

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

func (l *GetTodoListLogic) GetTodoList(in *add.AddReq) (*add.AddResp, error) {
	// todo: add your logic here and delete this line
	resp := &add.AddResp{Ok: true}
	todoList, err := l.svcCtx.Model.GetTodoListByAccountIDAndTime(l.ctx, in.AccountID, in.Months)
	if err != nil {
		resp.Ok = false
		return resp, err
	}
	logx.WithContext(l.ctx).Infof("todoList: %v", todoList)
	return resp, nil
}
