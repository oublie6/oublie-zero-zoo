package basic

import (
	"context"
	"github.com/oublie6/oublie-zero-zoo/app/user-rpc/pb"

	"github.com/oublie6/oublie-zero-zoo/app/user-api/internal/svc"
	"github.com/oublie6/oublie-zero-zoo/app/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BasicIndexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBasicIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BasicIndexLogic {
	return &BasicIndexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BasicIndexLogic) BasicIndex(req *types.BasicIndexReq) (resp *types.BasicIndexResp, err error) {
	rpcResp, err := l.svcCtx.UserRpcClient.ReturnMsg(l.ctx, &pb.ReturnMsgReq{
		Data: req.Data,
	})
	if err != nil {
		return nil, err
	}

	return &types.BasicIndexResp{
		Msg: rpcResp.Data,
	}, nil
}
