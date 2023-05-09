package main

import (
	"blog_server/core"
	"blog_server/global"
	"blog_server/utils/jwts"
	"fmt"
)

func main() {
	core.InitConf()
	global.Log = core.InitLogger()
	token, err := jwts.GetToken(jwts.JwtPayLoad{
		UserID:   1,
		Role:     1,
		Username: "liucan",
		Nickname: "lc",
	})
	fmt.Println(token, err)
	claims, err := jwts.ParseToken(token)
	fmt.Println(claims, err)
}
