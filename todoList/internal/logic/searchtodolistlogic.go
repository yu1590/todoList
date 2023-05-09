package logic

import (
	"context"

	"example.com/m/v2/todoList/internal/svc"
	"example.com/m/v2/todoList/todoList"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchTodoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchTodoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchTodoListLogic {
	return &SearchTodoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchTodoListLogic) SearchTodoList(in *todoList.SearchTodoListReq) (*todoList.SearchTodoListResp, error) {
	// todo: add your logic here and delete this line

	return &todoList.SearchTodoListResp{
		Ok: true,
		TodoList: []*todoList.TodoList{{
			Id:        1,
			AccountId: 1,
			Time:      1684342920,
			Month:     5,
			Extra:     "123323",
			Status:    2,
		}},
	}, nil
}
