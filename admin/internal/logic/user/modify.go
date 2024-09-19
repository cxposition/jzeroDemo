package user

import (
	"admin/internal/model/user"
	"context"

	"admin/internal/svc"
	"admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Modify struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModify(ctx context.Context, svcCtx *svc.ServiceContext) *Modify {
	return &Modify{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Modify) Modify(req *types.ModifyUserRequest) (resp *types.ModifyUserResponse, err error) {
	// todo: add your logic here and delete this line
	err = l.svcCtx.Model.User.Update(l.ctx, &user.User{
		Id:   uint64(req.Id),
		Name: req.Name,
	})
	if err != nil {
		l.Errorf("update user failed, err: %v", err)
		return nil, err
	}

	resp = &types.ModifyUserResponse{
		Status: 0,
	}
	return resp, nil
}
