package routers

import (
	"blog_server/api"
	"github.com/gin-gonic/gin"
)

func SettingsRouter(r *gin.Engine) {
	settingsapi := api.ApiGroupApp.SettingsApi
	r.GET("/settings", settingsapi.SettingApiInfoView)
}
