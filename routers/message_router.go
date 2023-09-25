package routers

import (
	"blog_server/api"
	"blog_server/middleware"
)

func (r RouterGroup) MessageRouter() {
	app := api.ApiGroupApp.MessageApi
	r.POST("/message", middleware.JwtAuth(), app.MessageCreateView)
	r.GET("/message_all", middleware.JwtAdmin(), app.MessageListAllView)
	r.GET("/message", middleware.JwtAuth(), app.MessageListView)
	//r.PUT("/settings/:name", settingsapi.SettingApiInfoUpdateView)

}
