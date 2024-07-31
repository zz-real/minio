package logic

import (
	"context"
	"github.com/minio/minio-go/v7"

	"minio/api/internal/svc"
	"minio/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListObjLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListObjLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListObjLogic {
	return &ListObjLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListObjLogic) ListObj(req *types.ObjReq) (resp *types.ListBucketResp, err error) {
	resp = &types.ListBucketResp{}
	objCh := l.svcCtx.Mc.ListObjects(l.ctx, req.Bucket, minio.ListObjectsOptions{})
	var arr []types.BucketResp
	for obj := range objCh {
		arr = append(arr, types.BucketResp{
			Name:         obj.Key,
			CreationDate: obj.LastModified.Format("2006-01-02 15:04:05"),
		})
	}
	resp.List = arr
	return
}
