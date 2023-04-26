package settings_api

import "github.com/gin-gonic/gin"

func (SettingApi) SettingApiInfoView(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "settings_api_info_view",
	})
}
