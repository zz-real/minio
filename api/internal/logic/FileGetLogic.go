package logic

import (
	"context"

	"minio/api/internal/svc"
	"minio/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileGetLogic {
	return &FileGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileGetLogic) FileGet(req *types.FileGetRequest) error {
	// todo: add your logic here and delete this line

	return nil
}
