package service

import (
	"context"

	"github.com/Yra-A/Douyin_Simple_Demo/cmd/favorite/dal/db"
	"github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/favorite"
)

type ActionFavoriteService struct {
	ctx context.Context
}

// NewUploadVideoService new CheckUserService
func NewActionFavoriteService(ctx context.Context) *ActionFavoriteService {
	return &ActionFavoriteService{
		ctx: ctx,
	}
}

func (s *ActionFavoriteService) ActionFavorite(req *favorite.FavoriteActionRequest) error {
	//1点赞
	if req.ActionType == 1 {
		return db.Add(s.ctx, req)
	} else {
		//2取消点赞
		return db.Delete(s.ctx, req)
	}

}
