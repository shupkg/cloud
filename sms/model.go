package sms

//短信内容
type Message struct {
	Sign           string   //短息签名
	TemplateID     string   //短信模板
	TemplateKeys   []string //短信模板字段
	TemplateParams []string //短信模板参数
}

//创建模板的请求
type CreateTemplateOptions struct {
	Name     string //模板名称
	Template string //模板内容
	SmsType  uint64 //0表示普通短信, 1表示营销短信
	Internal bool   //是为国内短信，否为国际/港澳台短信
	Reason   string //申请原因
}

//发送短信的请求
type SendOptions struct {
	Message    Message  //短信内容
	Recipients []string //接收手机号码，采用 e.164 标准，+[国家或地区码][手机号] * 例如+8613711112222
}
