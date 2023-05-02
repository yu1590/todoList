package logic

import (
	"context"
	"example.com/m/v2/todoList/adder"
	"github.com/jinzhu/copier"

	"example.com/m/v2/api/internal/svc"
	"example.com/m/v2/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTodoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTodoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTodoListLogic {
	return &GetTodoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTodoListLogic) GetTodoList(req *types.GetTodoListReq) (resp *types.GetTodoListResp, err error) {
	// todo: add your logic here and delete this line
	// 手动代码开始
	resp = &types.GetTodoListResp{
		TodoList: make([]*types.TodoList, 0),
		Ok:       true,
	}
	r, err := l.svcCtx.Adder.GetTodoList(l.ctx, &adder.GetTodoListReq{
		AccountID: req.AccountID,
		Months:    req.Months,
	})

	err = copier.Copy(&resp.TodoList, &r.TodoList)
	if err != nil {
		resp.Ok = false
		return resp, err
	}

	return resp, nil

}
