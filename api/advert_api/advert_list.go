package advert_api

import (
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/common"
	"github.com/gin-gonic/gin"
	"strings"
)

func (AdvertApi) AdvertListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	referer := c.GetHeader("Referer")
	isShow := true
	if strings.Contains(referer, "admin") {
		//admin来的,返回所有记录
		isShow = false
	}
	//如果不是admin来的,只返回isshow=true的记录
	list, count, err := common.ComList(models.AdvertModel{IsShow: isShow}, common.Option{
		PageInfo: cr,
		Debug:    true,
	})
	res.OkWithList(list, count, c) //响应列表优化
	return
}
