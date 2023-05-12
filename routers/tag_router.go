package routers

import (
	"blog_server/api"
	"blog_server/middleware"
)

func (r RouterGroup) TagRouter() {
	app := api.ApiGroupApp.TagApi
	r.POST("/tags", middleware.JwtAuth(), app.TagCreateView)
	r.GET("/tags", middleware.JwtAdmin(), app.TagListView)

	r.PUT("/tags/:id", middleware.JwtAuth(), app.TagUpdateView)
	r.DELETE("/tags", middleware.JwtAdmin(), app.TagRemoveView)
}
