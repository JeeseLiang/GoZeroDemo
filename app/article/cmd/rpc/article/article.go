// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.2
// Source: article.proto

package article

import (
	"context"

	"GoZeroDemo/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddArticleReq      = pb.AddArticleReq
	AddArticleResp     = pb.AddArticleResp
	Article            = pb.Article
	DelArticleReq      = pb.DelArticleReq
	DelArticleResp     = pb.DelArticleResp
	GetArticleByIdReq  = pb.GetArticleByIdReq
	GetArticleByIdResp = pb.GetArticleByIdResp
	SearchArticleReq   = pb.SearchArticleReq
	SearchArticleResp  = pb.SearchArticleResp
	UpdateArticleReq   = pb.UpdateArticleReq
	UpdateArticleResp  = pb.UpdateArticleResp

	ArticleZrpcClient interface {
		// -----------------------article-----------------------
		AddArticle(ctx context.Context, in *AddArticleReq, opts ...grpc.CallOption) (*AddArticleResp, error)
		UpdateArticle(ctx context.Context, in *UpdateArticleReq, opts ...grpc.CallOption) (*UpdateArticleResp, error)
		DelArticle(ctx context.Context, in *DelArticleReq, opts ...grpc.CallOption) (*DelArticleResp, error)
		GetArticleById(ctx context.Context, in *GetArticleByIdReq, opts ...grpc.CallOption) (*GetArticleByIdResp, error)
		SearchArticle(ctx context.Context, in *SearchArticleReq, opts ...grpc.CallOption) (*SearchArticleResp, error)
	}

	defaultArticleZrpcClient struct {
		cli zrpc.Client
	}
)

func NewArticleZrpcClient(cli zrpc.Client) ArticleZrpcClient {
	return &defaultArticleZrpcClient{
		cli: cli,
	}
}

// -----------------------article-----------------------
func (m *defaultArticleZrpcClient) AddArticle(ctx context.Context, in *AddArticleReq, opts ...grpc.CallOption) (*AddArticleResp, error) {
	client := pb.NewArticleClient(m.cli.Conn())
	return client.AddArticle(ctx, in, opts...)
}

func (m *defaultArticleZrpcClient) UpdateArticle(ctx context.Context, in *UpdateArticleReq, opts ...grpc.CallOption) (*UpdateArticleResp, error) {
	client := pb.NewArticleClient(m.cli.Conn())
	return client.UpdateArticle(ctx, in, opts...)
}

func (m *defaultArticleZrpcClient) DelArticle(ctx context.Context, in *DelArticleReq, opts ...grpc.CallOption) (*DelArticleResp, error) {
	client := pb.NewArticleClient(m.cli.Conn())
	return client.DelArticle(ctx, in, opts...)
}

func (m *defaultArticleZrpcClient) GetArticleById(ctx context.Context, in *GetArticleByIdReq, opts ...grpc.CallOption) (*GetArticleByIdResp, error) {
	client := pb.NewArticleClient(m.cli.Conn())
	return client.GetArticleById(ctx, in, opts...)
}

func (m *defaultArticleZrpcClient) SearchArticle(ctx context.Context, in *SearchArticleReq, opts ...grpc.CallOption) (*SearchArticleResp, error) {
	client := pb.NewArticleClient(m.cli.Conn())
	return client.SearchArticle(ctx, in, opts...)
}
