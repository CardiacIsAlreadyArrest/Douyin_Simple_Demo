// Code generated by hertz generator.

package main

import (
	handler "github.com/Yra-A/Douyin_Simple_Demo/cmd/api/biz/handler"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)
}
