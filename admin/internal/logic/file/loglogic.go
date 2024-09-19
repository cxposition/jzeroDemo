package file

import (
	"admin/internal/model/log"
	"admin/internal/svc"
	"admin/internal/types"
	"context"
)

type LogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogLogic {
	return &LogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogLogic) RecordOperationLog(operatorLogger *types.OperateLogger) (err error) {

	log := &log.Log{
		Uuid:        operatorLogger.Uuid,
		Operator:    operatorLogger.Operator,
		Ipaddr:      operatorLogger.Ipaddr,
		Actions:     operatorLogger.Actions,
		Contents:    operatorLogger.Contents,
		Status:      int64(operatorLogger.Status),
		CreateTime:  int64(operatorLogger.CreateTime),
		EventType:   int64(operatorLogger.EventType),
		EventLevel:  int64(operatorLogger.EventLevel),
		AuditStatus: int64(operatorLogger.AuditStatus),
	}

	ret, err := l.svcCtx.Model.Log.Insert(l.ctx, log)
	if err != nil {
		//log.Errorf("insert log error: %s", err)
	}
	return nil
}
