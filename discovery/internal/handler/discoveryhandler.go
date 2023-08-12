package handler

import (
	"net/http"

	"discovery/internal/logic"
	"discovery/internal/svc"
	"discovery/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DiscoveryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDiscoveryLogic(r.Context(), svcCtx)
		resp, err := l.Discovery(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
