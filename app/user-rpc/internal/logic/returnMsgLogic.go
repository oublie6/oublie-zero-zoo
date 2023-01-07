package logic

import (
	"context"
	"os"

	"github.com/oublie6/oublie-zero-zoo/app/user-rpc/internal/svc"
	"github.com/oublie6/oublie-zero-zoo/app/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReturnMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReturnMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReturnMsgLogic {
	return &ReturnMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ReturnMsgLogic) ReturnMsg(in *pb.ReturnMsgReq) (*pb.ReturnMsgResp, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	return &pb.ReturnMsgResp{
		Data: in.Data + " rpc --->>>> " + hostname,
	}, nil
}
