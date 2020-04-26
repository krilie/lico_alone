package third_api

import (
	"context"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"io"
)

type OssQiNiu struct {
	AccessKey  string
	SecretKey  string
	BucketName string
	qboxMac    *qbox.Mac
}

func (o *OssQiNiu) UploadFile(ctx context.Context, name string, file io.ReadSeeker, size int64) (content, bucket, key string, err error) {
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret.Key, ret.Hash)
	panic("implement me")
}

func (o *OssQiNiu) DeleteFile(ctx context.Context, userId, key string) error {
	panic("implement me")
}

func (o *OssQiNiu) GetFullUrl(ctx context.Context, isPub bool, key string) (url string) {
	panic("implement me")
}

func (o *OssQiNiu) GetBaseUrl(ctx context.Context) string {
	panic("implement me")
}

func (o *OssQiNiu) GetBucketName(ctx context.Context) string {
	panic("implement me")
}
