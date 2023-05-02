package handler

import (
	"net/http"

	"example.com/m/v2/api/internal/logic"
	"example.com/m/v2/api/internal/svc"
	"example.com/m/v2/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SaveTodoListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SaveTodoListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSaveTodoListLogic(r.Context(), svcCtx)
		resp, err := l.SaveTodoList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
