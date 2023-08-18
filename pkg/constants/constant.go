// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package constants

// minio
const (
	MinioEndPoint        = "localhost:18001" // 对象存储的 URL
	MinioAccessKeyID     = "douyin"          // 账户 ID
	MinioSecretAccessKey = "douyin111"       // 账户密码
	MiniouseSSL          = false             // 不是用 https

	MinioVideoBucketName = "videobucket"
	MinioImgBucketName   = "imagebucket"
)

const (
	UserTableName = "users"

	SecretKey   = "secret key"
	IdentityKey = "user id"

	UserServiceName    = "user"
	PublishServiceName = "publish"

	CPURateLimit float64 = 80.0

	MaxVideoSize int64 = 128 * 1024 * 1024 // 可上传的单个视频大小最大为 128 MB
)

var (
	MySQLDefaultDSN = "gorm:gorm@tcp(localhost:18000)/test_douyin?charset=utf8&parseTime=True&loc=Local"
	EtcdAddress     = "localhost:2379"
)
