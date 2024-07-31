package logic

import (
	"context"

	"minio/api/internal/svc"
	"minio/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListBucketLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListBucketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListBucketLogic {
	return &ListBucketLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListBucketLogic) ListBucket() (resp *types.ListBucketResp, err error) {
	resp = &types.ListBucketResp{}
	list, err := l.svcCtx.Mc.ListBuckets(l.ctx)
	var arr []types.BucketResp
	for _, info := range list {
		arr = append(arr, types.BucketResp{
			info.Name, info.CreationDate.Format("2006-01-02 15:04:05"),
		})
	}
	resp.List = arr
	return
}
