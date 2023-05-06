package menu_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"github.com/gin-gonic/gin"
)

type Banner struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
}

type MenuResponse struct {
	models.MenuModel
	Banners []Banner `json:"banners"`
}

// MenuListView 菜单列表   自定义连接表的图片顺序输出
func (MenuApi) MenuListView(c *gin.Context) {
	//先查菜单
	var menuList []models.MenuModel
	var menuIdList []uint
	//根据菜单列表中存在的id赋给menuIdList表
	global.DB.Order("sort desc").Find(&menuList).Select("id").Scan(&menuIdList)
	//查连接表
	var menuBanners []models.MenuBannerModel
	//与菜单有关联图片全部取出来存进menuBanners中
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBanners, "menu_id in ?", menuIdList)
	var menus []MenuResponse //每个菜单连着图片当作一条响应
	for _, model := range menuList {
		//var banners []Banner //存当前菜单下的图片
		var banners = make([]Banner, 0) //解决不传图片造成返回值为nil造成的无限循环问题
		for _, banner := range menuBanners {
			if model.ID != banner.MenuID {
				continue
			}
			banners = append(banners, Banner{
				ID:   banner.BannerID,
				Path: banner.BannerModel.Path,
			})
		}
		menus = append(menus, MenuResponse{
			MenuModel: model,
			Banners:   banners,
		})
	}
	res.OkWithData(menus, c)
	//fmt.Println(menuList, menuIdList)
}
