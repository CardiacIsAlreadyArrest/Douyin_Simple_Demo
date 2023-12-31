// Code generated by Kitex v0.6.2. DO NOT EDIT.

package favoriteservice

import (
	"context"
	favorite "github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/favorite"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return favoriteServiceServiceInfo
}

var favoriteServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "FavoriteService"
	handlerType := (*favorite.FavoriteService)(nil)
	methods := map[string]kitex.MethodInfo{
		"FavoriteAction": kitex.NewMethodInfo(favoriteActionHandler, newFavoriteServiceFavoriteActionArgs, newFavoriteServiceFavoriteActionResult, false),
		"FavoriteList":   kitex.NewMethodInfo(favoriteListHandler, newFavoriteServiceFavoriteListArgs, newFavoriteServiceFavoriteListResult, false),
		"FavoriteCount":  kitex.NewMethodInfo(favoriteCountHandler, newFavoriteServiceFavoriteCountArgs, newFavoriteServiceFavoriteCountResult, false),
		"IsFavorite":     kitex.NewMethodInfo(isFavoriteHandler, newFavoriteServiceIsFavoriteArgs, newFavoriteServiceIsFavoriteResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "favorite",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.6.2",
		Extra:           extra,
	}
	return svcInfo
}

func favoriteActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*favorite.FavoriteServiceFavoriteActionArgs)
	realResult := result.(*favorite.FavoriteServiceFavoriteActionResult)
	success, err := handler.(favorite.FavoriteService).FavoriteAction(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFavoriteServiceFavoriteActionArgs() interface{} {
	return favorite.NewFavoriteServiceFavoriteActionArgs()
}

func newFavoriteServiceFavoriteActionResult() interface{} {
	return favorite.NewFavoriteServiceFavoriteActionResult()
}

func favoriteListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*favorite.FavoriteServiceFavoriteListArgs)
	realResult := result.(*favorite.FavoriteServiceFavoriteListResult)
	success, err := handler.(favorite.FavoriteService).FavoriteList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFavoriteServiceFavoriteListArgs() interface{} {
	return favorite.NewFavoriteServiceFavoriteListArgs()
}

func newFavoriteServiceFavoriteListResult() interface{} {
	return favorite.NewFavoriteServiceFavoriteListResult()
}

func favoriteCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*favorite.FavoriteServiceFavoriteCountArgs)
	realResult := result.(*favorite.FavoriteServiceFavoriteCountResult)
	success, err := handler.(favorite.FavoriteService).FavoriteCount(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFavoriteServiceFavoriteCountArgs() interface{} {
	return favorite.NewFavoriteServiceFavoriteCountArgs()
}

func newFavoriteServiceFavoriteCountResult() interface{} {
	return favorite.NewFavoriteServiceFavoriteCountResult()
}

func isFavoriteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*favorite.FavoriteServiceIsFavoriteArgs)
	realResult := result.(*favorite.FavoriteServiceIsFavoriteResult)
	success, err := handler.(favorite.FavoriteService).IsFavorite(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFavoriteServiceIsFavoriteArgs() interface{} {
	return favorite.NewFavoriteServiceIsFavoriteArgs()
}

func newFavoriteServiceIsFavoriteResult() interface{} {
	return favorite.NewFavoriteServiceIsFavoriteResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) FavoriteAction(ctx context.Context, req *favorite.FavoriteActionRequest) (r *favorite.FavoriteActionResponse, err error) {
	var _args favorite.FavoriteServiceFavoriteActionArgs
	_args.Req = req
	var _result favorite.FavoriteServiceFavoriteActionResult
	if err = p.c.Call(ctx, "FavoriteAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FavoriteList(ctx context.Context, req *favorite.FavoriteListRequest) (r *favorite.FavoriteListResponse, err error) {
	var _args favorite.FavoriteServiceFavoriteListArgs
	_args.Req = req
	var _result favorite.FavoriteServiceFavoriteListResult
	if err = p.c.Call(ctx, "FavoriteList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FavoriteCount(ctx context.Context, req *favorite.FavoriteCountRequest) (r *favorite.FavoriteCountResponse, err error) {
	var _args favorite.FavoriteServiceFavoriteCountArgs
	_args.Req = req
	var _result favorite.FavoriteServiceFavoriteCountResult
	if err = p.c.Call(ctx, "FavoriteCount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) IsFavorite(ctx context.Context, req *favorite.IsFavoriteRequest) (r *favorite.IsFavoriteResponse, err error) {
	var _args favorite.FavoriteServiceIsFavoriteArgs
	_args.Req = req
	var _result favorite.FavoriteServiceIsFavoriteResult
	if err = p.c.Call(ctx, "IsFavorite", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
