package user

import (
	"net/http"

	"admin/internal/logic/user"
	"admin/internal/svc"
	"admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func Modify(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ModifyUserRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewModify(r.Context(), svcCtx)
		resp, err := l.Modify(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
