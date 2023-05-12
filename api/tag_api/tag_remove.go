package tag_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (TagApi) TagRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var tagList []models.TagModel
	count := global.DB.Find(&tagList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("标签不存在", c)
		return
	}

	//TODO :删标签前一定要确保当前标签的文章没了
	global.DB.Delete(&tagList)
	res.OkWithMessage(fmt.Sprintf("成功删除 %d 条标签", count), c)
}
