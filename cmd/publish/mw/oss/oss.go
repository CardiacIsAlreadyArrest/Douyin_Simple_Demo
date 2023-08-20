package oss

import (
  "bytes"
  conf "github.com/Yra-A/Douyin_Simple_Demo/pkg/configs/oss"
)

// Download 根据文件名，从 OSS 上下载文件
func Download(filename string) (data []byte, err error) {
  objectName := conf.BaseURL + filename
  reader, err := bucket.GetObject(objectName)
  buf := new(bytes.Buffer)
  _, err = buf.ReadFrom(reader)
  if err != nil {
    return nil, err
  }
  data = buf.Bytes()
  return data, nil
}

// UploadFile 上传视频文件
func UploadFile(filename string, data []byte) error {
  reader := bytes.NewReader(data)
  objectName := conf.BaseURL + filename
  err := bucket.PutObject(objectName, reader)
  if err != nil {
    return err
  }
  return nil
}
