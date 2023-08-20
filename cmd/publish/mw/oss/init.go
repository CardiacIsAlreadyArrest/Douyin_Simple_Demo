package oss

import (
	conf "github.com/Yra-A/Douyin_Simple_Demo/pkg/configs/oss"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	client *oss.Client
	bucket *oss.Bucket
)

func Init() {
	// 创建OSSClient实例
	client, err := oss.New(conf.EndPoint, conf.AccessKeyID, conf.AccessKeySecret)
	if err != nil {
		panic(err)
	}
	// 获取存储空间。
	bucket, err = client.Bucket(conf.OssBucketName)
	if err != nil {
		panic(err)
	}

}
