package cos

import "github.com/shupkg/cloud"

//Options is Service's config
type Options struct {
	cloud.Credential
	SecurityToken string `json:"security_token"`                     //临时令牌
	Bucket        string `json:"bucket" toml:"bucket" yaml:"bucket"` //存储桶
}

func (option *Options) Default() {
}
