package svc

import (
	"admin/internal/model/log"
	"admin/internal/model/user"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type Model struct {
	User user.UserModel
	Log  log.LogModel
}

func NewModel(conn sqlx.SqlConn) Model {
	return Model{
		User: user.NewUserModel(conn),
		Log:  log.NewLogModel(conn),
	}
}
