package images_api

import (
	"blog_server/global"
	"blog_server/models/res"
	"blog_server/service"
	"blog_server/service/image_ser"
	"github.com/gin-gonic/gin"
	"io/fs"
	"os"
)

// ImageUploadView 上传单个图片,并返回图片的url
func (ImagesApi) ImageUploadView(c *gin.Context) {
	//上传多个文件
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	filelist, ok := form.File["images"]
	if !ok {
		res.FailWithMessage("文件不存在", c)
		return
	}
	//判断路径是否存在,不存在就创建
	basePath := global.Config.Upload.Path
	_, err = os.ReadDir(basePath)
	if err != nil {
		//递归创建文件夹
		err = os.MkdirAll(basePath, fs.ModePerm)
		if err != nil {
			global.Log.Error(err)
		}
	}
	//多个文件响应码
	var resList []image_ser.FileUploadResponse

	for _, file := range filelist {
		serviceRes := service.ServiceApp.ImageService.ImageUploadService(file)
		if !serviceRes.IsSuccess {
			resList = append(resList, serviceRes)
			continue
		}
		//上传成功
		if !global.Config.QiNiu.Enable {
			//不是七牛,本地还要保存
			err = c.SaveUploadedFile(file, serviceRes.FileName)
			if err != nil {
				global.Log.Error(err)
				serviceRes.Msg = err.Error()
				serviceRes.IsSuccess = false
				resList = append(resList, serviceRes)
				continue
			}
		}
		resList = append(resList, serviceRes)

	}
	res.OkWithData(resList, c)

}

//单个图片
//file, err := c.FormFile("photo")
//if err != nil {
//	res.FailWithMessage(err.Error(), c)
//	return
//}
//fmt.Println(file.Header)
//fmt.Println(file.Size)
//fmt.Println(file.Filename)
