package config

type QQ struct {
	AppID    string `json:"app_id" yaml:"app_id" binding:"required"`
	Key      string `json:"key" yaml:"key"`
	Redirect string `json:"redirect" yaml:"redirect"` //登录之后的回调地址
}

//缺了
