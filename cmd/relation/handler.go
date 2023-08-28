package main

import (
	"context"
	"github.com/Yra-A/Douyin_Simple_Demo/cmd/relation/service"
	relation "github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/relation"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// RelationAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationAction(ctx context.Context, req *relation.RelationActionRequest) (resp *relation.RelationActionResponse, err error) {
	resp = new(relation.RelationActionResponse)

	invalidMsg := "invalid request"
	wrongMsg := "relation action failed"
	correctMsg := "relation action success"

	if req.ToUserId < 0 || req.ActionType < 0 || req.ActionType > 2 {
		resp.StatusCode = -1
		resp.StatusMsg = &invalidMsg
		return
	}
	err = service.NewRelationActionService(ctx).RelationAction(req)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = &wrongMsg
		return
	}
	resp.StatusCode = 0
	resp.StatusMsg = &correctMsg
	return
}

// RelationFollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowList(ctx context.Context, req *relation.RelationFollowerListRequest) (resp *relation.RelationFollowListResponse, err error) {
	resp = new(relation.RelationFollowListResponse)

	invalidMsg := "invalid request"
	failMsg := "relation follow list failed"
	correctMsg := "relation follow list success"
	// 没有找到用户和用户的关注列表
	if req.UserId < 0 || req.MUserId < 0 {
		resp.StatusCode = -1
		resp.StatusMsg = &invalidMsg
		return resp, nil
	}

	var followList []*relation.User
	followList, err = service.NewRelationFollowerListService(ctx).RelationFollowerList(req)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = &failMsg
		return resp, nil
	}

	resp.StatusCode = 0
	resp.StatusMsg = &correctMsg
	resp.UserList = followList
	return resp, nil
}

// RelationFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowerList(ctx context.Context, req *relation.RelationFollowerListRequest) (resp *relation.RelationFollowerListResponse, err error) {
	resp = new(relation.RelationFollowerListResponse)

	invalidMsg := "invalid request"
	failMsg := "Not Found your follower list"
	correctMsg := "Search your follower list success"

	if req.UserId < 0 || req.MUserId < 0 {
		resp.StatusCode = -1
		resp.StatusMsg = &invalidMsg
		return resp, nil
	}
	var users []*relation.User
	users, err = service.NewRelationFollowerListService(ctx).RelationFollowerList(req)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = &failMsg
		return resp, nil
	}

	resp.StatusCode = 0
	resp.StatusMsg = &correctMsg
	resp.UserList = users

	return resp, nil
}

// RelationFriendList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFriendList(ctx context.Context, req *relation.RelationFriendListRequest) (resp *relation.RelationFriendListResponse, err error) {
	resp = new(relation.RelationFriendListResponse)

	inValidMsg := "Relation Friendlist request inValid"
	failMsg := "Not Found your friend list"
	successMsg := "Search your friend list success"

	if req.UserId < 0 || req.MUserId < 0 {
		resp.StatusCode = -1
		resp.StatusMsg = &inValidMsg
		return resp, nil
	}

	var users []*relation.User
	users, err = service.NewRelationFriendListService(ctx).RelationFriendList(req)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = &failMsg
		return resp, nil
	}

	resp.StatusCode = 0
	resp.StatusMsg = &successMsg
	resp.UserList = users

	return resp, nil
}
