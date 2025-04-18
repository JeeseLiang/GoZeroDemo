package svc

import (
	"GoZeroDemo/app/article/cmd/api/internal/config"
	"GoZeroDemo/app/article/cmd/rpc/article"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	ArticleRpc article.ArticleZrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		ArticleRpc: article.NewArticleZrpcClient(zrpc.MustNewClient(c.ArticleRpcConf)),
	}
}
