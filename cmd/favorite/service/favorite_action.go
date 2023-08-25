package service

import (
	"context"

	"github.com/Yra-A/Douyin_Simple_Demo/cmd/favorite/dal/db"
	"github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/favorite"
)

type FavoriteActionService struct {
	ctx context.Context
}

// NewUploadVideoService new CheckUserService
func NewFavoriteActionService(ctx context.Context) *FavoriteActionService {
	return &FavoriteActionService{
		ctx: ctx,
	}
}

func (s *FavoriteActionService) FavoriteAction(req *favorite.FavoriteActionRequest) error {
	//1点赞
	if req.ActionType == 1 {
		return db.Add(s.ctx, req)
	} else {
		//2取消点赞
		return db.Delete(s.ctx, req)
	}

}
