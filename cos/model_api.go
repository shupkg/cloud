package cos

import "time"

//GetPreSignedURLOptions is options for Service.GetPreSignedURL
type GetPreSignedURLOptions struct {
	Name      string        `json:"name,omitempty"`
	Method    string        `json:"method,omitempty"`
	AK        string        `json:"ak,omitempty"`
	SK        string        `json:"sk,omitempty"`
	ExpiredIn time.Duration `json:"expired_in,omitempty"`
}

//ListObjectsResult is the result of Service.ListObjects
type ListObjectsResult struct {
	Name       string   `json:"name,omitempty"`        //存储桶名称
	Prefix     string   `json:"prefix,omitempty"`      //前缀
	Marker     string   `json:"marker,omitempty"`      //开始标记
	NextMarker string   `json:"next_marker,omitempty"` //下一次标记
	Delimiter  string   `json:"delimiter,omitempty"`   //分隔符
	MaxKeys    int      `json:"max_keys,omitempty"`    //最大获取数量
	Contents   []Object `json:"contents,omitempty"`    //对象集合
}

//ListObjectsOptions is the option of Service.ListObjects
type ListObjectsOptions struct {
	Prefix       string `json:"prefix,omitempty"`        //前缀
	Delimiter    string `json:"delimiter,omitempty"`     //分隔符
	EncodingType string `json:"encoding_type,omitempty"` //编码类型
	Marker       string `json:"marker,omitempty"`        //开始标记
	MaxKeys      int    `json:"max_keys,omitempty"`      //最大获取数量
}
