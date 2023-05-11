package user_ser

import (
	"blog_server/service/redis_ser"
	"blog_server/utils/jwts"
	"context"
	"time"
)

var ctx = context.Background()

type UserService struct {
}

// Logout用户注销后redis创建key
func (UserService) Logout(claims *jwts.CustomClaims, token string) error {
	//计算过期时间
	exp := claims.ExpiresAt
	now := time.Now()
	diff := exp.Time.Sub(now)

	return redis_ser.Logout(token, diff)
}
