package logic

import (
	"context"
	"github.com/minio/minio-go/v7"
	"minio/api/internal/svc"
	"minio/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadLogic) FileUpload(req *types.FileUploadRequest) (resp *types.FileUploadResponse, err error) {
	l.svcCtx.Mc.FGetObject(l.ctx, "test", "2.gif", "", minio.GetObjectOptions{})
	return
}
