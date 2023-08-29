package db

import (
	"context"
	"errors"
	"github.com/Yra-A/Douyin_Simple_Demo/pkg/constants"
	"gorm.io/gorm"
)

type Relation struct {
	gorm.Model
	UserId   int64
	ToUserId int64
}

type User struct {
	gorm.Model
}

func (r *Relation) TableName() string {
	return constants.RelationTableName
}

// CreateRelation create a relation db with 2 users
func CreateRelation(ctx context.Context, relations *Relation) error {
	if err := DB.WithContext(ctx).Where("user_id = ? AND to_user_id = ?", relations.UserId, relations.ToUserId).
		First(&Relation{}).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}
	return DB.WithContext(ctx).Create(relations).Error
}

func DeleteRelation(ctx context.Context, userId, toUserId int64) error {
	return DB.WithContext(ctx).Where("user_id = ? AND to_user_id = ?", userId, toUserId).
		Delete(&Relation{}).Error
}

// QueryFollowList query followed users' id
func QueryFollowList(ctx context.Context, userId int64) ([]int64, error) {
	var relationFound []*Relation

	var resp []int64

	if err := DB.WithContext(ctx).Where("user_id = ?", userId).Find(&relationFound).Error; err != nil {
		return resp, err
	}

	for _, r := range relationFound {
		resp = append(resp, r.ToUserId)
	}

	return resp, nil
}

// QueryFollowerList query users who followed userId
func QueryFollowerList(ctx context.Context, userId int64) ([]int64, error) {
	var relationFound []*Relation

	var resp []int64

	if err := DB.WithContext(ctx).Where("to_user_id = ?", userId).Find(&relationFound).Error; err != nil {
		return resp, err
	}

	for _, r := range relationFound {
		resp = append(resp, r.UserId)
	}

	return resp, nil
}

// QueryFriendList query users who followed userId and userId followed him/her
func QueryFriendList(ctx context.Context, userId int64) ([]int64, error) {
	var resp []int64

	var relationFound []*Relation
	if err := DB.WithContext(ctx).Where("to_user_id = ?", userId).Find(&relationFound).Error; err != nil {
		return resp, err
	}

	for _, r := range relationFound {
		var friendFound *Relation
		if err := DB.WithContext(ctx).Where("to_user_id = ? and user_id = ?", r.UserId, userId).Find(&friendFound).Error; err != nil {
			return resp, err
		}
		if friendFound.ToUserId > 0 {
			resp = append(resp, friendFound.ToUserId)
		}
	}

	return resp, nil
}

// QueryFollowCount return number of users who are followed by userId
func QueryFollowCount(ctx context.Context, userId int64) (int64, error) {
	var relationFound []*Relation

	var resp int64

	if err := DB.WithContext(ctx).Where("user_id = ?", userId).Find(&relationFound).Error; err != nil {
		return resp, err
	}

	resp = int64(len(relationFound))

	return resp, nil
}

// QueryFollowerCount return number of users who followed userId
func QueryFollowerCount(ctx context.Context, userId int64) (int64, error) {
	var relationFound []*Relation

	var resp int64

	if err := DB.WithContext(ctx).Where("to_user_id = ?", userId).Find(&relationFound).Error; err != nil {
		return resp, err
	}

	resp = int64(len(relationFound))

	return resp, nil
}

// QueryIsFollow return whether userId followed toUserId
func QueryIsFollow(ctx context.Context, userId int64, toUserId int64) (bool, error) {
	var relationFound *Relation

	var resp bool

	if err := DB.WithContext(ctx).Where("user_id = ? and to_user_id = ?", userId, toUserId).First(&relationFound).Error; err != nil {
		resp = false
		return resp, nil
	}

	resp = true

	return resp, nil
}
