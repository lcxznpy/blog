package settings_api

import (
	"blog_server/config"
	"blog_server/core"
	"blog_server/global"
	"blog_server/models/res"
	"github.com/gin-gonic/gin"
)

// SettingApiInfoUpdateView 修改配置文件
func (SettingApi) SettingApiInfoUpdateView(c *gin.Context) {
	var cr config.SiteInfo
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	global.Config.SiteInfo = cr
	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWith(c)
}
