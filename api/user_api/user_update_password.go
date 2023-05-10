package user_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/utils/jwts"
	"blog_server/utils/pwd"
	"github.com/gin-gonic/gin"
)

type UpdatePassRequest struct {
	OldPwd string `json:"old_pwd" binding:"required" msg:"请输入旧密码"` //旧密码
	Pwd    string `json:"pwd" binding:"required" msg:"请输入新密码"`     //新密码
}

func (UserApi) UserUpdatePassword(c *gin.Context) {
	var cr UpdatePassRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	_cliams, _ := c.Get("claims") //set的东西是个泛型
	claims := _cliams.(*jwts.CustomClaims)
	var user models.UserModel
	err = global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c) //可能有获取token后把用户删的情况
		return
	}
	//密码是否一致
	if !pwd.CheckPwd(user.Password, cr.OldPwd) {
		res.FailWithMessage("密码错误", c)
		return
	}
	hashPwd := pwd.HashPwd(cr.Pwd)
	err = global.DB.Model(&user).Update("password", hashPwd).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改密码失败", c)
		return
	}
	res.OkWithMessage("修改密码成功", c)
	return
}
