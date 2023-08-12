package main

import (
	"context"
	"fmt"
	"github.com/Yra-A/Douyin_Simple_Demo/cmd/user/service"
	user "github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/user"
	"github.com/Yra-A/Douyin_Simple_Demo/pkg/errno"
	"github.com/cloudwego/kitex/pkg/klog"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	klog.CtxDebugf(ctx, "UserRegister called: %s", req.GetUsername()+" "+req.GetPassword())
	resp = new(user.UserRegisterResponse)
	if len(req.Username) == 0 || len(req.Password) == 0 {
		fmt.Println(req.Username, req.Password, "!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!rpcserver")
		resp = &user.UserRegisterResponse{
			StatusCode: errno.ParamErrCode,
			StatusMsg:  &errno.ParamErr.ErrMsg,
		}
		return resp, nil
	}
	err = service.NewCreateUserService(ctx).CreateUser(req)
	if err != nil {
		ErrMsg := err.Error()
		resp = &user.UserRegisterResponse{
			StatusCode: errno.FalseCode,
			StatusMsg:  &ErrMsg,
		}
		return resp, nil
	}
	return resp, nil
}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	// TODO: Your code here...
	return
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoRequest) (resp *user.UserInfoResponse, err error) {
	// TODO: Your code here...
	return
}
