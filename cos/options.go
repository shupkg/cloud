package cos

//Options is Cos's config
type Options struct {
	AppID         string `json:"app_id"`
	SecretID      string `json:"secret_id" toml:"secret_id" yaml:"secret_id"`    //安全编号
	SecretKey     string `json:"secret_key" toml:"secret_key" yaml:"secret_key"` //安全密钥
	SecurityToken string `json:"security_token"`                                 //临时令牌
	Bucket        string `json:"bucket" toml:"bucket" yaml:"bucket"`             //存储桶
	Region        string `json:"region" toml:"region" yaml:"region"`             //储存桶区域
}

func (option *Options) Default() {
}
