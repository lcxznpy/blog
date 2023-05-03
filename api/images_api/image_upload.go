package images_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/fs"
	"os"
	"path"
	"strings"
)

var (
	// WhiteImageList 图片白名单
	WhiteImageList = []string{
		"jpg",
		"png",
		"jepg",
		"ico",
		"tiff",
		"gif",
		"svg",
		"webp",
	}
)

// 多个文件响应码
type FileUploadResponse struct {
	FileName  string `json:"file_name"`
	IsSuccess bool   `json:"is_success"` //是否上传成功
	Msg       string `json:"msg"`        //返回的消息
}

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
	var resList []FileUploadResponse

	for _, file := range filelist {
		fileName := file.Filename
		namelist := strings.Split(fileName, ".")
		suffix := strings.ToLower(namelist[len(namelist)-1])
		if !utils.InList(suffix, WhiteImageList) {
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       "非法文件",
			})
			continue
		}

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

		fileObj, err := file.Open()
		if err != nil {
			global.Log.Error(err)
		}
		byteData, err := io.ReadAll(fileObj)
		imageHash := utils.Md5(byteData)

		//去数据库中通过hash找图片是否存在
		var bannerModel models.BannerModel
		err = global.DB.Take(&bannerModel, "hash = ?", imageHash).Error
		//找到了
		if err == nil {
			resList = append(resList, FileUploadResponse{
				FileName:  bannerModel.Path,
				IsSuccess: false,
				Msg:       "图片已存在",
			})
			continue
		}

		err = c.SaveUploadedFile(file, filePath)
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
		//图片存入数据库
		global.DB.Create(&models.BannerModel{
			Path:      filePath,
			Hash:      imageHash,
			Name:      fileName,
			ImageType: 1,
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
