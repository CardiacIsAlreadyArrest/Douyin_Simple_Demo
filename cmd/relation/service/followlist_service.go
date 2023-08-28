package service

import (
	"context"
	"github.com/Yra-A/Douyin_Simple_Demo/cmd/publish/rpc"
	"github.com/Yra-A/Douyin_Simple_Demo/cmd/relation/dal/db"
	"github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/relation"
	"github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/user"
)

type RelationFollowListService struct {
	ctx context.Context
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
