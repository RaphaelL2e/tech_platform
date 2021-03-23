package ossutil

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"mime/multipart"
)

type OSSHelper struct {
	OSSClient *oss.Client
}

func NewClient() OSSHelper{
	client, err := oss.New("oss-cn-beijing.aliyuncs.com", "LTAI5tM7A2B2d1VzSsSYSfCn", "rUxPbk3ZFmx5hw4GKqkbWG5KQkQ2fO", oss.EnableCRC(false))
	if err != nil {
		fmt.Println("Error:", err)
		return OSSHelper{}
	}

	// 添加Referer白名单，且不允许空Referer。Referer参数支持通配符星号（*）和问号（？）。
	referers := []string{}
	err = client.SetBucketReferer("techplatform", referers, false)
	if err != nil {
		fmt.Println("Error:", err)
		return OSSHelper{}
	}
	return OSSHelper{OSSClient: client}
}

func (h OSSHelper) Upload(file *multipart.FileHeader) (string, error) {
	result :="https://techplatform.oss-cn-beijing.aliyuncs.com/"
	bucket, err := h.OSSClient.Bucket("techplatform")
	if err != nil {
		fmt.Println("Error:", err)
		return "",err
	}

	// 指定存储类型为标准存储，缺省也为标准存储。
	storageType := oss.ObjectStorageClass(oss.StorageStandard)

	// 指定访问权限为公共读，缺省为继承bucket的权限。
	objectAcl := oss.ObjectACL(oss.ACLPublicRead)
	// 上传file。
	src,err :=file.Open()
	if err!=nil{
		return "", err
	}
	err = bucket.PutObject(file.Filename,src, storageType, objectAcl)
	if err != nil {
		return "", err
	}
	result+=file.Filename
	return result, nil
}
