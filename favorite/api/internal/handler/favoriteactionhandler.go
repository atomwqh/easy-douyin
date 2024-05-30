package handler

import (
	"net/http"

	"github.com/huangsihao7/scooter-WSVA/favorite/api/internal/logic"
	"github.com/huangsihao7/scooter-WSVA/favorite/api/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/favorite/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FavoriteActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ActionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewFavoriteActionLogic(r.Context(), svcCtx)
		resp, err := l.FavoriteAction(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
