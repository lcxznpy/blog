package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primarykey" json:"id,select($any)" structs:"-"` // 主键ID
	CreatedAt time.Time `json:"created_at,select($any)" structs:"-"`           // 创建时间
	UpdatedAt time.Time `json:"-" structs:"-"`                                 // 更新时间
}

// 批量删除
type RemoveRequest struct {
	IDList []uint `json:"id_list"`
}

type ESIDRequest struct {
	ID string `json:"id" form:"id" uri:"id"`
}
type ESIDListRequest struct {
	IDList []string `json:"id_list" binding:"required"`
}

// PageInfo 分页查询的封装
type PageInfo struct {
	Page  int    `form:"page"`  //第几页
	Key   string `form:"key"`   //模糊查询的key
	Limit int    `form:"limit"` //每页限制的记录数量
	Sort  string `form:"sort"`  //排序方法
}
