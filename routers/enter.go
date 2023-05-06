package routers

import (
	_ "blog_server/docs" // 千万不要忘了导入把你上一步生成的docs
	"blog_server/global"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

type RouterGroup struct {
	*gin.RouterGroup
}

// 路由分层分组
func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	r := gin.Default()
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	apiGroup := r.Group("api")
	routerGroupApp := RouterGroup{apiGroup}
	routerGroupApp.SettingsRouter() //系统配置api
	routerGroupApp.ImagesRouter()   //图片上传api
	routerGroupApp.AdvertRouter()   //广告api
	return r
}
