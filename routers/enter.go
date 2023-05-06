package routers

import (
	"blog_server/global"
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.RouterGroup
}

// 路由分层分组
func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	r := gin.Default()
	apiGroup := r.Group("api")
	routerGroupApp := RouterGroup{apiGroup}
	routerGroupApp.SettingsRouter() //系统配置api
	routerGroupApp.ImagesRouter()   //图片上传api
	routerGroupApp.AdvertRouter()   //广告api
	return r
}
