package routers

import "blog_server/api"

func (r RouterGroup) AdvertRouter() {
	app := api.ApiGroupApp.AdvertApi
	r.POST("/adverts", app.AdvertCreateView)
	//r.GET("/images", app.ImageListView)
	//r.DELETE("/images", app.ImageRemoveView)
	//r.PUT("/images", app.ImageUpdateView)
}
