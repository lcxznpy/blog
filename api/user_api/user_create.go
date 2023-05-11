package user_api

import (
	"blog_server/global"
	"blog_server/models/ctype"
	"blog_server/models/res"
	"blog_server/service/user_ser"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserCreateRequest struct {
	NickName string     `json:"nick_name,select(c|info)" binding:"required" msg:"请输入昵称"` // 昵称
	UserName string     `json:"user_name" binding:"required" msg:"请输入用户名"`               // 用户名
	Password string     `json:"password" binding:"required" msg:"请输入密码"`                 // 密码
	Role     ctype.Role `json:"role,select(info)" binding:"required" msg:"请输入权限级别"`      // 权限  1 管理员  2 普通用户  3 游客
}

func (UserApi) UserCreateView(c *gin.Context) {
	var cr UserCreateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	//调用创建用户函数
	err = user_ser.UserService{}.CreateUser(cr.UserName, cr.NickName, cr.Password, cr.Role, "", c.ClientIP())
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage(fmt.Sprintf("用户%s创建成功", cr.UserName), c)
	return
}
