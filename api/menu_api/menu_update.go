package menu_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

func (MenuApi) MenuUpdateView(c *gin.Context) {
	var cr MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	id := c.Param("id")
	//先把之前与当前菜单相关的图片删除
	var menuModel models.MenuModel
	err = global.DB.Take(&menuModel, id).Error
	if err != nil {
		res.FailWithMessage("菜单记录不存在", c)
		return
	}
	global.DB.Model(&menuModel).Association("Banners").Clear()
	//如果选择了图片,那就添加进第三张表
	if len(cr.ImageSortList) > 0 {
		var bannerList []models.MenuBannerModel
		for _, sort := range cr.ImageSortList {
			bannerList = append(bannerList, models.MenuBannerModel{
				MenuID:   menuModel.ID,
				BannerID: sort.ImageId,
				Sort:     sort.Sort,
			})
		}
		err = global.DB.Create(&bannerList).Error
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("创建菜单图片失败", c)
			return
		}
	}

	maps := structs.Map(&cr) //结构体转map,用的第三方包
	//想要把true值修改为false值直接传false值gorm会自动忽略,需要传map进行修改
	err = global.DB.Model(&menuModel).Updates(maps).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改菜单失败", c)
		return
	}
	res.OkWithMessage("修改菜单成功", c)
}

//????如果修改的值与数据库中有重复该怎么办？
//????如果如何保证更新的原子性？
