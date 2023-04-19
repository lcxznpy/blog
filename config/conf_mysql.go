package config

type Mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Db       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	//Config  string `yaml:"config"`
	LogLevel string `yaml:"log_level"` //日志等级,debug输出全部sql,
}
