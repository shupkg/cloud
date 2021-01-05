package qcos

import (
	"context"
	"errors"
	"fmt"
	"github.com/shupkg/cloud/cos"
	"github.com/shupkg/cloud/utils/iox"
	"io"
	"net/http"
	"net/url"

	q "github.com/tencentyun/cos-go-sdk-v5"
)

func init() {
	cos.Register("q", New, false)
}

func New(options cos.Options) cos.Service {
	return &Service{Options: options}
}

//https://estar-1302894783.cos.ap-shanghai.myqcloud.com
type Service struct {
	cos.Options
}

func (s *Service) getClient() *q.Client {
	u, _ := url.Parse(fmt.Sprintf("https://%s.cos.%s.myqcloud.com", s.Bucket, s.Region))
	return q.NewClient(
		&q.BaseURL{
			BucketURL: u,
		},
		&http.Client{
			Transport: &q.AuthorizationTransport{
				SecretID:  s.SecretID,
				SecretKey: s.SecretKey,
			},
		},
	)
}

//获取预签名URL
func (s *Service) GetPreSignedURL(ctx context.Context, options cos.GetPreSignedURLOptions) (string, error) {
	u, err := s.getClient().Object.GetPresignedURL(ctx, options.Method, options.Name, s.SecretID, s.SecretKey, options.ExpiredIn, nil)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

//获取对象列表
func (s *Service) ListObjects(ctx context.Context, options cos.ListObjectsOptions) (cos.ListObjectsResult, error) {
	qResult, _, err := s.getClient().Bucket.Get(ctx, &q.BucketGetOptions{
		Prefix:       options.Prefix,
		Delimiter:    options.Delimiter,
		EncodingType: options.EncodingType,
		Marker:       options.Marker,
		MaxKeys:      options.MaxKeys,
	})
	if err != nil {
		return cos.ListObjectsResult{}, err
	}

	result := cos.ListObjectsResult{
		Name:       qResult.Name,
		Prefix:     qResult.Prefix,
		Marker:     qResult.Marker,
		NextMarker: qResult.NextMarker,
		Delimiter:  qResult.Delimiter,
		MaxKeys:    qResult.MaxKeys,
	}

	for _, content := range qResult.Contents {
		result.Contents = append(result.Contents, cos.Object{
			Key:          content.Key,
			ETag:         content.ETag,
			Size:         content.Size,
			LastModified: content.LastModified,
			StorageClass: content.StorageClass,
			VersionId:    content.VersionId,
		})
	}

	return result, nil
}

//下载对象, 下载一个 Object（文件/对象）至本地
func (s *Service) GetObject(ctx context.Context, name string) (io.ReadCloser, error) {
	resp, err := s.getClient().Object.Get(ctx, name, nil)
	if err != nil {
		return nil, s.checkError(err)
	}
	return resp.Body, nil
}

//下载对象, 下载一个 Object（文件/对象）至本地
func (s *Service) DownloadObject(ctx context.Context, name, filepath string) error {
	body, err := s.GetObject(ctx, name)
	if err != nil {
		return s.checkError(err)
	}
	return iox.ReadToFile(filepath,body)
}

//简单上传对象, 上传一个 Object（文件/对象）至 Bucket
func (s *Service) PutObject(ctx context.Context, name string, src io.Reader) error {
	_, err := s.getClient().Object.Put(ctx, name, src, nil)
	return s.checkError(err)
}

//分块上传文件, 上传一个 Object（文件/对象）至 Bucket
func (s *Service) UploadObject(ctx context.Context, name string, filepath string) error {
	_, _, err := s.getClient().Object.Upload(ctx, name, filepath, nil)
	return s.checkError(err)
}

//删除一个或多个 Object（文件/对象）
func (s *Service) DeleteObject(ctx context.Context, names ...string) (map[string]string, error) {
	if len(names) == 0 {
		return nil, nil
	}

	if len(names) == 1 {
		_, err := s.getClient().Object.Delete(ctx, names[0], nil)
		return nil, s.checkError(err)
	}

	opt := &q.ObjectDeleteMultiOptions{Quiet: true}
	for _, v := range names {
		opt.Objects = append(opt.Objects, q.Object{Key: v})
	}

	delResult, _, err := s.getClient().Object.DeleteMulti(ctx, opt)
	if err != nil {
		return nil, s.checkError(err)
	}

	if len(delResult.Errors) > 0 {
		result := make(map[string]string, len(delResult.Errors))
		for _, delError := range delResult.Errors {
			result[delError.Key] = fmt.Sprintf("[%s]%s", delError.Code, delError.Message)
		}
		return result, nil
	}
	return nil, nil
}

//复制一个文件到另一个路径
func (s *Service) CopyObject(ctx context.Context, name, srcName string) error {
	c := s.getClient()
	_, _, err := c.Object.Copy(ctx, name, fmt.Sprintf("%s/%s", c.BaseURL.BucketURL.Host, srcName), nil)
	return s.checkError(err)
}

//复制一个文件到另一个路径
func (s *Service) RenameObject(ctx context.Context, name, sourceName string) error {
	err := s.CopyObject(ctx, name, sourceName)
	if err == nil {
		_, err = s.DeleteObject(ctx, sourceName)
	}
	return s.checkError(err)
}

func (s *Service) checkError(err error) error {
	var ex = new(q.ErrorResponse)
	if errors.As(err, &ex) {
		return cos.Error{
			Code:      ex.Code,
			Message:   ex.Message,
			RequestID: ex.RequestID,
			Internal:  ex,
		}
	}

	return cos.Error{
		Code:      "SDK",
		Message:   err.Error(),
		RequestID: "",
		Internal:  err,
	}
}

var _ q.ErrorResponse
