package svc

import (
	"admin/internal/config"
	"admin/internal/custom"
	"admin/internal/logic/file"
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	Custom *custom.Custom

	SqlxConn sqlx.SqlConn
	Model    Model
	LogLogic *file.LogLogic
}

func NewServiceContext(c config.Config) *ServiceContext {
	svcContext := &ServiceContext{
		Config:   c,
		Custom:   custom.New(),
		SqlxConn: MustSqlConn(c),
	}

	svcContext.LogLogic = file.NewLogLogic(context.Background(), svcContext)
	svcContext.Model = NewModel(svcContext.SqlxConn)
	return svcContext
}
