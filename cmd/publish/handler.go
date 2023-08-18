package main

import (
	"context"
	publish "github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/publish"
	"github.com/cloudwego/kitex/pkg/klog"
)

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

// PublishAction implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishAction(ctx context.Context, req *publish.PublishActionRequest) (resp *publish.PublishActionResponse, err error) {
	klog.CtxDebugf(ctx, "PublishAction called: %s")
	return
}

// PublishList implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishList(ctx context.Context, req *publish.PublishListRequest) (resp *publish.PublishListResponse, err error) {
	// TODO: Your code here...
	return
}
