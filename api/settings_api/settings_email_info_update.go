package settings_api

import (
	"blog_server/config"
	"blog_server/core"
	"blog_server/global"
	"blog_server/models/res"
	"github.com/gin-gonic/gin"
)

// SettingApiEmailInfoUpdateView 修改邮箱配置文件
func (SettingApi) SettingApiEmailInfoUpdateView(c *gin.Context) {
	var em config.Email
	err := c.ShouldBindJSON(&em)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	global.Config.Email = em
	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWith(c)
}
