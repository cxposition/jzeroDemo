package middleware

import (
	"admin/internal/logic/file"
	"admin/internal/types"
	"bytes"
	"net/http"
)

type OperationLogMiddleware struct {
	logLogic *file.LogLogic
}

func NewOperationLogMiddleware(logLogic *file.LogLogic) *OperationLogMiddleware {
	return &OperationLogMiddleware{
		logLogic: logLogic,
	}
}

func (m *OperationLogMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 响应捕获
		recorder := &responseRecorder{w, 200, bytes.Buffer{}}
		next(recorder, r)

		// 获取响应状态码和响应体
		//respBody := recorder.body.String()
		//statusCode := recorder.statusCode

		// 假设我们已经有用户ID（从 JWT Token 或其他方式获得）
		//userID := int64(1)                     // 示例用户ID
		//action := "request action description" // 具体操作

		operatorLogger := &types.OperateLogger{
			Uuid:        "",
			Operator:    "",
			Ipaddr:      "",
			Actions:     "",
			Contents:    "",
			Status:      0,
			CreateTime:  0,
			EventType:   0,
			EventLevel:  0,
			AuditStatus: 0,
		}

		// 记录日志
		err := m.logLogic.RecordOperationLog(operatorLogger)
		if err != nil {
			// 错误处理
		}
	}
}

// responseRecorder 用于捕获响应
type responseRecorder struct {
	http.ResponseWriter
	statusCode int
	body       bytes.Buffer
}

func (r *responseRecorder) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func (r *responseRecorder) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}
