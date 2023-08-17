package main

import (
	"context"
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
		resp = &user.UserRegisterResponse{
			StatusCode: errno.UserAlreadyExistErrCode,
			StatusMsg:  &errno.UsernameOrPasswordNilErr.ErrMsg,
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
	resp = &user.UserRegisterResponse{
		StatusCode: errno.SuccessCode,
		StatusMsg:  &errno.Success.ErrMsg,
	}
	return resp, nil
}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	klog.CtxDebugf(ctx, "UserLogin called: %s", req.GetUsername()+" "+req.GetPassword())
	resp = new(user.UserLoginResponse)
	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp = &user.UserLoginResponse{
			StatusCode: errno.UsernameOrPasswordNilErrCode,
			StatusMsg:  &errno.UsernameOrPasswordNilErr.ErrMsg,
		}
		return resp, nil
	}
	uid, err := service.NewCheckUserService(ctx).CheckUser(req)
	if err != nil {
		ErrMsg := err.Error()
		resp = &user.UserLoginResponse{
			StatusCode: errno.FalseCode,
			StatusMsg:  &ErrMsg,
		}
		return resp, nil
	}
	resp = &user.UserLoginResponse{
		StatusCode: errno.SuccessCode,
		StatusMsg:  &errno.Success.ErrMsg,
		UserId:     uid,
	}
	return resp, nil
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoRequest) (resp *user.UserInfoResponse, err error) {
	klog.CtxDebugf(ctx, "UserLogin called: uid is %s", req.GetUserId())
	resp = new(user.UserInfoResponse)

	// TODO： 带其他模块完成后补全，暂时返回一个空用户信息
	resp = &user.UserInfoResponse{
		StatusCode: errno.SuccessCode,
		StatusMsg:  &errno.Success.ErrMsg,
		User:       nil,
	}
	return resp, nil
}
