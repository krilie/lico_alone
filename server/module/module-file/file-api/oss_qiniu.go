package file_api

import (
	"context"
	"github.com/juju/ratelimit"
	"github.com/krilie/lico_alone/component/ncfg"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"io"
	url2 "net/url"
	"strings"
	"time"
)

type OssQiNiu struct {
	AccessKey  string
	SecretKey  string
	BucketName string
	BaseUrl    string
	qboxMac    *qbox.Mac
}

func (o *OssQiNiu) UploadFile(ctx context.Context, fileName string, fileStream io.ReadSeeker, fileSize int64) (url, key string, err error) {
	putPolicy := storage.PutPolicy{
		Scope: o.BucketName,
	}
	upToken := putPolicy.UploadToken(o.qboxMac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = true
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	// 限制速度 150kb/s
	bucket := ratelimit.NewBucketWithRate(150*1024, 150*1024)
	limitReader := ratelimit.Reader(fileStream, bucket)
	// 上传
	err = formUploader.Put(ctx, &ret, upToken, fileName, limitReader, fileSize, &putExtra)
	if err != nil {
		return "", "", err
	}
	url, _ = o.GetUrl(ctx, true, ret.Key)
	return url, ret.Key, nil
}

func (o *OssQiNiu) DeleteFile(ctx context.Context, fileKey string) error {
	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: true,
	}
	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	//cfg.Zone=&storage.ZoneHuabei
	bucketManager := storage.NewBucketManager(o.qboxMac, &cfg)
	err := bucketManager.Delete(o.BucketName, fileKey)
	if err != nil {
		return err
	}
	return nil
}

func (o *OssQiNiu) GetUrl(ctx context.Context, isPub bool, fileKey string) (url string, err error) {
	if isPub {
		publicAccessURL := storage.MakePublicURL(o.BaseUrl, fileKey)
		return publicAccessURL, nil
	} else {
		deadline := time.Now().Add(time.Second * 60 * 20).Unix() // 20分钟有效期
		privateAccessURL := storage.MakePrivateURL(o.qboxMac, o.BaseUrl, fileKey, deadline)
		return privateAccessURL, nil
	}
}

func (o *OssQiNiu) DeleteFileByUrl(ctx context.Context, url string) error {
	return o.DeleteFile(ctx, o.GetKeyByUrl(ctx, url))
}

func (o *OssQiNiu) GetKeyByUrl(ctx context.Context, url string) string {
	key, err := url2.QueryUnescape(strings.TrimPrefix(url, o.BaseUrl+"/"))
	if err != nil {
		return ""
	}
	return key
}

func (o *OssQiNiu) GetBaseUrl(ctx context.Context) string {
	return o.BaseUrl
}

func (o *OssQiNiu) GetBucketName(ctx context.Context) string {
	return o.BucketName
}

func NewOssQiNiu(cfg *ncfg.FileSave) *OssQiNiu {
	return &OssQiNiu{
		AccessKey:  cfg.OssKey,
		SecretKey:  cfg.OssSecret,
		BucketName: cfg.OssBucket,
		BaseUrl:    cfg.OssEndPoint,
		qboxMac:    qbox.NewMac(cfg.OssKey, cfg.OssSecret),
	}
}
