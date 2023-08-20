package service

import (
	"context"
	"github.com/Yra-A/Douyin_Simple_Demo/cmd/publish/dal/db"
	"github.com/Yra-A/Douyin_Simple_Demo/cmd/publish/rpc"
	"github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/publish"
	"github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/common/hlog"
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
	videos, err := db.GetVideoByUserID(req.UserId)
	if err != nil {
		hlog.CtxInfof(s.ctx, "获取该用户视频列表失败: "+err.Error())
		return nil, err
	}

	var vs []*publish.Video
	for _, v := range videos {
		author, err := rpc.GetUser(context.Background(), &user.UserInfoRequest{
			UserId: v.AuthorID,
			Token:  req.Token,
		})
		if err != nil {
			hlog.CtxInfof(s.ctx, "用户信息获取失败： "+err.Error())
			return nil, err
		}

		pubAuthor := &publish.User{
			Id:              author.Id,
			Name:            author.Name,
			FollowCount:     author.FollowCount,
			FollowerCount:   author.FollowerCount,
			IsFollow:        author.IsFollow,
			Avatar:          author.Avatar,
			BackgroundImage: author.BackgroundImage,
			Signature:       author.Signature,
			TotalFavorited:  author.TotalFavorited,
			WorkCount:       author.WorkCount,
			FavoriteCount:   author.FavoriteCount,
		}

		// TODO: 可以一次 MGet，而不是每轮循环 get yici

		// TODO： 获取 favoriteCount
		// favoriteCount
		//
		//TODO： 获取 commentCount
		// commentCount

		// TODO： 获取 isFavorite
		//isFavorite

		vs = append(vs, &publish.Video{
			Id:            v.ID,
			Author:        pubAuthor,
			PlayUrl:       v.PlayURL,
			CoverUrl:      v.CoverURL,
			FavoriteCount: 0,     // TODO
			CommentCount:  0,     // TODO
			IsFavorite:    false, // TODO
			Title:         v.Title,
		})
	}

	return vs, nil
}
