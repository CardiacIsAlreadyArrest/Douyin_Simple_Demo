package db

import (
	"context"
	"fmt"
	"time"

	"github.com/Yra-A/Douyin_Simple_Demo/cmd/publish/dal/db"
	"github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/favorite"
	"github.com/Yra-A/Douyin_Simple_Demo/pkg/constants"
	"gorm.io/gorm"
)

type Favorite struct {
	gorm.Model
	// Id int64 `gorm:"primaryKey;autoIncrement" json:id`
	UserId  int64 `json:"user_id"`
	VideoId int64 `json:"video_id"`
}

func (u *Favorite) TableName() string {
	return constants.FavoriteTableName
}

type Video struct {
	gorm.Model
	ID          int64
	AuthorID    int64
	PlayURL     string
	CoverURL    string
	PublishTime time.Time
	Title       string
}

func (Video) TableName() string {
	return constants.VideosTableName
}

func Delete(ctx context.Context, req *favorite.FavoriteActionRequest) error {
	return DB.WithContext(ctx).Where("user_id = ? and video_id = ?", req.UserId, req.VideoId).Delete(&Favorite{}).Error
}

// add favorite
func Add(ctx context.Context, req *favorite.FavoriteActionRequest) error {
	var res []*Favorite
	//如果存在，不再增加
	if err := DB.WithContext(ctx).Where("user_id = ? and video_id = ?", req.UserId, req.VideoId).Find(&res).Error; err != nil {
		fmt.Println("如果存在，不再增加")
		return nil
	}

	if len(res) != 0 {
		fmt.Println("db.Add 如果存在，不再增加", req.VideoId)
		return nil
	}

	//如果不存在，增加
	u := new(Favorite)
	u.UserId = req.UserId
	u.VideoId = req.VideoId
	if err := DB.WithContext(ctx).Create(u).Error; err != nil {
		fmt.Println("db.Add 增加失败 UserId   VideoId", req.UserId, req.VideoId)
		return err
	}
	fmt.Println("db.Add 增加成功 UserId   VideoId", req.UserId, req.VideoId)
	return nil
}

// getlist
func QueryUsr(ctx context.Context, usrid int64) ([]int64, error) {
	res := make([]*Favorite, 0)
	if err := DB.WithContext(ctx).Where("user_id = ?", usrid).Find(&res).Error; err != nil {
		return nil, err
	}

	var video_id []int64
	for _, users := range res {
		video_id = append(video_id, users.VideoId)
	}
	fmt.Print("user_id:", usrid)
	for _, vid := range video_id {
		fmt.Print("  video_id:", vid)
	}
	fmt.Println()
	return video_id, nil
}

func GetVideoListByVideoIDList(ctx context.Context, video_id_list []int64) ([]*db.Video, error) {
	var video_list []*db.Video
	var err error
	for _, item := range video_id_list {
		var video *db.Video
		err = DB.WithContext(ctx).Where("id = ?", item).Find(&video).Error
		if err != nil {
			return video_list, err
		}
		video_list = append(video_list, video)
	}

	return video_list, err
}

// favorite_count  videoid  how many people like
func QueryFavoriteCount(ctx context.Context, video_id int64) (int64, error) {
	var favorite_count int64
	favorite_count = 0
	res := make([]*Favorite, 0)
	if err := DB.WithContext(ctx).Where("video_id = ?", video_id).Find(&res).Error; err != nil {
		fmt.Println("favorite_count:", favorite_count)
		return favorite_count, err
	}

	favorite_count = int64(len(res))
	fmt.Println("favorite_count:", favorite_count)
	return favorite_count, nil
}

// is_favorite   ueser_id like video_id
func QueryIsFavorite(ctx context.Context, req *favorite.IsFavoriteRequest) (bool, error) {
	res := make([]*Favorite, 0)
	DB.WithContext(ctx).Where("user_id = ? and video_id = ?", req.UserId, req.VideoId).Find(&res)

	if len(res) == 0 {
		fmt.Println("user like this video  UserId   VideoId", req.UserId, req.VideoId)
		return false, nil
	}

	fmt.Println("user do not like this video")
	return true, nil
}
