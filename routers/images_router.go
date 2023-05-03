package routers

import "blog_server/api"

func (r RouterGroup) ImagesRouter() {
	app := api.ApiGroupApp.ImagesApi
	r.GET("/images", app.ImageListView)
	r.POST("/images", app.ImageUploadView)
	r.DELETE("/images", app.ImageRemoveView)
}
