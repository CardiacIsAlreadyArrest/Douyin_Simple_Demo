package handler

import (
	"github.com/Yra-A/Douyin_Simple_Demo/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

type Response struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

// BadResponse 返回错误状态码与错误信息
func BadResponse(c *app.RequestContext, err error) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusBadRequest, Response{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
	})
}

// SendResponse pack response
func SendResponse(c *app.RequestContext, data interface{}) {
	c.JSON(http.StatusOK, data)
}
