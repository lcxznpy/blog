package user_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/utils/jwts"
	"blog_server/utils/pwd"
	"github.com/gin-gonic/gin"
)

type EmailLoginRequest struct {
	UserName string `json:"user_name" binding:"required" msg:"请输入用户名"`
	Password string `json:"password" binding:"required" msg:"请输入密码"`
}

// 邮箱或用户名登录
func (UserApi) EmailLoginView(c *gin.Context) {
	var cr EmailLoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	var userModel models.UserModel
	err = global.DB.Take(&userModel, "user_name = ? or email = ?", cr.UserName, cr.UserName).Error
	if err != nil {
		//没找到用户
		global.Log.Warnln("用户名不存在")
		res.FailWithMessage("用户名或密码不存在", c)
		return
	}
	//校验密码
	isCheck := pwd.CheckPwd(userModel.Password, cr.Password)
	if !isCheck {
		global.Log.Warn("密码错误")
		res.FailWithMessage("用户名或密码错误", c)
		return
	}
	//登录成功,生成token
	token, err := jwts.GetToken(jwts.JwtPayLoad{
		Nickname: userModel.NickName,
		Role:     int(userModel.Role),
		UserID:   userModel.ID,
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("生成token失败", c)
		return
	}
	res.OkWithData(token, c)

}
