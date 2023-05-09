package logic

import (
	"context"
	"example.com/m/v2/todoList/adder"
	"github.com/jinzhu/copier"

	"example.com/m/v2/api/internal/svc"
	"example.com/m/v2/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchTodoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchTodoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchTodoListLogic {
	return &SearchTodoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchTodoListLogic) SearchTodoList(req *types.SearchTodoListReq) (resp *types.SearchTodoListResp, err error) {
	// todo: add your logic here and delete this line
	resp = &types.SearchTodoListResp{
		TodoList: make([]*types.TodoList, 0),
		Ok:       true,
	}
	r, err := l.svcCtx.Adder.SearchTodoList(l.ctx, &adder.SearchTodoListReq{
		AccountID: req.AccountID,
		Keyword:   req.Keyword,
	})

	err = copier.Copy(&resp.TodoList, &r.TodoList)
	if err != nil {
		resp.Ok = false
		return resp, err
	}
	return
}
