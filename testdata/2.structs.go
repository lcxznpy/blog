package main

import (
	"fmt"
	"github.com/fatih/structs"
)

type AdvertModel struct {
	Title  string `json:"title" binding:"required" msg:"请输入标题" structs:"title"`       // 显示的标题
	Href   string `json:"href" binding:"required,url" msg:"跳转链接不合法" structs:"-"`      // 跳转链接
	Images string `json:"images" binding:"required,url" msg:"图片地址不合法" structs:"-"`    // 图片
	IsShow bool   `json:"is_show" binding:"required" msg:"请选择是否展示" structs:"is_show"` // 是否展示
}

func main() {
	u1 := AdvertModel{
		Title:  "xxx",
		Href:   "xxx",
		Images: "xxx",
		IsShow: true,
	}
	m3 := structs.Map(&u1)
	fmt.Println(m3)
}
