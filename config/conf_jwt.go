package config

type Jwy struct {
	Secret  string `json:"secret" yaml:"secret" binding:"required"` //密钥
	Expires int    `json:"expires" yaml:"expires"`                  //过期时间
	Issuer  string `json:"issuer" yaml:"issuer"`                    //颁发人
}
