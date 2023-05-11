package redis_ser

import (
	"blog_server/global"
	"blog_server/utils"
	"context"
	"time"
)

const prefix = "logout_"

var ctx = context.Background()

// 注销用户并创建key和设置失效时间
func Logout(token string, diff time.Duration) error {
	err := global.Redis.Set(ctx, prefix+token, "", diff).Err()
	return err
}

// 判断是否在redis中,在的话就是token失效了
func ChekcLogout(token string) bool {
	keys := global.Redis.Keys(ctx, prefix+"*").Val()
	if utils.InList("logout_"+token, keys) {
		return true
	}
	return false
}
