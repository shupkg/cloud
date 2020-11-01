package sms

import "github.com/shupkg/cloud"

type Options struct {
	cloud.Credential
	SDKAppID  string
	SenderID  string //国际/港澳台短信 SenderID
}
