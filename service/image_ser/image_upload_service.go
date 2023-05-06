package image_ser

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/ctype"
	"blog_server/plugins/qiniu"
	"blog_server/utils"
	"fmt"
	"io"
	"mime/multipart"
	"path"
	"strings"
)

// WhiteImageList 图片白名单
var WhiteImageList = []string{
	"jpg",
	"png",
	"jepg",
	"ico",
	"tiff",
	"gif",
	"svg",
	"webp",
	"mp4",
}

type FileUploadResponse struct {
	FileName  string `json:"file_name"`
	IsSuccess bool   `json:"is_success"` //是否上传成功
	Msg       string `json:"msg"`        //返回的消息
}

// ImageUploadService 文件上传
func (ImageService) ImageUploadService(file *multipart.FileHeader) (res FileUploadResponse) {
	fileName := file.Filename
	basePath := global.Config.Upload.Path

	filePath := path.Join(basePath, file.Filename)
	res.FileName = filePath

	//文件白名单判断
	namelist := strings.Split(fileName, ".")
	suffix := strings.ToLower(namelist[len(namelist)-1])
	if !utils.InList(suffix, WhiteImageList) {
		res.Msg = "非法文件"
		return

	}

	//判断图片大小
	size := float64(file.Size) / float64(1024*1024)
	if size >= float64(global.Config.Upload.Size) {
		res.Msg = fmt.Sprintf("图片上传失败,当前大小为:%.2f MB,大小不能超过:%dMB", size, global.Config.Upload.Size)
		return res
	}
	//读取文件内容并上传文件和加密
	fileObj, err := file.Open()
	if err != nil {
		global.Log.Error(err)
	}
	byteData, err := io.ReadAll(fileObj)
	imageHash := utils.Md5(byteData)

	//去数据库中通过hash找图片是否存在
	var bannerModel models.BannerModel
	err = global.DB.Take(&bannerModel, "hash = ?", imageHash).Error
	//找到了该文件,不保存了
	if err == nil {
		res.Msg = "图片已存在"
		res.FileName = bannerModel.Path
		return res
	}
	fileType := ctype.Local

	res.Msg = "图片上传成功"
	res.IsSuccess = true

	//是否上传七牛云
	if global.Config.QiNiu.Enable {
		filePath, err = qiniu.UploadImage(byteData, fileName, global.Config.QiNiu.Prefix)
		if err != nil {

			global.Log.Error(err)
			res.Msg = err.Error()
			return res
		}
		res.FileName = filePath
		res.Msg = "成功上传七牛云"
		fileType = ctype.QiNiu
	}

	//图片存入数据库
	global.DB.Create(&models.BannerModel{
		Path:      filePath,
		Hash:      imageHash,
		Name:      fileName,
		ImageType: fileType,
	})
	return
}
