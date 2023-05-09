package logic

import (
	"context"
	"example.com/m/v2/todoList/internal/infra"
	"example.com/m/v2/todoList/internal/svc"
	"example.com/m/v2/todoList/todoList"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7/esutil"
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
	//新建查询
	resp := &todoList.SearchTodoListResp{
		Ok:       true,
		TodoList: make([]*todoList.TodoList, 0),
	}

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"wildcard": map[string]interface{}{
				"extra": map[string]interface{}{
					"value": "*" + in.Keyword + "*",
				},
			},
		},
	}
	res, err := infra.ESClient.Search(
		infra.ESClient.Search.WithContext(context.Background()),
		infra.ESClient.Search.WithIndex("todo_list"),
		infra.ESClient.Search.WithBody(esutil.NewJSONReader(&query)),
		infra.ESClient.Search.WithTrackTotalHits(true),
		infra.ESClient.Search.WithPretty(),
	)
	if err != nil {
		fmt.Println("Error getting response: ", err)
		return resp, err
	}
	fmt.Printf("%v", res)
	return resp, nil
}
