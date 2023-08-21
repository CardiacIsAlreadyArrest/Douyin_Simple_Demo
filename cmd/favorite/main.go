package main

import (
	"log"
	"net"

	favorite "github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/favorite/favoriteservice"
	"github.com/cloudwego/kitex/server"
)

func main() {

	addr, _ := net.ResolveTCPAddr("tcp", ":8892")
	var opts []server.Option
	opts = append(opts, server.WithServiceAddr(addr))
	svr := favorite.NewServer(new(FavoriteServiceImpl), opts...)
	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
