/*
 * Copyright 2023 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package db

import (
	"gorm.io/gorm"
	"time"

	"github.com/Yra-A/Douyin_Simple_Demo/pkg/constants"
)

type Video struct {
	gorm.Model
	ID          int64
	AuthorID    int64
	PlayURL     string
	CoverURL    string
	PublishTime time.Time
	Title       string
}

func (Video) TableName() string {
	return constants.VideosTableName
}

func CreateVideo(video *Video) (int64, error) {
	err := DB.Create(video).Error
	if err != nil {
		return 0, err
	}
	return video.ID, err
}

func GetVideoByUserId(user_id int64) ([]*Video, error) {
	var videos []*Video
	err := DB.Where("author_id = ?", user_id).Find(&videos).Error
	if err != nil {
		return videos, err
	}
	return videos, err
}

func GetVideoListByVideoIDList(videoIds []int64) ([]*Video, error) {
	var video_list []*Video
	var err error
	for _, id := range videoIds {
		var video *Video

		if err := DB.Where("id = ?", id).Find(&video).Error; err != nil {
			return video_list, err
		}
		video_list = append(video_list, video)
	}
	return video_list, err
}
