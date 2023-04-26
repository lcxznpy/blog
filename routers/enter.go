package routers

import (
	"blog_server/global"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	r := gin.Default()
	SettingsRouter(r) //系统配置api
	return r
}
