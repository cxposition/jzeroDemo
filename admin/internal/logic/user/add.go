package user

import (
	"admin/internal/model/user"
	"admin/internal/svc"
	"admin/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type Add struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdd(ctx context.Context, svcCtx *svc.ServiceContext) *Add {
	return &Add{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Add) Add(req *types.AddUserRequest) (resp *types.AddUserResponse, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.Model.User.Insert(l.ctx, &user.User{
		Name: req.Name,
	})
	if err != nil {
		l.Errorf("AddUser failed, err: %v", err)
		return nil, err
	}

	resp = &types.AddUserResponse{
		Status: 0,
	}
	return resp, nil
}
