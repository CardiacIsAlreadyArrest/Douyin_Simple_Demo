/*
 * Copyright 2022 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package jwt

import (
	"context"
	"fmt"
	huser "github.com/Yra-A/Douyin_Simple_Demo/cmd/api/biz/model/user"
	"github.com/Yra-A/Douyin_Simple_Demo/cmd/api/rpc"
	kuser "github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/user"
	"github.com/Yra-A/Douyin_Simple_Demo/pkg/constants"
	"github.com/Yra-A/Douyin_Simple_Demo/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/jwt"
	"net/http"
	"time"
)

var JwtMiddleware *jwt.HertzJWTMiddleware

func InitJwt() {
	var err error
	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Key:         []byte(constants.SecretKey),
		TimeFunc:    time.Now,
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		TokenLookup: "query:token,form:token",
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			hlog.CtxInfof(ctx, "Login success ，token is issued clientIP: "+c.ClientIP())
			c.Set("token", token)
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			fmt.Println("Authenticator 成功")
			var req huser.UserLoginRequest
			if err := c.BindAndValidate(&req); err != nil {
				return nil, jwt.ErrMissingLoginValues
			}
			if len(req.Username) == 0 || len(req.Password) == 0 {
				return nil, jwt.ErrMissingLoginValues
			}

			user, err := rpc.UserLogin(context.Background(), &kuser.UserLoginRequest{
				Username: req.Username,
				Password: req.Password,
			})
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}
			c.Set("user_id", user.UserId)
			return user.UserId, nil
		},
		// Verify token and get the id of logged-in user
		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			//fmt.Println("Authorization 中", data, reflect.TypeOf(data))
			if v, ok := data.(float64); ok {
				fmt.Println("Authorization 成功")
				current_user_id := int64(v)
				c.Set("current_user_id", current_user_id)
				hlog.CtxInfof(ctx, "Token is verified clientIP: "+c.ClientIP())
				return true
			}
			return false
		},
		IdentityKey: constants.IdentityKey,

		PayloadFunc: func(data interface{}) jwt.MapClaims {
			// 将 userid 存入 token 的负载部分
			//fmt.Println("PayloadFunc 中", data, reflect.TypeOf(data))
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					constants.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			hlog.CtxErrorf(ctx, "jwt biz err = %+v", e.Error())
			return e.Error()
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusOK, huser.UserLoginResponse{
				StatusCode: errno.AuthorizationFailedErrCode,
				StatusMsg:  &errno.AuthorizationFailedErr.ErrMsg,
			})
		},
		//ParseOptions: []jwtv4.ParserOption{
		//	jwtv4.WithValidMethods([]string{"HS256"}),
		//	//jwtv4.WithJSONNumber(),
		//},
	})
	if err != nil {
		panic(err)
	}
}
