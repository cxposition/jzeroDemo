package role

import (
	"context"

	"admin/internal/svc"
	"admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Delete struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelete(ctx context.Context, svcCtx *svc.ServiceContext) *Delete {
	return &Delete{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Delete) Delete(req *types.DeleteRoleRequest) (resp *types.DeleteRoleResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
