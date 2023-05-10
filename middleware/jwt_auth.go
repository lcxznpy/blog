package middleware

import (
	"blog_server/models/ctype"
	"blog_server/models/res"
	"blog_server/utils/jwts"
	"github.com/gin-gonic/gin"
)

// JwtAuth 普通用户认证中间件
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			res.FailWithMessage("未接受到token", c)
			c.Abort() //没登陆成功
			return
		}
		claims, err := jwts.ParseToken(token)
		if err != nil {
			res.FailWithMessage("token错误", c)
			c.Abort() //没登陆成功
			return
		}
		//登录的用户
		c.Set("claims", claims)
	}
}

// JwtAuth 管理员认证中间件
func JwtAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			res.FailWithMessage("未接受到token", c)
			c.Abort() //没登陆成功
			return
		}
		claims, err := jwts.ParseToken(token)
		if err != nil {
			res.FailWithMessage("token错误", c)
			c.Abort() //没登陆成功
			return
		}
		if claims.Role != int(ctype.PermissionAdmin) {
			res.FailWithMessage("权限错误", c)
			c.Abort() //没登陆成功
			return
		}
		//登录的用户
		c.Set("claims", claims)
	}
}
