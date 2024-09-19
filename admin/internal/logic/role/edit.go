package role

import (
	"context"

	"admin/internal/svc"
	"admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Edit struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEdit(ctx context.Context, svcCtx *svc.ServiceContext) *Edit {
	return &Edit{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Edit) Edit(req *types.EditRoleRequest) (resp *types.EditRoleResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
