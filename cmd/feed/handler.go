package main

import (
	"context"
	"github.com/Yra-A/Douyin_Simple_Demo/cmd/feed/service"
	feed "github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/feed"
	"github.com/Yra-A/Douyin_Simple_Demo/pkg/errno"
	"github.com/cloudwego/kitex/pkg/klog"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// Feed implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) Feed(ctx context.Context, req *feed.FeedRequest) (resp *feed.FeedResponse, err error) {
	klog.CtxDebugf(ctx, "【Feed called】Latest Time is %v", req.LatestTime)

	videos, nextTime, err := service.NewFeedService(ctx).Feed(req)
	if err != nil {
		ErrMsg := err.Error()
		resp = &feed.FeedResponse{
			StatusCode: errno.FalseCode,
			StatusMsg:  &ErrMsg,
		}
		return resp, nil
	}

	return &feed.FeedResponse{
		StatusCode: errno.SuccessCode,
		StatusMsg:  &errno.Success.ErrMsg,
		VideoList:  videos,
		NextTime:   nextTime,
	}, nil
}
