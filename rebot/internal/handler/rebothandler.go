package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"rebot/internal/logic"
	"rebot/internal/svc"
	"rebot/internal/types"
)

func RebotHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRebotLogic(r.Context(), svcCtx)
		resp, err := l.Rebot(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
