package service

import (
	"context"
	"github.com/Yra-A/Douyin_Simple_Demo/cmd/publish/dal/db"
	"github.com/Yra-A/Douyin_Simple_Demo/cmd/publish/rpc"
	"github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/favorite"
	"github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/publish"
	"github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/klog"
	"sync"
)

type PublishListService struct {
	ctx context.Context
}

// NewPublishListService new CheckUserService
func NewPublishListService(ctx context.Context) *PublishListService {
	return &PublishListService{
		ctx: ctx,
	}
}

func (s *PublishListService) PublishList(req *publish.PublishListRequest) ([]*publish.Video, error) {
	videos, err := db.GetVideoByUserId(req.UserId)
	if err != nil {
		klog.CtxInfof(s.ctx, "获取该用户视频列表失败: "+err.Error())
		return nil, err
	}

	var vs []*publish.Video
	for _, v := range videos {
		favoriteCount, commentCount, isFavorite := s.getOtherVideoInfo(v.ID, req.UserId)
		author := s.getAuthorByUserId(v.AuthorID)

		vs = append(vs, &publish.Video{
			Id:            v.ID,
			Author:        author,
			PlayUrl:       v.PlayURL,
			CoverUrl:      v.CoverURL,
			FavoriteCount: favoriteCount,
			CommentCount:  commentCount,
			IsFavorite:    isFavorite,
			Title:         v.Title,
		})
	}

	return vs, nil
}

// TODO: 实现获取视频相关信息 (并发获取?)
func (s *PublishListService) getOtherVideoInfo(videoID int64, myId int64) (int64, int64, bool) {
	var favoriteCount int64
	var commentCount int64
	var isFavorite bool
	var wg sync.WaitGroup
	wg.Add(2)

	// 获取 favorite count
	go func() {
		defer wg.Done()
		resp, err := rpc.FavoriteCount(s.ctx, &favorite.FavoriteCountRequest{VideoId: videoID})
		favoriteCount = resp.FavoriteCount
		if err != nil {
			klog.CtxInfof(s.ctx, "获取 favorite count 出现错误："+err.Error())
		}
	}()

	// TODO 获取 comment count

	// 获取 is favorite
	go func() {
		defer wg.Done()
		resp, err := rpc.IsFavorite(s.ctx, &favorite.IsFavoriteRequest{UserId: myId, VideoId: videoID})
		isFavorite = resp.IsFavorite
		if err != nil {
			klog.CtxInfof(s.ctx, "获取 is favorite 出现错误："+err.Error())
		}
	}()

	wg.Wait()

	return favoriteCount, commentCount, isFavorite
}

func (s *PublishListService) getAuthorByUserId(UserId int64) *publish.User {
	u, err := rpc.GetUser(context.Background(), &user.UserInfoRequest{
		UserId: UserId,
	})
	if err != nil {
		klog.CtxInfof(s.ctx, "获取视频的 Author 出现错误："+err.Error())

	}
	return &publish.User{
		Id:              u.Id,
		Name:            u.Name,
		FollowCount:     u.FollowCount,
		FollowerCount:   u.FollowerCount,
		IsFollow:        u.IsFollow,
		Avatar:          u.Avatar,
		BackgroundImage: u.BackgroundImage,
		Signature:       u.Signature,
		TotalFavorited:  u.TotalFavorited,
		WorkCount:       u.WorkCount,
		FavoriteCount:   u.FavoriteCount,
	}
}
