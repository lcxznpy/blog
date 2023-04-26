package api

import "blog_server/api/settings_api"

type ApiGroup struct {
	SettingsApi settings_api.SettingApi
}

var ApiGroupApp = new(ApiGroup)
