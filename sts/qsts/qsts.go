package qsts

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/shupkg/cloud/sts"
	"github.com/shupkg/cloud/utils/httpx"
	"github.com/shupkg/cloud/utils/p"

	qCommon "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	qErrs "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	qProfile "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	q "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sts/v20180813"
)

func init() {
	sts.Register("q", New, false)
}

func New(options sts.Options) sts.Service {
	return &Service{Options: options}
}

type Service struct {
	sts.Options
}

func (s *Service) getClient(ctx context.Context) (*q.Client, error) {
	client, err := q.NewClient(qCommon.NewCredential(s.SecretID, s.SecretKey), s.Region, qProfile.NewClientProfile())
	if err != nil {
		return nil, err
	}
	client.WithHttpTransport(httpx.Context(ctx, false, nil))
	return client, nil
}

func (s *Service) GetFederationToken(ctx context.Context, options sts.GetFederationTokenOptions) (result sts.GetFederationTokenResult, err error) {
	//// 密钥的权限列表。简单上传、表单上传和分片上传需要以下的权限
	//// 其他权限列表请看 https://cloud.tencent.com/document/product/436/31923
	//policy := Policy{
	//	Version: "2.0",
	//	Statement: []Statement{
	//		{
	//			Effect: "allow",
	//			Action: []string{
	//				"name/cos:PostObject",
	//				"name/cos:PutObject",
	//				"name/cos:InitiateMultipartUpload", // 分片上传
	//				"name/cos:ListMultipartUploads",
	//				"name/cos:ListParts",
	//				"name/cos:UploadPart",
	//				"name/cos:CompleteMultipartUpload",
	//			},
	//			Resource: []string{
	//				// 这里改成允许的路径前缀，可以根据自己网站的用户登录态判断允许上传的具体路径
	//				// 例子： a.jpg 或者 a/* 或者 * (使用通配符*存在重大安全风险, 请谨慎评估使用)
	//				// qcs::cos:cn-shanghai:uid/<APPID>:<BucketName-APPID>/app/avatar/<Username>.jpg
	//				"qcs::cos:" + s.cfg.Region + ":uid/" + input.AppId + ":" + s.cfg.Bucket + "/" + input.FileType + "/*",
	//			},
	//		},
	//	},
	//}

	client, err := s.getClient(ctx)
	if err != nil {
		return result, err
	}

	req := q.NewGetFederationTokenRequest()
	req.Name = p.StringP(options.Name)
	expireSeconds := int64(options.ExpireIn.Seconds())
	req.DurationSeconds = p.Uint64P(uint64(expireSeconds))

	v, _ := json.MarshalIndent(options.Policy, "", "  ")
	req.Policy = p.StringP(string(v))

	resp, err := client.GetFederationToken(req)
	if err != nil {
		return result, err
	}

	if !p.IsNil(resp) && !p.IsNil(resp.Response) && !p.IsNil(resp.Response.Credentials) {
		result.TmpSecretID = p.String(resp.Response.Credentials.TmpSecretId)
		result.SessionToken = p.String(resp.Response.Credentials.Token)
		result.TmpSecretKey = p.String(resp.Response.Credentials.TmpSecretKey)
		result.ExpiredTime = int64(p.Uint64(resp.Response.ExpiredTime))
		result.Expiration = p.String(resp.Response.Expiration)
		result.StartTime = result.ExpiredTime - expireSeconds
	}
	return result, nil
}

func (s *Service) handleError(err error) (bizId string, ex error) {
	if err == nil {
		return "", nil
	}
	qErr := &qErrs.TencentCloudSDKError{}
	if errors.As(err, &qErr) {
		return qErr.RequestId, qErr
	}
	return "", err
}
