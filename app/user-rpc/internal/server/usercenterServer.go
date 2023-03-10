// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package server

import (
	"context"

	"github.com/oublie6/oublie-zero-zoo/app/user-rpc/internal/logic"
	"github.com/oublie6/oublie-zero-zoo/app/user-rpc/internal/svc"
	"github.com/oublie6/oublie-zero-zoo/app/user-rpc/pb"
)

type UsercenterServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUsercenterServer
}

func NewUsercenterServer(svcCtx *svc.ServiceContext) *UsercenterServer {
	return &UsercenterServer{
		svcCtx: svcCtx,
	}
}

func (s *UsercenterServer) GetUserInfo(ctx context.Context, in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	l := logic.NewGetUserInfoLogic(ctx, s.svcCtx)
	return l.GetUserInfo(in)
}

func (s *UsercenterServer) ReturnMsg(ctx context.Context, in *pb.ReturnMsgReq) (*pb.ReturnMsgResp, error) {
	l := logic.NewReturnMsgLogic(ctx, s.svcCtx)
	return l.ReturnMsg(in)
}
