package service

import (
	"context"

	"github.com/Yra-A/Douyin_Simple_Demo/cmd/favorite/dal/db"
	"github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/favorite"
)

type FavoriteCountService struct {
	ctx context.Context
}

// NewFavoriteCountService new FavoriteCountService
func NewFavoriteCountService(ctx context.Context) *FavoriteCountService {
	return &FavoriteCountService{
		ctx: ctx,
	}
}

func (s *FavoriteCountService) FavoriteCount(req *favorite.FavoriteCountRequest) (int64, error) {
	//拉取
	return db.QueryFavoriteCount(s.ctx, req.VideoId)
}
