// Code generated by hertz generator.

package user

import (
	"context"
	"fmt"
	"github.com/Yra-A/Douyin_Simple_Demo/cmd/api/biz/handler"
	huser "github.com/Yra-A/Douyin_Simple_Demo/cmd/api/biz/model/user"
	"github.com/Yra-A/Douyin_Simple_Demo/cmd/api/biz/mw"
	"github.com/Yra-A/Douyin_Simple_Demo/cmd/api/rpc"
	kuser "github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/user"
	"github.com/Yra-A/Douyin_Simple_Demo/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

// UserRegister .
// @router /douyin/user/register/ [POST]
func UserRegister(ctx context.Context, c *app.RequestContext) {
	var req huser.UserRegisterRequest
	if err := c.BindAndValidate(&req); err != nil {
		handler.BadResponse(c, errno.ConvertErr(err))
		return
	}
	if len(req.Username) == 0 || len(req.Password) == 0 {
		handler.BadResponse(c, errno.ParamErr)
		return
	}
	if _, err := rpc.UserRegister(context.Background(), &kuser.UserRegisterRequest{
		Username: req.Username,
		Password: req.Password,
	}); err != nil {
		handler.BadResponse(c, errno.ConvertErr(err))
		return
	}
	mw.JwtMiddleware.LoginHandler(ctx, c)
	v, _ := c.Get("user_id")
	user_id := v.(int64)
	token := c.GetString("token")

	handler.SendResponse(c, utils.H{
		"status_code": errno.Success.ErrCode,
		"status_msg":  errno.Success.ErrMsg,
		"user_id":     user_id,
		"token":       token,
	})
}

// UserLogin .
// @router /douyin/user/login/ [POST]
func UserLogin(ctx context.Context, c *app.RequestContext) {
	mw.JwtMiddleware.LoginHandler(ctx, c)
	v, ok := c.Get("user_id")
	if !ok {
		fmt.Println("密码不对!")
		handler.BadResponse(c, errno.LoginFailedErr)
		return
	}
	user_id := v.(int64)
	token := c.GetString("token")

	handler.SendResponse(c, utils.H{
		"status_code": errno.Success.ErrCode,
		"status_msg":  errno.Success.ErrMsg,
		"user_id":     user_id,
		"token":       token,
	})
}

// UserInfo .
// @router /douyin/user/ [GET]
func UserInfo(ctx context.Context, c *app.RequestContext) {
	fmt.Println("????????????????????????!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	var req huser.UserInfoRequest
	if err := c.BindAndValidate(&req); err != nil {
		handler.BadResponse(c, errno.ConvertErr(err))
		return
	}

	u, err := rpc.UserInfo(context.Background(), &kuser.UserInfoRequest{
		UserId: req.UserID,
	})
	if err != nil {
		handler.BadResponse(c, errno.ConvertErr(err))
		return
	}
	resp := huser.UserInfoResponse{
		StatusCode: errno.SuccessCode,
		StatusMsg:  &errno.Success.ErrMsg,
		User: &huser.User{
			ID:   req.UserID,
			Name: u.User.Name,

			// TODO 待其他模块完善后补全
			FollowCount:     nil,
			FollowerCount:   nil,
			IsFollow:        false,
			Avatar:          nil,
			BackgroundImage: nil,
			Signature:       nil,
			TotalFavorited:  nil,
			WorkCount:       nil,
			FavoriteCount:   nil,
		},
	}
	handler.SendResponse(c, resp)
}
