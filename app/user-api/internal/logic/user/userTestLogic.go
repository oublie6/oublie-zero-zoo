package user

import (
	"context"
	"fmt"

	"github.com/oublie6/oublie-zero-zoo/app/user-api/internal/svc"
	"github.com/oublie6/oublie-zero-zoo/app/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserTestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserTestLogic {
	return &UserTestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserTestLogic) UserTest(req *types.UserTestReq) (resp *types.UserTestResp, err error) {
	fmt.Println("user test ing")
	return
}
