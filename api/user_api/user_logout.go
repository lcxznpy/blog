package user_api

import (
	"blog_server/global"
	"blog_server/models/res"
	"blog_server/utils/jwts"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

var ctx = context.Background()

func (UserApi) LogoutView(c *gin.Context) {
	_cliams, _ := c.Get("claims") //set的东西是个泛型
	claims := _cliams.(*jwts.CustomClaims)
	token := c.Request.Header.Get("token")
	//计算过期时间
	exp := claims.ExpiresAt
	now := time.Now()
	diff := exp.Time.Sub(now)
	err := global.Redis.Set(ctx, fmt.Sprintf("logout_%s", token), "", diff).Err()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("注销失败", c)
		return
	}
	res.OkWithMessage("注销成功", c)
}
