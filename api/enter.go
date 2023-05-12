package api

import (
	"blog_server/api/advert_api"
	"blog_server/api/images_api"
	"blog_server/api/menu_api"
	"blog_server/api/settings_api"
	"blog_server/api/tag_api"
	"blog_server/api/user_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingApi
	ImagesApi   images_api.ImagesApi
	AdvertApi   advert_api.AdvertApi
	MenuApi     menu_api.MenuApi
	UserApi     user_api.UserApi
	TagApi      tag_api.TagApi
}

var ApiGroupApp = new(ApiGroup)
