package user

import (
	"net/http"

	"github.com/oublie6/oublie-zero-zoo/app/user-api/internal/logic/user"
	"github.com/oublie6/oublie-zero-zoo/app/user-api/internal/svc"
	"github.com/oublie6/oublie-zero-zoo/app/user-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewUserUpdateLogic(r.Context(), svcCtx)
		resp, err := l.UserUpdate(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
