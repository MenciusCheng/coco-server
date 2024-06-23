package model

import "time"

// 流式生成器配置
type GenStreamConf struct {
	// ID
	Id int64 `gorm:"primary_key;column:id" json:"id"`
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

type GenStreamConfQueryReq struct {
	Page         int    `json:"page"`         // 页码
	Size         int    `json:"size"`         // 每页大小
	ConfNameLike string `json:"confNameLike"` // 配置名称模糊匹配
}

type GenStreamConfQueryRes struct {
	List  []*GenStreamConf `json:"list"`  // 列表
	Total int64            `json:"total"` // 总数
}

type GenStreamConfCreateReq struct {
	ConfName string `json:"confName"` // 配置名称
	Extend   string `json:"extend"`   // 扩展配置
}

type GenStreamConfCreateRes struct {
	Id int64 `json:"id"` // ID
}

type GenStreamConfUpdateReq struct {
	Id       int64  `json:"id"`       // ID
	ConfName string `json:"confName"` // 配置名称
	Extend   string `json:"extend"`   // 扩展配置
}

type GenStreamConfUpdateRes struct {
}

type GenStreamConfDeleteReq struct {
	Id int64 `json:"id"` // ID
}

type GenStreamConfDeleteRes struct {
}
