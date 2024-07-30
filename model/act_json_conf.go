package model

import (
	"time"
)

// 活动JSON配置
type ActJsonConf struct {
	// ID
	Id int64 `gorm:"primary_key;column:id" json:"id"`
	// 名称
	Name string `gorm:"column:name" json:"name"`
	// 扩展配置
	Extend string `gorm:"column:extend" json:"extend"`
	// 创建时间
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	// 更新时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

func (obj *ActJsonConf) TableName() string {
	return "act_json_conf"
}

type ActJsonConfQueryReq struct {
	Page   int    `json:"page"`   // 页码
	Size   int    `json:"size"`   // 每页大小
	Name   string `json:"name"`   // 名称
	Extend string `json:"extend"` // 扩展配置
}

type ActJsonConfQueryRes struct {
	List  []*ActJsonConf `json:"list"`  // 列表
	Total int64          `json:"total"` // 总数
}

type ActJsonConfCreateReq struct {
	Name   string `json:"name"`   // 名称
	Extend string `json:"extend"` // 扩展配置
}

type ActJsonConfCreateRes struct {
	Id int64 `json:"id"` // ID
}

type ActJsonConfUpdateReq struct {
	Id     int64  `json:"id"`     // ID
	Name   string `json:"name"`   // 名称
	Extend string `json:"extend"` // 扩展配置
}

type ActJsonConfUpdateRes struct {
}

type ActJsonConfDeleteReq struct {
	Id int64 `json:"id"` // ID
}

type ActJsonConfDeleteRes struct {
}
