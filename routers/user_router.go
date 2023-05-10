package routers

import (
	"blog_server/api"
	"blog_server/middleware"
)

func (r RouterGroup) UserRouter() {
	app := api.ApiGroupApp.UserApi
	r.POST("/email_login", app.EmailLoginView)
	r.GET("/users", middleware.JwtAuth(), app.UserListView)
	r.PUT("/user_role", middleware.JwtAdmin(), app.UserUpdateRoleView)
	r.PUT("/user_pwd", middleware.JwtAuth(), app.UserUpdatePassword)
	r.POST("/logout", middleware.JwtAuth(), app.LogoutView)

	//r.DELETE("/menus", app.MenuRemoveView)
}
