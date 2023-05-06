package routers

import "blog_server/api"

func (r RouterGroup) ImagesRouter() {
	app := api.ApiGroupApp.ImagesApi
	r.GET("/images", app.ImageListView)
	r.GET("/image_names", app.ImageNameListView)
	r.POST("/images", app.ImageUploadView)
	r.DELETE("/images", app.ImageRemoveView)
	r.PUT("/images", app.ImageUpdateView)
}
