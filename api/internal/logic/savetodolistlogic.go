package logic

import (
	"context"
	"example.com/m/v2/todoList/todoList"
	"time"

	"example.com/m/v2/api/internal/svc"
	"example.com/m/v2/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveTodoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveTodoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveTodoListLogic {
	return &SaveTodoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveTodoListLogic) SaveTodoList(req *types.SaveTodoListReq) (resp *types.SaveTodoListResp, err error) {
	// todo: add your logic here and delete this line
	resp = &types.SaveTodoListResp{
		Ok: true,
	}
	todo := &todoList.TodoList{
		AccountId: 123,
		Time:      req.Time,
		Month:     0,
		Extra:     req.Extra,
		Status:    0,
	}
	t := time.Unix(req.Time, 0)
	todo.Month = int64(t.Month())

	rsp, err := l.svcCtx.Adder.SaveTodoList(l.ctx, &todoList.SaveTodoListReq{
		AccountID: 123,
		TodoList:  []*todoList.TodoList{todo},
	})
	if err != nil || rsp.Ok == false {
		resp.Ok = false
		return resp, err
	}
	return resp, nil
}
