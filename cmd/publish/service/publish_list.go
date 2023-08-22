package service

import (
	"context"
	"github.com/Yra-A/Douyin_Simple_Demo/cmd/publish/dal/db"
	"github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/publish"
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
		favoriteCount, commentCount, isFavorite, err := getOtherVideoInfo(v.ID)
		if err != nil {
			return nil, err
		}
		author, err := getAuthorByUserId(v.AuthorID)
		if err != nil {
			return nil, err
		}

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
func getOtherVideoInfo(videoID int64) (int64, int64, bool, error) {
	return 0, 0, false, nil
}

// TODO: 实现 rpc.GetUser
func getAuthorByUserId(userId int64) (*publish.User, error) {
	var author *publish.User
	var err error

	//userAuthor, err = rpc.GetUser(context.Background(), &user.UserInfoRequest{
	//	UserId: v.AuthorID,
	//	Token:  req.Token,
	//})
	return author, err
}
