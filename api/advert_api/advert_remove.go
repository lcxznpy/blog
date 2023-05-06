package advert_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 批量删除
func (AdvertApi) AdvertRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var advertList []models.AdvertModel
	count := global.DB.Find(&advertList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("广告不存在", c)
		return
	}
	global.DB.Delete(&advertList)
	res.OkWithMessage(fmt.Sprintf("成功删除 %d 广告", count), c)
}
