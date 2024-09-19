package file

import (
	"net/http"

	"admin/internal/logic/file"
	"admin/internal/svc"
	"admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func Download(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DownloadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := file.NewDownload(r.Context(), svcCtx, w)
		err := l.Download(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
