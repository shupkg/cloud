package sts

//授权策略
type Policy struct {
	Version   string      `json:"version,omitempty"`   //版本
	Statement []Statement `json:"statement,omitempty"` //授权声明
}

//授权声明
type Statement struct {
	Action    []string                          `json:"action,omitempty"`    //动作
	Effect    string                            `json:"effect,omitempty"`    //运行还是拒绝
	Resource  []string                          `json:"resource,omitempty"`  //资源
	Condition map[string]map[string]interface{} `json:"condition,omitempty"` //条件
}
