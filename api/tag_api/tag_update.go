package tag_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

func (TagApi) TagUpdateView(c *gin.Context) {
	id := c.Param("id")
	var cr TagRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	//根据id判断该广告是否存在
	var tag models.TagModel
	err = global.DB.Take(&tag, id).Error
	if err != nil {
		res.FailWithMessage("该标签不存在", c)
		return
	}
	maps := structs.Map(&cr) //结构体转map,用的第三方包
	//想要把true值修改为false值直接传false值gorm会自动忽略,需要传map进行修改
	err = global.DB.Model(&tag).Updates(maps).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改标签记录失败", c)
		return
	}
	res.OkWithMessage("修改标签记录成功", c)
}
