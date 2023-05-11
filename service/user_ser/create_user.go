package user_ser

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/ctype"
	"blog_server/utils/pwd"
	"errors"
)

const Avatar = "/blog_server/uploads/avatar/xiaolan.png"

func (UserService) CreateUser(userName, nickName, password string, role ctype.Role, email string, ip string) error {
	//判断用户是否存在
	var userModel models.UserModel
	err := global.DB.Take(&userModel, "user_name = ?", userName).Error
	if err == nil {
		return errors.New("用户名已存在")
	}
	//头像
	//1.默认头像
	//2.随机头像

	//哈希密码
	hashPwd := pwd.HashPwd(password)

	err = global.DB.Create(&models.UserModel{
		NickName:   nickName,
		UserName:   userName,
		Password:   hashPwd,
		Email:      email,
		Role:       role,
		Avatar:     Avatar,
		Addr:       "上海市",
		IP:         ip,
		SignStatus: ctype.SignEmail,
	}).Error
	if err != nil {
		global.Log.Error(err)
		return err
	}

	return nil
}
