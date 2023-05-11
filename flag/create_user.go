package flag

import (
	"blog_server/global"
	"blog_server/models/ctype"
	"blog_server/service/user_ser"
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

	//密码一致
	if password != repassword {
		global.Log.Error("两次密码不一致")
		return
	}

	//管理员验证
	role := ctype.PermissionUser
	if permissions == "admin" {
		role = ctype.PermissionAdmin
	}
	//调用创建用户函数
	err := user_ser.UserService{}.CreateUser(username, nickname, password, role, email, "127.0.0.1")
	if err != nil {
		global.Log.Error(err)
		return
	}
	global.Log.Infof("用户%s创建成功", username)
}
