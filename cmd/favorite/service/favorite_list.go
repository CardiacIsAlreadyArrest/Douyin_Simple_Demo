package service

import (
	"context"
	"log"

	"github.com/Yra-A/Douyin_Simple_Demo/cmd/favorite/dal/db"
	"github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/favorite"
)

type FavoriteListService struct {
	ctx context.Context
}

// NewFavoriteListService new FavoriteListService
func NewFavoriteListService(ctx context.Context) *FavoriteListService {
	return &FavoriteListService{
		ctx: ctx,
	}
}

// getlist
func (s *FavoriteListService) FavoriteList(req *favorite.FavoriteListRequest) ([]*favorite.Video, error) {
	//本用户id + video_id[]  获取 video_list
	video_ids, _ := db.QueryUsr(s.ctx, req.UserId)

	if len(video_ids) == 0 {
		log.Println("FavoriteList : video_ids is blank")
	}

	// temp, err := rpc.PublishIds2List(s.ctx, &publish.PublishListRequest{UserId: req.UserId, Token: req.Token})
	// temp, err := db.GetVideoListByVideoIDList(s.ctx, video_ids)
	var resp []*favorite.Video
	// for _, a := range temp {
	// 	b := &favorite.Video{Id: a.ID, Author: (*favorite.User)(a.AuthorID), PlayUrl: a.PlayURL, CoverUrl: a.CoverURL, FavoriteCount: a., IsFavorite: a.IsFavorite, Title: a.Title}
	// 	resp = append(resp, b)
	// }
	// todo:等一个User的id->user和publish的videoID->video

	return resp, nil
}
