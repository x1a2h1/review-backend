package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

type KODOUploader struct {
	accessKey string
	secretKey string
	bucket    string
	server    string
}

func NewKODOUploader(cfg interface{}) *KODOUploader {
	// 类型断言
	if kodoConfig, ok := cfg.(struct {
		AccessKey string `mapstructure:"access_key"`
		SecretKey string `mapstructure:"secret_key"`
		Bucket    string `mapstructure:"bucket"`
		Server    string `mapstructure:"server"`
	}); ok {
		return &KODOUploader{
			accessKey: kodoConfig.AccessKey,
			secretKey: kodoConfig.SecretKey,
			bucket:    kodoConfig.Bucket,
			server:    kodoConfig.Server,
		}
	}
	// 默认返回空配置
	return &KODOUploader{}
}

func (k *KODOUploader) Upload(ctx context.Context, input *Input) (*Output, error) {
	// 裁剪文件后缀 - 文件类型
	suffixList := gstr.Split(input.FileName, ".")
	suffixStr := suffixList[len(suffixList)-1]

	putPolicy := storage.PutPolicy{
		Scope: k.bucket,
	}
	client := qbox.NewMac(
		k.accessKey,
		k.secretKey,
	)
	upToken := putPolicy.UploadToken(client)
	cfg := storage.Config{}
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{}
	// 开始上传 - 添加日期文件夹前缀
	dateFolder := time.Now().Format("2006-01/02/")
	key := fmt.Sprintf("%s%s.%s", dateFolder, input.Hash, suffixStr)
	if err := formUploader.PutFile(context.Background(), &ret, upToken, key, input.FilePath, &putExtra); err != nil {
		return nil, err
	}
	// 删除临时文件
	if err := gfile.Remove(input.FilePath); err != nil {
		return nil, err
	}
	return &Output{
		Name:       input.FileName,
		Type:       suffixStr,
		Url:        k.server + "/" + ret.Key,
		BucketHash: ret.Hash,
	}, nil
}
