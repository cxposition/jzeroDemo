package user

import (
	"context"
	"github.com/jzero-io/jzero-contrib/condition"

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

func (l *List) List(req *types.ListUserRequest) (resp *types.ListUserResponse, err error) {
	page, total, err := l.svcCtx.Model.User.PageByCondition(l.ctx, condition.Condition{
		Operator: condition.Limit,
		Value:    req.Size,
	}, condition.Condition{
		Operator: condition.Offset,
		Value:    (req.Page - 1) * req.Size,
	})
	if err != nil {
		l.Errorf("List failed, err: %v", err)
		return nil, err
	}

	resp = &types.ListUserResponse{}

	for _, v := range page {
		resp.List = append(resp.List, types.User{
			Id:   int(v.Id),
			Name: v.Name,
		})
	}
	resp.Total = int(total)
	return
}
