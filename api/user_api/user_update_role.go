package user_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/ctype"
	"blog_server/models/res"
	"github.com/gin-gonic/gin"
)

type UserRole struct {
	Role     ctype.Role `json:"role" binding:"required,oneof=1 2 3 4" msg:"权限必须在指定范围内"`
	UserID   uint       `json:"user_id" binding:"required" msg:"请输入用户id"`
	NickName string     `json:"nick_name"`
}

func (UserApi) UserUpdateRoleView(c *gin.Context) {
	var cr UserRole
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var user models.UserModel
	err = global.DB.Take(&user, cr.UserID).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("用户不存在", c)
		return
	}
	err = global.DB.Model(&user).Updates(map[string]any{
		"role":      cr.Role,
		"nick_name": cr.NickName,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改权限失败", c)
		return
	}
	res.OkWithMessage("修改权限成功", c)
}
