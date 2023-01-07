package user

import (
	"context"
	"github.com/oublie6/oublie-zero-zoo/app/user-api/internal/svc"
	"github.com/oublie6/oublie-zero-zoo/app/user-api/internal/types"
	"github.com/oublie6/oublie-zero-zoo/app/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {

	//if err := l.testOne(); err != nil {
	//	logx.Errorf("err:%+v", err)
	//}
	userResp, err := l.svcCtx.UserRpcClient.GetUserInfo(l.ctx, &pb.GetUserInfoReq{
		Id: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	//l.Logger.Error("log info")
	//
	//fmt.Println("user info ing")
	//
	//user, err := l.svcCtx.UserModel.FindOne(l.ctx, req.UserId)
	//if err != nil && err != model.ErrNotFound {
	//	return nil, errors.New("查询数据失败")
	//}
	//if user == nil {
	//	return nil, errors.New("用户不存在")
	//}

	return &types.UserInfoResp{
		UserId:   userResp.UserModel.Id,
		Nickname: userResp.UserModel.Nickname,
	}, nil
}

//func (l *UserInfoLogic) testOne() error {
//	return l.testTwo()
//}
//
//func (l *UserInfoLogic) testTwo() error {
//	return l.testThree()
//}
//
//func (l *UserInfoLogic) testThree() error {
//	return errors.Wrap(errors.New("这是故意的"), "enen")
//}
