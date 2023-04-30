package settings_api

import (
	"blog_server/global"
	"blog_server/models/res"
	"github.com/gin-gonic/gin"
)

// SettingApiInfoView 获取配置文件信息
func (SettingApi) SettingApiInfoView(c *gin.Context) {
	res.OkWithData(global.Config.SiteInfo, c)
}
