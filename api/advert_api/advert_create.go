package advert_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"github.com/gin-gonic/gin"
)

type AdvertModel struct {
	Title  string `json:"title" binding:"required" msg:"请输入标题" structs:"title"`         // 显示的标题
	Href   string `json:"href" binding:"required,url" msg:"跳转链接不合法" structs:"href"`     // 跳转链接
	Images string `json:"images" binding:"required,url" msg:"图片地址不合法" structs:"images"` // 图片
	IsShow bool   `json:"is_show" structs:"is_show"`                                    // 是否展示
}

func (AdvertApi) AdvertCreateView(c *gin.Context) {
	var cr AdvertModel
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	//重复广告的判断
	var advert models.AdvertModel
	err = global.DB.Take(&advert, "title = ?", cr.Title).Error
	if err == nil {
		res.FailWithMessage("该广告已存在", c)
		return
	}

	err = global.DB.Create(&models.AdvertModel{
		Title:  cr.Title,
		Href:   cr.Href,
		Images: cr.Images,
		IsShow: cr.IsShow,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("添加广告记录失败", c)
		return
	}
	res.OkWithMessage("添加广告记录成功", c)
}
