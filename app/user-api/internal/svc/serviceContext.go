package svc

import (
	"github.com/oublie6/oublie-zero-zoo/app/user-api/internal/config"
	"github.com/oublie6/oublie-zero-zoo/app/user-api/internal/middleware"
	"github.com/oublie6/oublie-zero-zoo/app/user-api/model"
	"github.com/oublie6/oublie-zero-zoo/app/user-rpc/usercenter"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	TestMiddleware rest.Middleware
	UserModel      model.UserModel
	UserDataModel  model.UserDataModel
	UserRpcClient  usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		TestMiddleware: middleware.NewTestMiddleware().Handle,
		UserModel:      model.NewUserModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		UserDataModel:  model.NewUserDataModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		UserRpcClient:  usercenter.NewUsercenter(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
