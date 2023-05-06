package common

import (
	"blog_server/global"
	"blog_server/models"
	"gorm.io/gorm"
)

type Option struct {
	models.PageInfo
	Debug bool
}

// ComList 分页查询统一函数
func ComList[T any](model T, option Option) (list []T, count int64, err error) {
	DB := global.DB
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog}) //独属的数据库日志等级
	}

	if option.Sort == "" {
		option.Sort = "created_at desc" // 默认按照时间往前排
	}
	//添加查询参数
	query := DB.Where(model)

	count = query.Debug().Select("id").Find(&list).RowsAffected //查询count优化
	query = DB.Where(model)
	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}
	err = query.Debug().Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error
	return list, count, err
}
