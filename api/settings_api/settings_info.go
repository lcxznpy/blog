package settings_api

import (
	"blog_server/models/res"
	"github.com/gin-gonic/gin"
)

func (SettingApi) SettingApiInfoView(c *gin.Context) {
	res.FailWithCode(1, c)
}
