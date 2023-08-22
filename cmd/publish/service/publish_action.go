package service

import (
	"context"
	"fmt"
	"github.com/Yra-A/Douyin_Simple_Demo/cmd/publish/dal/db"
	"github.com/Yra-A/Douyin_Simple_Demo/cmd/publish/mw/ffmpeg"
	"github.com/Yra-A/Douyin_Simple_Demo/cmd/publish/mw/oss"
	"github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/publish"
	conf "github.com/Yra-A/Douyin_Simple_Demo/pkg/configs/oss"
	"github.com/Yra-A/Douyin_Simple_Demo/pkg/utils"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"strconv"
	"time"
)

type UploadVideoService struct {
	ctx context.Context
}

// NewUploadVideoService new CheckUserService
func NewUploadVideoService(ctx context.Context) *UploadVideoService {
	return &UploadVideoService{
		ctx: ctx,
	}
}

func (s *UploadVideoService) UploadVideo(req *publish.PublishActionRequest) error {
	title := req.Title
	user_id := req.UserId
	videoBytes := req.Data
	nowTime := time.Now()

	// TODO：支持自定义扩展名，现在默认为 .mp4 和 .png
	videoName := utils.NewVideoName(user_id, nowTime.Unix()) // user_id.nowTime.Unix().mp4
	imageName := utils.NewImageName(user_id, nowTime.Unix()) // user_id.nowTime.Unix().png

	/*------------------------------------------------------*/
	/*------------------上传视频到 minio 服务------------------*/
	/*------------------------------------------------------*/
	err := oss.UploadFile(videoName, videoBytes)
	hlog.CtxInfof(s.ctx, "视频上传大小为: "+strconv.FormatInt(int64(len(videoBytes)), 10)+"B")
	if err != nil {
		hlog.CtxInfof(s.ctx, "上传视频出现错误: "+err.Error())
	}
	PlayURL := conf.PublicURL + videoName
	if err != nil {
		hlog.CtxInfof(s.ctx, "获取视频 URL 出现错误: "+err.Error())
	}
	fmt.Println("视频 URL 为 " + PlayURL)

	/*---------------------------------------------------------------------------------*/
	/*------------------先用 ffmpeg 获取视频封面，再上传视频封面到 minio 服务------------------*/
	/*---------------------------------------------------------------------------------*/
	imageBuf, err := ffmpeg.GetSnapshot(PlayURL)

	err = oss.UploadFile(imageName, imageBuf.Bytes())
	hlog.CtxInfof(s.ctx, "封面上传大小为:"+strconv.FormatInt(int64(imageBuf.Len()), 10)+"B")
	if err != nil {
		hlog.CtxInfof(s.ctx, "上传封面出现错误: "+err.Error())
	}
	CoverURL := conf.PublicURL + imageName
	if err != nil {
		hlog.CtxInfof(s.ctx, "获取封面 URL 出现错误: "+err.Error())
	}
	fmt.Println("封面 URL 为 " + CoverURL)

	// 将相关内容上传到 MySQL 数据库
	_, err = db.CreateVideo(&db.Video{
		AuthorID:    user_id,
		PlayURL:     PlayURL,
		CoverURL:    CoverURL,
		PublishTime: nowTime,
		Title:       title,
	})
	return err
}
