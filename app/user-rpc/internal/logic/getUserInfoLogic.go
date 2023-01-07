package logic

import (
	"context"

	"github.com/oublie6/oublie-zero-zoo/app/user-rpc/internal/svc"
	"github.com/oublie6/oublie-zero-zoo/app/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	nickname := "unknown"
	m := map[int64]string{
		1: "张三",
		2: "李四",
	}
	if name, ok := m[in.Id]; ok {
		nickname = name
	}

	return &pb.GetUserInfoResp{
		UserModel: &pb.UserModel{
			Id:       in.Id,
			Nickname: nickname,
		},
	}, nil
}
