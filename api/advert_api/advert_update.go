package advert_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

// AdvertUpdateView 更新广告
// @Tags 广告管理
// @Summary 更新广告
// @Description 更新广告
// @Param data body AdvertRequest true "广告的一些参数"
// @Failure 400 {object} string "请求错误"
// @Produce json
// @Router /api/adverts/:id [put]
// @Success 200 {object} res.Response{date=string}
func (AdvertApi) AdvertUpdateView(c *gin.Context) {
	id := c.Param("id")
	var cr AdvertRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	//根据id判断该广告是否存在
	var advert models.AdvertModel
	err = global.DB.Take(&advert, id).Error
	if err != nil {
		res.FailWithMessage("该广告不存在", c)
		return
	}
	maps := structs.Map(&cr) //结构体转map,用的第三方包
	//想要把true值修改为false值直接传false值gorm会自动忽略,需要传map进行修改
	err = global.DB.Model(&advert).Updates(maps).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改广告记录失败", c)
		return
	}
	res.OkWithMessage("修改广告记录成功", c)
}
