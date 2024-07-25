package model

import (
	"time"
)

// 流式生成器配置项模版
type GenStreamDetailTmpl struct {
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

func (obj *GenStreamDetailTmpl) TableName() string {
	return "gen_stream_detail_tmpl"
}

type GenStreamDetailTmplQueryReq struct {
	Page   int    `json:"page"`   // 页码
	Size   int    `json:"size"`   // 每页大小
	Name   string `json:"name"`   // 名称
	Extend string `json:"extend"` // 扩展配置
}

type GenStreamDetailTmplQueryRes struct {
	List  []*GenStreamDetailTmpl `json:"list"`  // 列表
	Total int64                  `json:"total"` // 总数
}

type GenStreamDetailTmplCreateReq struct {
	Name   string `json:"name"`   // 名称
	Extend string `json:"extend"` // 扩展配置
}

type GenStreamDetailTmplCreateRes struct {
	Id int64 `json:"id"` // ID
}

type GenStreamDetailTmplUpdateReq struct {
	Id     int64  `json:"id"`     // ID
	Name   string `json:"name"`   // 名称
	Extend string `json:"extend"` // 扩展配置
}

type GenStreamDetailTmplUpdateRes struct {
}

type GenStreamDetailTmplDeleteReq struct {
	Id int64 `json:"id"` // ID
}

type GenStreamDetailTmplDeleteRes struct {
}
