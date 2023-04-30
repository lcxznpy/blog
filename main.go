package main

import (
	"blog_server/core"
	"blog_server/flag"
	"blog_server/global"
	"blog_server/routers"
)

func main() {
	//读取配置文件
	core.InitConf()
	//fmt.Println(global.Config)
	//初始化日志
	global.Log = core.InitLogger()

	//连接数据库
	global.DB = core.InitGorm()

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
