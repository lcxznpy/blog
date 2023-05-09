package jwts

import (
	"blog_server/global"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"time"
)

// JwtPayLoad jwt中payload的数据
type JwtPayLoad struct {
	//Username string `json:"username"` //用户名
	Nickname string `json:"nickname"` //昵称
	Role     int    `json:"role"`     //权限 1admin 2user 3 游客
	UserID   uint   `json:"user_id"`  //用户id
}

type CustomClaims struct {
	JwtPayLoad
	jwt.StandardClaims
}

var Mysecret []byte

// 创建token
func GetToken(user JwtPayLoad) (string, error) {
	Mysecret = []byte(global.Config.Jwy.Secret)
	claim := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * time.Duration(global.Config.Jwy.Expires))), //默认过期时间2小时
			Issuer:    global.Config.Jwy.Issuer,                                                     //签发人
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(Mysecret)
}

// 解析token
func ParseToken(tokenStr string) (*CustomClaims, error) {
	Mysecret = []byte(global.Config.Jwy.Secret)
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Mysecret, nil
	})
	if err != nil {
		global.Log.Error(fmt.Sprintf("token parse err:%s", err.Error()))
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
