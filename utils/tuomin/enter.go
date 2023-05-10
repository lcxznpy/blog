package tuomin

import "strings"

// 手机号脱敏
func DesensitizationTel(tel string) string {
	if len(tel) != 11 {
		return ""
	}
	return tel[:3] + "****@" + tel[7:]

}

// 邮箱脱敏
func DesensitizationEmail(email string) string {
	elist := strings.Split(email, "@")
	if len(elist) != 2 {
		return ""
	}
	return elist[0][0:1] + "****" + elist[1]
}
