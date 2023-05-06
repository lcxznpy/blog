package routers

import "blog_server/api"

func (r RouterGroup) MenuRouter() {
	app := api.ApiGroupApp.MenuApi
	r.GET("/menus", app.MenuListView)
	r.GET("/menu_names", app.MenuNameList)
	r.POST("/menus", app.MenuCreateView)
	//r.DELETE("/images", app.ImageRemoveView)
	r.PUT("/menus/:id", app.MenuUpdateView)
}
