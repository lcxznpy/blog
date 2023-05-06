package api

import (
	"blog_server/api/advert_api"
	"blog_server/api/images_api"
	"blog_server/api/settings_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingApi
	ImagesApi   images_api.ImagesApi
	AdvertApi   advert_api.AdvertApi
}

var ApiGroupApp = new(ApiGroup)
