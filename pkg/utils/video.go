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

package utils

import (
	"fmt"
	hpublish "github.com/Yra-A/Douyin_Simple_Demo/cmd/api/biz/model/publish"
	"io"

	"github.com/cloudwego/hertz/pkg/app"
)

// NewVideoName Splicing user_id and time to make unique video name
func NewVideoName(user_id, time int64) string {
	return fmt.Sprintf("%d.%d.mp4", user_id, time)
}

// NewImageName Splicing user_id and time to make unique Image name
func NewImageName(user_id, time int64) string {
	return fmt.Sprintf("%d.%d.png", user_id, time)
}

//// URLconvert Convert the path in the database into a complete url accessible by the front end
//func URLconvert(ctx context.Context, c *app.RequestContext, path string) (fullURL string) {
//	if len(path) == 0 {
//		return ""
//	}
//	arr := strings.Split(path, "/")
//	u, err := minio.GetObjURL(ctx, arr[0], arr[1])
//	if err != nil {
//		hlog.CtxInfof(ctx, err.Error())
//		return ""
//	}
//	u.Scheme = string(c.URI().Scheme())
//	u.Host = string(c.URI().Host())
//	u.Path = "/src" + u.Path
//	return u.String()
//}

func ParsePubActionRequest(req *hpublish.PublishActionRequest, c *app.RequestContext) error {
	file, err := c.FormFile("data")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	bytes, err := io.ReadAll(src)
	if err != nil {
		return err
	}
	req.Data = bytes
	req.Token = c.PostForm("token")
	req.Title = c.PostForm("title")
	return nil
}
