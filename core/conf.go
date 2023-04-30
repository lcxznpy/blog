package core

import (
	"blog_server/config"
	"blog_server/global"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/fs"
	"log"

	//"io"
	"io/ioutil"
)

const ConfigFile = "settings.yaml"

// InitConf 读取yaml中的配置文件
func InitConf() {

	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlcomf error : %s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config init unmarshal : %v", err)
	}
	log.Println("config yaml load init success")
	global.Config = c
}

func SetYaml() error {
	//转换数据格式
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		return err
	}
	//写yaml文件
	err = ioutil.WriteFile(ConfigFile, byteData, fs.ModePerm)
	if err != nil {
		return err
	}
	global.Log.Info("修改配置成功")
	return nil
}
