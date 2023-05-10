package main

import (
	"blog_server/core"
	_ "blog_server/docs"
	"blog_server/flag"
	"blog_server/global"
	"blog_server/routers"
)

// @title gvb_server API文档
// @version 1.0
// @description Api文档
// @host localhost:8080
// @BasePath /
func main() {
	//读取配置文件
	core.InitConf()
	//fmt.Println(global.Config)
	//初始化日志
	global.Log = core.InitLogger()

	//连接数据库
	global.DB = core.InitGorm()
	//连接redis
	global.Redis = core.ConnectionRedis()
	//命令行参数绑定
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}

	r := routers.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("blog_server运行在:%s上", addr)
	err := r.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}
}
