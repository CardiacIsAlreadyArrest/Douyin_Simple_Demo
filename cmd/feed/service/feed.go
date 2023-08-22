package service

import (
  "context"
  "fmt"
  "github.com/Yra-A/Douyin_Simple_Demo/cmd/feed/dal/db"
  "github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/feed"
  "github.com/Yra-A/Douyin_Simple_Demo/pkg/constants"
  "time"
)

type FeedService struct {
  ctx context.Context
}

func NewFeedService(ctx context.Context) *FeedService {
  return &FeedService{
    ctx: ctx,
  }
}

func (s *FeedService) Feed(req *feed.FeedRequest) ([]*feed.Video, int64, error) {
  var lastTime time.Time
  if req.LatestTime == 0 {
    lastTime = time.Now()
  } else {
    lastTime = time.Unix(req.LatestTime/1000, 0)
  }
  fmt.Printf("视频的最新投稿时间戳为 %v\n", lastTime)

  dbVideos, err := db.GetVideosByLastTime(lastTime)
  if err != nil {
    return nil, 0, err
  }

  videos := make([]*feed.Video, 0, constants.MaxFeedCount)
  var nextTime int64
  if len(dbVideos) > 0 {
    nextTime = dbVideos[len(dbVideos)-1].PublishTime.Unix()
  }

  for _, v := range dbVideos {
    favoriteCount, commentCount, isFavorite, err := getOtherVideoInfo()
    if err != nil {
      return nil, 0, err
    }
    author, err := getAuthorByUserId(v.AuthorID)
    if err != nil {
      return nil, 0, err
    }
    nv := &feed.Video{
      Id:            v.ID,
      Author:        author,
      PlayUrl:       v.PlayURL,
      CoverUrl:      v.CoverURL,
      Title:         v.Title,
      FavoriteCount: favoriteCount,
      CommentCount:  commentCount,
      IsFavorite:    isFavorite,
    }
    videos = append(videos, nv)
  }

  return videos, nextTime, nil
}

// TODO: 实现获取视频相关信息 (并发获取?)
func getOtherVideoInfo() (int64, int64, bool, error) {
  return 0, 0, false, nil
}

// TODO: 实现 rpc.GetUser
func getAuthorByUserId(userId int64) (*feed.User, error) {
  var author *feed.User
  var err error

  //userAuthor, err = rpc.GetUser(context.Background(), &user.UserInfoRequest{
  //	UserId: v.AuthorID,
  //	Token:  req.Token,
  //})
  return author, err
}
