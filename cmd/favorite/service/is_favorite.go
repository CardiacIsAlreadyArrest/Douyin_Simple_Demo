package service

import (
  "context"

  "github.com/Yra-A/Douyin_Simple_Demo/cmd/favorite/dal/db"
  "github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/favorite"
)

type IsFavoriteService struct {
  ctx context.Context
}

// NewIsFavoriteService new IsFavoriteService
func NewIsFavoriteService(ctx context.Context) *IsFavoriteService {
  return &IsFavoriteService{
    ctx: ctx,
  }
}

func (s *IsFavoriteService) IsFavorite(req *favorite.IsFavoriteRequest) (bool, error) {
  return db.QueryIsFavorite(s.ctx, req)
}
