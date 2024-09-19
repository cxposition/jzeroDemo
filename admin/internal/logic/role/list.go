package role

import (
	"context"

	"admin/internal/svc"
	"admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type List struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewList(ctx context.Context, svcCtx *svc.ServiceContext) *List {
	return &List{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *List) List(req *types.ListRoleRequest) (resp *types.ListRoleResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
