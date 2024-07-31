package handler

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"io"
	"net/http"
	"os"

	"github.com/zeromicro/go-zero/rest/httpx"
	"minio/api/internal/svc"
	"minio/api/internal/types"
)

func FileGetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileGetRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		b, err := svcCtx.Mc.GetObject(context.Background(), req.Bucket, req.Name, minio.GetObjectOptions{})
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			chunk, _ := io.ReadAll(b)
			fw, _ := os.Create(req.Path + "\\" + req.Name)
			defer fw.Close()
			_, err := io.WriteString(fw, string(chunk))
			if err != nil {
				fmt.Printf("err: %s\n", err)
				return
			}
			httpx.Ok(w)
		}
	}
}
