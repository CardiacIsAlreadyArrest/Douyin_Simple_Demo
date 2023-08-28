package service

import (
	"context"
	"github.com/Yra-A/Douyin_Simple_Demo/cmd/relation/dal/db"
	"github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/relation"
)

type RelationService struct {
	ctx context.Context
}

// RelationAction create a relation with 2 users
func (s *RelationService) RelationAction(req *relation.RelationActionRequest) error {
	//var userId int64
	//if tempUserId, exist := c.Get("current_user_id"); exist {
	//	userId = tempUserId.(int64)
	//} else {
	//	handler.BadResponse(c, errno.AuthorizationFailedErr)
	//	return nil
	//}
	userId := req.UserId
	toUserId := req.ToUserId
	actionType := req.ActionType
	// they are friends
	if actionType == 1 {
		return db.CreateRelation(s.ctx, &db.Relation{
			UserId:   userId,
			ToUserId: toUserId,
		})
	}
	// they are not friends
	return db.DeleteRelation(s.ctx, userId, toUserId)
}

func NewRelationActionService(ctx context.Context) *RelationService {
	return &RelationService{
		ctx: ctx,
	}
}
