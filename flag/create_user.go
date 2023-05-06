package flag

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/ctype"
	"blog_server/utils/pwd"
	"fmt"
)

func CreateUser(permissions string) {
	var (
		nickname   string
		username   string
		password   string
		repassword string
		email      string
	)
	fmt.Printf("请输入用户名:")
	fmt.Scan(&username)
	fmt.Printf("请输入昵称:")
	fmt.Scan(&nickname)
	fmt.Printf("请输入密码:")
	fmt.Scan(&password)
	fmt.Printf("请再次输入密码:")
	fmt.Scan(&repassword)
	fmt.Printf("请输入email:")
	fmt.Scan(&email)

	//判断用户是否存在
	var userModel models.UserModel
	err := global.DB.Take(&userModel, "user_name = ?", username).Error
	if err == nil {
		global.Log.Error("用户已存在,请重新输入新的用户名")
		return
	}
	//密码一致
	if password != repassword {
		global.Log.Error("两次密码不一致")
		return
	}
	avatar := "/blog_server/uploads/avatar/xiaolan.png"

	hashPwd := pwd.HashPwd(password)
	role := ctype.PermissionUser
	if permissions == "admin" {
		role = ctype.PermissionAdmin
	}
	err = global.DB.Create(&models.UserModel{
		NickName:   nickname,
		UserName:   username,
		Password:   hashPwd,
		Email:      email,
		Role:       role,
		Avatar:     avatar,
		Addr:       "上海市",
		IP:         "127.0.0.1",
		SignStatus: ctype.SignEmail,
	}).Error
	if err != nil {
		global.Log.Error(err)
		return
	}
	global.Log.Infof("用户%s创建成功", username)
}
