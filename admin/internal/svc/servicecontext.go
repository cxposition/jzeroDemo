package svc

import (
	"admin/internal/config"
	"admin/internal/custom"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	Custom *custom.Custom

	SqlxConn sqlx.SqlConn
	Model    Model
}

func NewServiceContext(c config.Config) *ServiceContext {
	svcContext := &ServiceContext{
		Config:   c,
		Custom:   custom.New(),
		SqlxConn: MustSqlConn(c),
	}

	svcContext.Model = NewModel(svcContext.SqlxConn)
	return svcContext
}
