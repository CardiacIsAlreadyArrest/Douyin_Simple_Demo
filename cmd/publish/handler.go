package main

import (
  "context"
  service "github.com/Yra-A/Douyin_Simple_Demo/cmd/publish/service"
  publish "github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/publish"
  "github.com/Yra-A/Douyin_Simple_Demo/pkg/constants"
  "github.com/Yra-A/Douyin_Simple_Demo/pkg/errno"
  "github.com/cloudwego/kitex/pkg/klog"
)

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

// PublishAction implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishAction(ctx context.Context, req *publish.PublishActionRequest) (resp *publish.PublishActionResponse, err error) {
  klog.CtxDebugf(ctx, "【PublishAction called】user id: %v; title: %v; data length: %v", req.UserId, req.Title, len(req.Data))

  if len(req.Data) == 0 || len(req.Title) == 0 {
    resp = &publish.PublishActionResponse{
      StatusCode: errno.ParamErrCode,
      StatusMsg:  &errno.ParamErr.ErrMsg,
    }
    return resp, nil
  }

  if int64(len(req.Data)) > constants.MaxVideoSize {
    resp = &publish.PublishActionResponse{
      StatusCode: errno.VideoExceedMaxSizeErrCode,
      StatusMsg:  &errno.VideoExceedMaxSizeErr.ErrMsg,
    }
    return resp, nil
  }

  err = service.NewUploadVideoService(ctx).UploadVideo(req)
  if err != nil {
    ErrMsg := err.Error()
    resp = &publish.PublishActionResponse{
      StatusCode: errno.FalseCode,
      StatusMsg:  &ErrMsg,
    }
    return resp, nil
  }

  resp = &publish.PublishActionResponse{
    StatusCode: errno.SuccessCode,
    StatusMsg:  &errno.Success.ErrMsg,
  }
  return resp, nil
}

// PublishList implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishList(ctx context.Context, req *publish.PublishListRequest) (resp *publish.PublishListResponse, err error) {
  klog.CtxDebugf(ctx, "【PublishAction called】user id: %v", req.UserId)

  if req.UserId == 0 {
    resp = &publish.PublishListResponse{
      StatusCode: errno.AuthorizationFailedErrCode,
      StatusMsg:  &errno.AuthorizationFailedErr.ErrMsg,
    }
    return resp, nil
  }

  videos, err := service.NewPublishListService(ctx).PublishList(req)
  if err != nil {
    ErrMsg := err.Error()
    resp = &publish.PublishListResponse{
      StatusCode: errno.FalseCode,
      StatusMsg:  &ErrMsg,
    }
    return resp, nil
  }

  resp = &publish.PublishListResponse{
    StatusCode: errno.SuccessCode,
    StatusMsg:  &errno.Success.ErrMsg,
    VideoList:  videos,
  }
  return resp, nil
}

// GetVideoList implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) GetVideoList(ctx context.Context, req *publish.GetVideoListRequest) (resp *publish.GetVideoListResponse, err error) {
  klog.CtxDebugf(ctx, "【GetVideoList called】")

  if req.UserId == 0 {
    resp = &publish.GetVideoListResponse{
      StatusCode: errno.AuthorizationFailedErrCode,
      StatusMsg:  &errno.AuthorizationFailedErr.ErrMsg,
    }
    return resp, nil
  }

  videos, err := service.NewGetVideoListService(ctx).GetVideoList(req)
  if err != nil {
    ErrMsg := err.Error()
    resp = &publish.GetVideoListResponse{
      StatusCode: errno.FalseCode,
      StatusMsg:  &ErrMsg,
    }
    return resp, nil
  }

  resp = &publish.GetVideoListResponse{
    StatusCode: errno.SuccessCode,
    StatusMsg:  &errno.Success.ErrMsg,
    VideoList:  videos,
  }
  return resp, nil
}
