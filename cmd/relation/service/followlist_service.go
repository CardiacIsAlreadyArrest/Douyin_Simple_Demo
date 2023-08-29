package service

import (
	"context"
	"github.com/Yra-A/Douyin_Simple_Demo/cmd/relation/dal/db"
	"github.com/Yra-A/Douyin_Simple_Demo/cmd/relation/rpc"
	"github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/relation"
	"github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/user"
)

type RelationFollowListService struct {
	ctx context.Context
}

func (s *RelationFollowListService) RelationFollowList(req *relation.RelationFollowListRequest) ([]*relation.User, error) {
	var resp []*relation.User

	userID := req.UserId

	followIDs, err := db.QueryFollowList(s.ctx, userID)
	if err != nil {
		return resp, err
	}

	for _, id := range followIDs {
		u, err := rpc.GetUser(s.ctx, &user.UserInfoRequest{
			UserId: id,
			//MUserId: req.MUserId,
		})
		if err != nil {
			return resp, err
		}
		resp = append(resp, &relation.User{
			Id:            u.Id,
			Name:          u.Name,
			FollowCount:   &u.FollowCount,
			FollowerCount: &u.FollowerCount,
			IsFollow:      u.IsFollow,
		})
	}

	return resp, nil
}

// NewRelationFollowListService new RelationAction
func NewRelationFollowListService(ctx context.Context) *RelationFollowListService {
	return &RelationFollowListService{ctx: ctx}
}

func NewRelationFollowerListService(ctx context.Context) *RelationService {
	return &RelationService{
		ctx: ctx,
	}
}

func (s *RelationService) RelationFollowerList(req *relation.RelationFollowerListRequest) ([]*relation.User, error) {
	var resp []*relation.User

	userID := req.UserId

	followerIDs, err := db.QueryFollowerList(s.ctx, userID)
	if err != nil {
		return resp, err
	}

	for _, id := range followerIDs {
		u, err := rpc.GetUser(s.ctx, &user.UserInfoRequest{
			UserId: id,
			//MUserId: req.MUserId,
		})
		if err != nil {
			return resp, err
		}
		resp = append(resp, &relation.User{
			Id:            u.Id,
			Name:          u.Name,
			FollowCount:   &u.FollowCount,
			FollowerCount: &u.FollowerCount,
			IsFollow:      u.IsFollow,
		})
	}

	return resp, nil
}

type RelationFriendListService struct {
	ctx context.Context
}

func NewRelationFriendListService(ctx context.Context) *RelationFriendListService {
	return &RelationFriendListService{ctx: ctx}
}

// RelationFriendList return users who followed each other
func (s *RelationFriendListService) RelationFriendList(req *relation.RelationFriendListRequest) ([]*relation.User, error) {
	var resp []*relation.User

	userID := req.UserId

	friendIDs, err := db.QueryFriendList(s.ctx, userID)
	if err != nil {
		return resp, err
	}

	for _, id := range friendIDs {
		u, err := rpc.GetUser(s.ctx, &user.UserInfoRequest{
			UserId: id,
			//MUserId: req.MUserId,
		})
		if err != nil {
			return resp, err
		}
		resp = append(resp, &relation.User{
			Id:            u.Id,
			Name:          u.Name,
			FollowCount:   &u.FollowCount,
			FollowerCount: &u.FollowerCount,
			IsFollow:      u.IsFollow,
		})
	}

	return resp, nil
}

type RelationFollowerCountService struct {
	ctx context.Context
}

// NewRelationFollowerCountService new RelationFollowerCountService
func NewRelationFollowerCountService(ctx context.Context) *RelationFollowerCountService {
	return &RelationFollowerCountService{ctx: ctx}
}

// RelationFollowerCount return number of users who are followed
func (s *RelationFollowerCountService) RelationFollowerCount(req *relation.RelationFollowCountRequest) (int64, error) {
	var resp int64

	userID := req.UserId

	followerCount, err := db.QueryFollowerCount(s.ctx, userID)
	if err != nil {
		return resp, err
	}

	resp = followerCount

	return resp, nil
}

type RelationFollowCountService struct {
	ctx context.Context
}

// NewRelationFollowCountService new RelationFollowCountService
func NewRelationFollowCountService(ctx context.Context) *RelationFollowCountService {
	return &RelationFollowCountService{ctx: ctx}
}

// RelationFollowCount return number of users who followed each other
func (s *RelationFollowCountService) RelationFollowCount(req *relation.RelationFollowCountRequest) (int64, error) {
	var resp int64

	userID := req.UserId

	followCount, err := db.QueryFollowCount(s.ctx, userID)
	if err != nil {
		return resp, err
	}

	resp = followCount

	return resp, nil
}

// NewRelationIsFollowService new RelationIsFollowService
func NewRelationIsFollowService(ctx context.Context) *RelationIsFollowService {
	return &RelationIsFollowService{ctx: ctx}
}

type RelationIsFollowService struct {
	ctx context.Context
}

// RelationFollowCount return number of users who followed each other
func (s *RelationIsFollowService) RelationIsFollow(req *relation.RelationIsFollowRequest) (bool, error) {
	var resp bool

	userId := req.UserId
	toUserId := req.ToUserId

	isFollow, err := db.QueryIsFollow(s.ctx, userId, toUserId)
	if err != nil {
		return resp, err
	}

	resp = isFollow

	return resp, nil
}
