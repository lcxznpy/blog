package service

import (
	"blog_server/service/image_ser"
	"blog_server/service/user_ser"
)

type ServiceGroup struct {
	ImageService image_ser.ImageService
	UserService  user_ser.UserService
}

var ServiceApp = new(ServiceGroup)
