package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"minio/api/internal/logic"
	"minio/api/internal/svc"
)

func ListBucketHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewListBucketLogic(r.Context(), svcCtx)
		resp, err := l.ListBucket()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
