package main

import (
	"context"

	"github.com/Yra-A/Douyin_Simple_Demo/cmd/favorite/service"
	favorite "github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/favorite"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// ActionFavorite implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteAction(ctx context.Context, req *favorite.FavoriteActionRequest) (resp *favorite.FavoriteActionResponse, err error) {
	resp = new(favorite.FavoriteActionResponse)

	if req.ActionType <= 0 || req.ActionType >= 3 {
		resp.StatusCode = 1
		*resp.StatusMsg = "点赞不合法,action_type只能是1或者2"
		return resp, nil
	}

	err = service.NewActionFavoriteService(ctx).ActionFavorite(req)

	if err != nil {
		resp.StatusCode = 1
		*resp.StatusMsg = "点赞失败"
		return resp, nil
	}
	resp.StatusCode = 0
	*resp.StatusMsg = "点赞成功"
	return resp, nil
}

// FavoriteList implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteList(ctx context.Context, req *favorite.FavoriteListRequest) (resp *favorite.FavoriteListResponse, err error) {
	resp = new(favorite.FavoriteListResponse)
	if req.UserId < 0 {
		resp.StatusCode = 1
		*resp.StatusMsg = "UserId非法"
		return resp, err
	}

	video_list, err := service.NewFavoriteListService(ctx).FavoriteList(req)

	if err != nil {
		resp.StatusCode = 1
		*resp.StatusMsg = "拉取点赞视频失败"
		return resp, nil
	}

	resp.StatusCode = 0
	*resp.StatusMsg = "拉取点赞视频成功"
	resp.VideoList = video_list
	return resp, nil
}

// FavoriteCount implements the FavoriteServiceImpl interface.
// videoid  how many people like
func (s *FavoriteServiceImpl) FavoriteCount(ctx context.Context, req *favorite.FavoriteCountRequest) (resp *favorite.FavoriteCountResponse, err error) {
	var favorite_count int64
	favorite_count = 0
	favorite_count, err = service.NewFavoriteCountService(ctx).FavoriteCount(req)

	resp = new(favorite.FavoriteCountResponse)
	resp.FavoriteCount = favorite_count
	return resp, err
}

// IsFavorite implements the FavoriteServiceImpl interface.
//
//	ueser_id like video_id
func (s *FavoriteServiceImpl) IsFavorite(ctx context.Context, req *favorite.IsFavoriteRequest) (resp *favorite.IsFavoriteResponse, err error) {
	resp = new(favorite.IsFavoriteResponse)
	resp.IsFavorite, err = service.NewIsFavoriteService(ctx).IsFavorite(req)
	return resp, err
}
