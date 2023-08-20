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

	//TODO : video_ids -> videos
	//----------------------下面注释的是别人写的-----------------
	// resp, err := rpc.PublishIds2List(s.ctx, &publish.Ids2ListRequest{VideoId: video_ids, UserId: req.UserId, MUserId: req.MUserId})
	// f_video := pack.PublishVideo2FavoriteVideo(resp.VideoList)
	// return f_video, err
	return nil, nil
}
