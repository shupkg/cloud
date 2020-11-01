package qsms

import (
	"context"

	"github.com/shupkg/cloud/sms"
	"github.com/shupkg/cloud/utils/httpx"
	"github.com/shupkg/cloud/utils/p"

	qCommon "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	qProfile "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	q "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20190711"
)

func init() {
	sms.Register("q", New, false)
}

func New(options sms.Options) sms.Service {
	return &Service{Options: options}
}

type Service struct {
	sms.Options
}

func (s *Service) getClient(ctx context.Context) (*q.Client, error) {
	client, err := q.NewClient(qCommon.NewCredential(s.SecretID, s.SecretKey), s.Region, qProfile.NewClientProfile())
	if err != nil {
		return nil, err
	}
	client.WithHttpTransport(httpx.Context(ctx, false, nil))
	return client, nil
}

func (s *Service) CreateTemplate(ctx context.Context, options sms.CreateTemplateOptions) (templateId string, err error) {
	request := q.NewAddSmsTemplateRequest()
	/* 模板名称 */
	request.TemplateName = p.StringP(options.Name)
	/* 模板内容 */
	request.TemplateContent = p.StringP(options.Template)
	/* 短信类型：0表示普通短信, 1表示营销短信 */
	request.SmsType = p.Uint64P(options.SmsType)
	/* 是否国际/港澳台短信： 0：表示国内短信 1：表示国际/港澳台短信 */
	if options.Internal {
		request.International = p.Uint64P(0)
	} else {
		request.International = p.Uint64P(1)
	}
	/* 模板备注：例如申请原因，使用场景等 */
	request.Remark = p.StringP(options.Reason)

	client, err := s.getClient(ctx)
	if err != nil {
		return "", err
	}

	response, err := client.AddSmsTemplate(request)
	if err != nil {
		return "", err
	}

	if !p.IsNil(response.Response) && !p.IsNil(response.Response.AddTemplateStatus) && !p.IsNil(response.Response.AddTemplateStatus.TemplateId) {
		templateId = p.String(response.Response.AddTemplateStatus.TemplateId)
	}

	return templateId, nil
}

//手机号码，采用 e.164 标准，+[国家或地区码][手机号] * 例如+8613711112222
func (s *Service) Send(ctx context.Context, options sms.SendOptions) (err error) {
	request := q.NewSendSmsRequest()
	request.SmsSdkAppid = p.StringP(s.SDKAppID)    //短信应用 ID
	request.SenderId = p.StringP(s.SenderID)       //国际/港澳台短信 SenderID
	request.Sign = p.StringP(options.Message.Sign) //短信签名
	request.TemplateParamSet = p.StringPs(options.Message.TemplateParams)
	request.TemplateID = p.StringP(options.Message.TemplateID)
	request.PhoneNumberSet = p.StringPs(options.Recipients)

	client, err := s.getClient(ctx)
	if err != nil {
		return err
	}

	_, err = client.SendSms(request)
	if err != nil {
		return err
	}

	return nil
}
