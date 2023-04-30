package routers

import (
	"blog_server/global"
	"github.com/gin-gonic/gin"
)

// 路由分层分组
func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	r := gin.Default()
	SettingsRouter(r) //系统配置api
	return r
}
