package file

import (
	"net/http"

	"admin/internal/logic/file"
	"admin/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func Upload(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := file.NewUpload(r.Context(), svcCtx, r)
		resp, err := l.Upload()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
