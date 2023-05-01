package logic

import (
	"context"
	"example.com/m/v2/rpc/adder"

	"example.com/m/v2/api/internal/svc"
	"example.com/m/v2/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req *types.AddReq) (resp *types.AddResp, err error) {
	// todo: add your logic here and delete this line
	// 手动代码开始
	r, err := l.svcCtx.Adder.GetTodoList(l.ctx, &adder.AddReq{
		AccountID: req.AccountID,
		Months:    req.Months,
	})
	if err != nil {
		return nil, err
	}

	return &types.AddResp{
		Ok: r.Ok,
	}, nil
	// 手动代码结束
	return
}
