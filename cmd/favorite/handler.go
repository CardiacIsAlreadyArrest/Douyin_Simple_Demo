package main

import (
  "context"
  "github.com/Yra-A/Douyin_Simple_Demo/pkg/errno"

  "github.com/Yra-A/Douyin_Simple_Demo/cmd/favorite/service"
  favorite "github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/favorite"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// FavoriteAction implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteAction(ctx context.Context, req *favorite.FavoriteActionRequest) (resp *favorite.FavoriteActionResponse, err error) {
  resp = new(favorite.FavoriteActionResponse)
  resp.StatusMsg = new(string)

  if req.UserId == 0 {
    resp.StatusCode = 1
    *resp.StatusMsg = "非登录状态，无法进行点赞操作"
    return resp, nil
  }

  if req.ActionType <= 0 || req.ActionType >= 3 {
    resp.StatusCode = 1
    *resp.StatusMsg = "点赞不合法,action_type只能是1或者2"
    return resp, nil
  }

  err = service.NewFavoriteActionService(ctx).FavoriteAction(req)

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
  resp.StatusMsg = new(string)
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
  favorite_count, err := service.NewFavoriteCountService(ctx).FavoriteCount(req)
  if err != nil {
    ErrMsg := err.Error()
    resp = &favorite.FavoriteCountResponse{
      StatusCode: errno.FalseCode,
      StatusMsg:  &ErrMsg,
    }
    return resp, nil
  }
  resp = &favorite.FavoriteCountResponse{
    StatusCode:    errno.SuccessCode,
    StatusMsg:     &errno.Success.ErrMsg,
    FavoriteCount: favorite_count,
  }
  return resp, err
}

// IsFavorite implements the FavoriteServiceImpl interface.
//	ueser_id like video_id
func (s *FavoriteServiceImpl) IsFavorite(ctx context.Context, req *favorite.IsFavoriteRequest) (resp *favorite.IsFavoriteResponse, err error) {
  resp = new(favorite.IsFavoriteResponse)
  resp.IsFavorite, err = service.NewIsFavoriteService(ctx).IsFavorite(req)
  if err != nil {
    ErrMsg := err.Error()
    resp = &favorite.IsFavoriteResponse{
      StatusCode: errno.FalseCode,
      StatusMsg:  &ErrMsg,
    }
    return resp, nil
  }
  return resp, err
}
