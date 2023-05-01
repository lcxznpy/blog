package routers

import (
	"blog_server/api"
)

func (r RouterGroup) SettingsRouter() {
	settingsapi := api.ApiGroupApp.SettingsApi
	r.GET("/settings/:name", settingsapi.SettingApiInfoView)
	r.PUT("/settings/:name", settingsapi.SettingApiInfoUpdateView)

}
