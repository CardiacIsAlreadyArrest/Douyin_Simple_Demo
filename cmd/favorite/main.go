package main

import (
	"net"

	"github.com/Yra-A/Douyin_Simple_Demo/cmd/favorite/dal"
	"github.com/Yra-A/Douyin_Simple_Demo/cmd/favorite/rpc"
	favorite "github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/favorite/favoriteservice"
	"github.com/Yra-A/Douyin_Simple_Demo/pkg/constants"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func Init() {
	dal.Init()
	rpc.InitRPCFavorite()
}

func main() {
	klog.SetLevel(klog.LevelDebug)

	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8893")
	if err != nil {
		panic(err)
	}

	Init()

	svr := favorite.NewServer(new(FavoriteServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.FavoriteServiceName}),
		server.WithServiceAddr(addr),
		server.WithRegistry(r))

	err = svr.Run()

	if err != nil {
		klog.Fatal(err)
	}
}
