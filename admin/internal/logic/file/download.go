package file

import (
	"admin/internal/svc"
	"admin/internal/types"
	"context"
	"io"
	"os"
	"path/filepath"

	"github.com/zeromicro/go-zero/core/logx"

	"net/http"
)

type Download struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	w http.ResponseWriter
}

func NewDownload(ctx context.Context, svcCtx *svc.ServiceContext, w http.ResponseWriter) *Download {
	return &Download{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,

		w: w,
	}
}

func (l *Download) Download(req *types.DownloadRequest) error {
	filePath := filepath.Join("file", req.File_id)

	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(l.w, "文件不存在", http.StatusNotFound)
		return err
	}
	defer file.Close()

	// 获取文件信息
	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(l.w, "无法获取文件信息", http.StatusInternalServerError)
		return err
	}

	// 设置Content-Disposition响应头，提示浏览器下载而不是显示文件
	l.w.Header().Set("Content-Disposition", "attachment; filename="+fileInfo.Name())

	// 将文件内容复制到HTTP响应中
	_, err = io.Copy(l.w, file)
	if err != nil {
		http.Error(l.w, "文件传输失败", http.StatusInternalServerError)
		return err
	}

	return nil
}
