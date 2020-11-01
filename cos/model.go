package cos

//Object is the meta info of the object
type Object struct {
	Key          string `json:"key,omitempty"`           //key（文件名）
	ETag         string `json:"e_tag,omitempty"`         //ETag
	Size         int    `json:"size,omitempty"`          //Size
	LastModified string `json:"last_modified,omitempty"` //最后修改时间
	StorageClass string `json:"storage_class,omitempty"` //存储类型
	VersionId    string `json:"version_id,omitempty"`    //版本
}
