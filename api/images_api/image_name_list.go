package images_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"github.com/gin-gonic/gin"
)

type ImageResponse struct {
	ID   uint   `json:"id"`
	Path string `json:"path"` //图片路径
	Name string `json:"name"` //图片名称
}

// ImageNameListView 返回图片的部分信息
// @Tags 图片管理
// @Summary 图片名称列表
// @Description 图片名称列表
// @Produce json
// @Router /api/image_names [get]
// @Success 200 {object} res.Response{data=[]ImageResponse}
func (ImagesApi) ImageNameListView(c *gin.Context) {
	var imageList []ImageResponse
	global.DB.Model(models.BannerModel{}).Select("id", "path", "name").Scan(&imageList)
	res.OkWithData(imageList, c)
}
