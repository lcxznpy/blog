package user_api

import (
	"blog_server/global"
	"blog_server/models/res"
	"blog_server/service"
	"blog_server/utils/jwts"
	"github.com/gin-gonic/gin"
)

// LogoutView 注销用户
func (UserApi) LogoutView(c *gin.Context) {
	_cliams, _ := c.Get("claims") //set的东西是个泛型
	claims := _cliams.(*jwts.CustomClaims)
	token := c.Request.Header.Get("token")
	err := service.ServiceApp.UserService.Logout(claims, token)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("注销失败", c)
		return
	}
	res.OkWithMessage("注销成功", c)
}
