package handler

import (
	"net/http"

	"example.com/m/v2/api/internal/logic"
	"example.com/m/v2/api/internal/svc"
	"example.com/m/v2/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SearchTodoListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchTodoListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSearchTodoListLogic(r.Context(), svcCtx)
		resp, err := l.SearchTodoList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
