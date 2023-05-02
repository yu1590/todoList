package handler

import (
	"net/http"

	"example.com/m/v2/api/internal/logic"
	"example.com/m/v2/api/internal/svc"
	"example.com/m/v2/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetTodoListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetTodoListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetTodoListLogic(r.Context(), svcCtx)
		resp, err := l.GetTodoList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
