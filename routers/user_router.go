package routers

import "blog_server/api"

func (r RouterGroup) UserRouter() {
	app := api.ApiGroupApp.UserApi
	r.POST("/email_login", app.EmailLoginView)
	//r.GET("/menus/:id", app.MenuDetailView)
	//r.GET("/menu_names", app.MenuNameList)
	//r.POST("/menus", app.MenuCreateView)
	//r.PUT("/menus/:id", app.MenuUpdateView)
	//r.DELETE("/menus", app.MenuRemoveView)
}
