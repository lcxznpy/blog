package routers

import (
	"blog_server/api"
	"blog_server/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

var store = cookie.NewStore([]byte("qwerqwer2134"))

func (r RouterGroup) UserRouter() {
	app := api.ApiGroupApp.UserApi
	r.Use(sessions.Sessions("sessionid", store))
	r.POST("/email_login", app.EmailLoginView)
	r.GET("/users", middleware.JwtAuth(), app.UserListView)
	r.POST("/users", middleware.JwtAdmin(), app.UserCreateView)
	r.PUT("/user_role", middleware.JwtAdmin(), app.UserUpdateRoleView)
	r.PUT("/user_pwd", middleware.JwtAuth(), app.UserUpdatePassword)
	r.POST("/logout", middleware.JwtAuth(), app.LogoutView)
	r.POST("/user_bind_email", middleware.JwtAuth(), app.UserBindEmailView)

	r.DELETE("/users", middleware.JwtAdmin(), app.UserRemoveView)
}
