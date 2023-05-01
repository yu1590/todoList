package logic

import (
	"context"
	"database/sql"
	"example.com/m/v2/model"
	"time"

	"example.com/m/v2/rpc/add"
	"example.com/m/v2/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogic) Add(in *add.AddReq) (*add.AddResp, error) {
	// todo: add your logic here and delete this line
	// 手动代码开始
	_, err := l.svcCtx.Model.Insert(l.ctx, &model.TodoList{
		Id:        1,
		AccountId: 111,
		Time:      time.Now(),
		Extra: sql.NullString{
			String: "12345",
			Valid:  true,
		},
		Status: 1,
	})
	if err != nil {
		return nil, err
	}

	return &add.AddResp{
		Ok: true,
	}, nil
	// 手动代码结束
	return &add.AddResp{}, nil
}
