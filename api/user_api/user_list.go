package user_api

import (
	"blog_server/models"
	"blog_server/models/ctype"
	"blog_server/models/res"
	"blog_server/service/common"
	"blog_server/utils/jwts"
	"blog_server/utils/tuomin"
	"github.com/gin-gonic/gin"
)

// UserListView 用户列表,不同权限看到的信息不同
func (UserApi) UserListView(c *gin.Context) {
	_cliams, _ := c.Get("claims") //set的东西是个泛型
	claims := _cliams.(*jwts.CustomClaims)
	var page models.PageInfo
	if err := c.ShouldBindQuery(&page); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var users []models.UserModel
	list, count, _ := common.ComList(models.UserModel{}, common.Option{
		PageInfo: page,
	})
	for _, user := range list {
		if ctype.Role(claims.Role) != ctype.PermissionAdmin {
			user.UserName = ""
		}
		//脱敏
		user.Tel = tuomin.DesensitizationTel(user.Tel)
		user.Email = tuomin.DesensitizationEmail(user.Email)
		users = append(users, user)
	}
	res.OkWithList(users, count, c)
}
