package cos

// Object is the meta info of the object
type Object struct {
	Key          string `json:"key,omitempty"`           //key（文件名）
	ETag         string `json:"e_tag,omitempty"`         //ETag
	Size         int    `json:"size,omitempty"`          //Size
	LastModified string `json:"last_modified,omitempty"` //最后修改时间
	StorageClass string `json:"storage_class,omitempty"` //存储类型
	VersionId    string `json:"version_id,omitempty"`    //版本
}

// ListObjectsResult is the result of Cos.ListObjects
type ListObjectsResult struct {
	Name       string   `json:"name,omitempty"`        //存储桶名称
	Prefix     string   `json:"prefix,omitempty"`      //前缀
	Marker     string   `json:"marker,omitempty"`      //开始标记
	NextMarker string   `json:"next_marker,omitempty"` //下一次标记
	Delimiter  string   `json:"delimiter,omitempty"`   //分隔符
	MaxKeys    int      `json:"max_keys,omitempty"`    //最大获取数量
	Contents   []Object `json:"contents,omitempty"`    //对象集合
}

// ListObjectsOptions is the option of Cos.ListObjects
type ListObjectsOptions struct {
	Prefix       string `json:"prefix,omitempty"`        //前缀
	Delimiter    string `json:"delimiter,omitempty"`     //分隔符
	EncodingType string `json:"encoding_type,omitempty"` //编码类型
	Marker       string `json:"marker,omitempty"`        //开始标记
	MaxKeys      int    `json:"max_keys,omitempty"`      //最大获取数量
}
