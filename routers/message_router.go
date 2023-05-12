package routers

import (
	"blog_server/api"
	"blog_server/middleware"
)

func (r RouterGroup) MessageRouter() {
	app := api.ApiGroupApp.MessageApi
	r.POST("/message", middleware.JwtAuth(), app.MessageCreateView)
	//r.PUT("/settings/:name", settingsapi.SettingApiInfoUpdateView)

}
