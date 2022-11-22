package oss

import (
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

const (
	defaultEndpoint = "https://oss-cn-chengdu.aliyuncs.com"
)

var regionIds = []string{
	"cn-hangzhou",
	"cn-shanghai",
	"cn-qingdao",
	"cn-beijing",
	"cn-zhangjiakou",
	"cn-huhehaote",
	"cn-wulanchabu",
	"cn-shenzhen",
	"cn-heyuan",
	"cn-guangzhou",
	"cn-chengdu",
	"cn-hongkong",
	"us-west-1",
	"us-east-1",
	"ap-northeast-1",
	"ap-northeast-2",
	"ap-southeast-1",
	"ap-southeast-2",
	"ap-southeast-3",
	"ap-southeast-5",
	"ap-southeast-6",
	"ap-southeast-7",
	"ap-south-1",
	"eu-central-1",
	"eu-west-1",
	"me-east-1",
}

type Config struct {
	AccessKeyId     string
	AccessKeySecret string
	Endpoint        string
	Bucket          string
}

type Client struct {
	OssClient *oss.Client
	OssBucket *oss.Bucket
	Config    *Config
	UseCname  bool
}

func New(cfg *Config) (*Client, error) {
	c := &Client{
		OssClient: nil,
		OssBucket: nil,
		Config:    cfg,
		UseCname:  false,
	}
	if err := c.Validate(); err != nil {
		return nil, err
	}

	ossClient, err := oss.New(c.Config.Endpoint, c.Config.AccessKeyId, c.Config.AccessKeySecret, oss.UseCname(c.UseCname))
	if err != nil {
		return nil, err
	}

	// 设置bucket
	bucket, err := ossClient.Bucket(c.Config.Bucket)
	if err != nil {
		return nil, err
	}
	c.OssClient = ossClient
	c.OssBucket = bucket
	return c, nil
}

// Validate 验证配置参数
func (client *Client) Validate() error {
	if client.Config.AccessKeyId == "" {
		return errors.New("oss.config.AccessKeyId cannot be empty")
	}
	if client.Config.AccessKeySecret == "" {
		return errors.New("oss.config.AccessKeySecret cannot be empty")
	}
	if client.Config.Bucket == "" {
		return errors.New("oss.config.Bucket cannot be empty")
	}
	if client.Config.Endpoint == "" {
		client.Config.Endpoint = defaultEndpoint
	}
	//是否使用自定义域名
	for _, item := range regionIds {
		if client.Config.Endpoint == fmt.Sprintf("https://oss-%s.aliyuncs.com", item) || client.Config.Endpoint == fmt.Sprintf("https://oss-%s-internal.aliyuncs.com", item) {
			client.UseCname = true
		}
	}
	return nil
}

// PutObjectFromFile 从本地文件上传
// remoteFile aa/bb/c.txt aa/c.txt c.txt
// localFile /tmp/c.txt D:\\aa\\bb\\c.txt
func (client *Client) PutObjectFromFile(remoteFile, localFile string) error {
	err := client.OssBucket.PutObjectFromFile(remoteFile, localFile)
	if err != nil {
		return err
	}
	return nil
}

// PutObjectMulFromFile 分片上传
func (client *Client) PutObjectMulFromFile(remoteFile, localFile string) error {
	fd, _ := os.Open(localFile)
	defer fd.Close()
	stat, err := fd.Stat()
	if err != nil {
		return err
	}
	if stat.Size() < 1024*1024*1 {
		return client.PutObjectFromFile(remoteFile, localFile)
	}
	chunks, _ := oss.SplitFileByPartNum(localFile, 3)
	// 指定过期时间。
	expires := time.Date(2049, time.January, 10, 23, 0, 0, 0, time.UTC)
	// 如果需要在初始化分片时设置请求头，请参考以下示例代码。
	options := []oss.Option{
		oss.MetadataDirective(oss.MetaReplace),
		oss.Expires(expires),
	}
	// 步骤1：初始化一个分片上传事件，并指定存储类型为标准存储。
	imur, _ := client.OssBucket.InitiateMultipartUpload(remoteFile, options...)
	// 步骤2：上传分片。
	var parts []oss.UploadPart
	for _, chunk := range chunks {
		_, err2 := fd.Seek(chunk.Offset, io.SeekStart)
		if err2 != nil {
			return err2
		}
		// 调用UploadPart方法上传每个分片。
		part, err := client.OssBucket.UploadPart(imur, fd, chunk.Size, chunk.Number)
		if err != nil {
			fmt.Println("Error:", err)
		}
		parts = append(parts, part)
	}
	// 指定Object的读写权限为公共读，默认为继承Bucket的读写权限。
	objectAcl := oss.ObjectACL(oss.ACLPublicRead)
	// 步骤3：完成分片上传，指定文件读写权限为公共读。
	cmur, err := client.OssBucket.CompleteMultipartUpload(imur, parts, objectAcl)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	fmt.Println("cmur:", cmur)
	return nil
}

// PutObject 从字节流上传
// remoteFile aa/bb/c.txt aa/c.txt c.txt
// reader bytes.NewReader([]byte("yourObjectValueByteArrary"))
func (client *Client) PutObject(remoteFile string, reader io.Reader) error {
	err := client.OssBucket.PutObject(remoteFile, reader)
	if err != nil {
		return err
	}
	return nil
}

// GetUrlByFile 获取文件完整链接地址
func (client *Client) GetUrlByFile(remoteFile string) string {
	if client.UseCname {
		return client.Config.Endpoint + "/" + remoteFile
	}
	parse, err := url.Parse(client.Config.Endpoint)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("https://%s.%s/%s", client.Config.Bucket, parse.Host, remoteFile)
}
