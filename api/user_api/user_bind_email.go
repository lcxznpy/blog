package user_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/plugins/email"
	"blog_server/utils/jwts"
	"blog_server/utils/pwd"
	"blog_server/utils/random"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type BindEmailRequest struct {
	Email    string  `json:"email" binding:"required,email" msg:"邮箱非法"`
	Code     *string `json:"code"`
	Password string  `json:"password" `
}

func (UserApi) UserBindEmailView(c *gin.Context) {
	_cliams, _ := c.Get("claims") //set的东西是个泛型
	claims := _cliams.(*jwts.CustomClaims)
	var cr BindEmailRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	session := sessions.Default(c)
	if cr.Code == nil {
		code := random.Code(4)
		session.Set("valid_code", code)
		session.Set("valid_email", cr.Email)
		err = session.Save()
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("session错误", c)
			return
		}
		err = email.NewCode().Send(cr.Email, "您的验证码是:"+code)
		if err != nil {
			global.Log.Error(err)
		}
		res.OkWithMessage("验证码已发送", c)
		return
	}
	code := session.Get("valid_code")
	if code != *cr.Code {
		res.FailWithMessage("验证码错误", c)
		return
	}
	email := session.Get("valid_email")
	if email != cr.Email {
		res.FailWithMessage("两次邮箱不一致", c)
		return
	}
	//修改用户邮箱
	var user models.UserModel
	err = global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}
	if len(cr.Password) < 4 {
		res.FailWithMessage("密码强度太低", c)
		return
	}
	hashPwd := pwd.HashPwd(cr.Password)
	err = global.DB.Model(&user).Updates(map[string]any{
		"email":    cr.Email,
		"password": hashPwd,
	}).Error
	if err != nil {
		res.FailWithMessage("修改邮箱失败", c)
		return
	}
	res.OkWithMessage("邮箱绑定成功", c)

}
