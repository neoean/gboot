package qiniu

import (
	"bytes"
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
	"step2money-server/common/config"
)

var (
	formUploader *storage.FormUploader
)

func Init() {
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Region = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	formUploader = storage.NewFormUploader(&cfg)
}

func Upload(data []byte, key string) {
	bucket := "s2m"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := auth.New(config.Config.Qiniu.AccessKey, config.Config.Qiniu.SecretKey)
	upToken := putPolicy.UploadToken(mac)

	dataLen := int64(len(data))

	ret := storage.PutRet{}
	err := formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), dataLen, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
