package sts

import "time"

//获取临时密钥的请求
type GetFederationTokenOptions struct {
	Name     string
	Policy   Policy
	ExpireIn time.Duration
}

//获取临时密钥的结果
type GetFederationTokenResult struct {
	TmpSecretID  string
	TmpSecretKey string
	SessionToken string
	ExpiredTime  int64
	Expiration   string
	StartTime    int64
	RequestId    string
}
