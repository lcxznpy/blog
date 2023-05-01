package routers

import "blog_server/api"

func (r RouterGroup) ImagesRouter() {
	app := api.ApiGroupApp.ImagesApi
	r.POST("/images", app.ImageUploadView)
}
