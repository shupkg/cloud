package sms

import (
	"context"
)

//短信服务
type Service interface {
	CreateTemplate(ctx context.Context, ctReq CreateTemplateOptions) (templateId string, err error)
	Send(ctx context.Context, options SendOptions) (err error)
}
