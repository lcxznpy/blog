package images_api

import (
	"blog_server/global"
	"blog_server/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/fs"
	"os"
	"path"
)

// 多个文件响应码
type FileUploadResponse struct {
	FileName  string `json:"file_name"`
	IsSuccess bool   `json:"is_success"` //是否上传成功
	Msg       string `json:"msg"`        //返回的消息
}

// ImageUploadView 上传单个图片,并返回图片的url
func (ImagesApi) ImageUploadView(c *gin.Context) {
	//多个文件
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
		err = os.MkdirAll(basePath, fs.ModePerm)
		if err != nil {
			global.Log.Error(err)
		}
	}
	//多个文件响应码
	var resList []FileUploadResponse

	for _, file := range filelist {
		filePath := path.Join(basePath, file.Filename)
		//判断图片大小
		size := float64(file.Size) / float64(1024*1024)
		if size >= float64(global.Config.Upload.Size) {
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       fmt.Sprintf("图片上传失败,当前大小为:%.2f MB,大小不能超过:%dMB", size, global.Config.Upload.Size),
			})
			continue
		}

		err := c.SaveUploadedFile(file, filePath)
		if err != nil {
			global.Log.Error(err)
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       err.Error(),
			})
			continue
		}

		resList = append(resList, FileUploadResponse{
			FileName:  filePath,
			IsSuccess: true,
			Msg:       "上传成功",
		})

	}
	res.OkWithData(resList, c)
	//单个图片
	//file, err := c.FormFile("photo")
	//if err != nil {
	//	res.FailWithMessage(err.Error(), c)
	//	return
	//}
	//fmt.Println(file.Header)
	//fmt.Println(file.Size)
	//fmt.Println(file.Filename)
}
