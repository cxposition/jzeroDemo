package file

import (
	"admin/internal/svc"
	"admin/internal/types"
	"context"
	"io"
	"os"

	"github.com/zeromicro/go-zero/core/logx"

	"net/http"
)

type Upload struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewUpload(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *Upload {
	return &Upload{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *Upload) Upload() (resp *types.UploadResponse, err error) {
	if l.r.Method == "POST" {
		err := l.r.ParseMultipartForm(10 << 20) // 10MB
		if err != nil {
			return nil, err
		}
	}

	// 假设表单中的文件字段名为"file"
	file, handler, err := l.r.FormFile("file")
	if err != nil {
		return
	}
	defer file.Close()

	// 创建文件保存路径
	dst, err := os.Create("./file/" + handler.Filename)
	if err != nil {
		l.Errorf("create file error: %v", err)
		return
	}
	defer dst.Close()

	// 将上传的文件内容复制到目标文件
	_, err = io.Copy(dst, file)
	if err != nil {
		l.Errorf("copy file error: %v", err)
		return
	}
	return
}
