package handler

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/zeromicro/go-zero/rest/httpx"
	"minio/api/internal/svc"
	"minio/api/internal/types"
	"net/http"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		// add file
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			return
		}
		defer file.Close()

		b1, err := svcCtx.Mc.PutObject(context.Background(), req.Bucket, fileHeader.Filename, file, fileHeader.Size, minio.PutObjectOptions{ContentType: "application/text"})
		fmt.Println(b1, err)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, b1)
		}
	}
}
