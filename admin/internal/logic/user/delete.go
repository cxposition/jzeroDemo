package user

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

func (l *Delete) Delete(req *types.DeleteUserRequest) (resp *types.DeleteUserResponse, err error) {
	// todo: add your logic here and delete this line
	err = l.svcCtx.Model.User.Delete(l.ctx, uint64(req.Id))
	if err != nil {
		l.Errorf("delete user error: %s", err)
		return nil, err
	}

	resp = &types.DeleteUserResponse{
		Status: 0,
	}
	return resp, nil
}
