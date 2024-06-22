package model

import "time"

// 流式生成器配置
type GenStreamConf struct {
	// ID
	Id int64 `gorm:"column:id" json:"id"`
	// 配置名称
	ConfName string `gorm:"column:conf_name" json:"confName"`
	// 扩展配置
	Extend string `gorm:"column:extend" json:"extend"`
	// 创建时间
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	// 更新时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

// get table name of GenStreamConf
func (obj *GenStreamConf) TableName() string {
	return "gen_stream_conf"
}
