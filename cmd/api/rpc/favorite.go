package rpc

import (
	"context"
	"fmt"
	"time"

	"github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/favorite"
	"github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/favorite/favoriteservice"
	"github.com/Yra-A/Douyin_Simple_Demo/pkg/constants"
	"github.com/Yra-A/Douyin_Simple_Demo/pkg/errno"
	"github.com/Yra-A/Douyin_Simple_Demo/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var favoriteClient favoriteservice.Client

func initfavoriteRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress}) // 服务发现
	if err != nil {
		panic(err)
	}

	c, err := favoriteservice.NewClient(
		constants.FavoriteServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(tracing.NewClientSuite()),        // tracer
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	favoriteClient = c
}

// favoriteAction 视频投稿【rpc 客户端】
func FavoriteAction(ctx context.Context, req *favorite.FavoriteActionRequest) (*favorite.FavoriteActionResponse, error) {
	resp := new(favorite.FavoriteActionResponse)
	resp, err := favoriteClient.FavoriteAction(ctx, req)
	if err != nil {
		return resp, err
	}
	if resp.StatusCode != 0 {
		return resp, errno.NewErrNo(int64(resp.StatusCode), *resp.StatusMsg)
	}
	return resp, nil
}

func FavoriteList(ctx context.Context, req *favorite.FavoriteListRequest) (*favorite.FavoriteListResponse, error) {
	resp, err := favoriteClient.FavoriteList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int64(resp.StatusCode), *resp.StatusMsg)
	}
	fmt.Println("resp.StatusMsg:", resp.StatusMsg)
	return resp, nil
}
