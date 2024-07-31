package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"minio/api/internal/logic"
	"minio/api/internal/svc"
	"minio/api/internal/types"
)

func ListObjHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ObjReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewListObjLogic(r.Context(), svcCtx)
		resp, err := l.ListObj(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
