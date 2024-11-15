package model

import (
	"time"
)

// 书签信息
type Bookmark struct {
	// ID
	Id int64 `gorm:"primary_key;column:id" json:"id"`
	// 名称
	Name string `gorm:"column:name" json:"name"`
	// 书签URL
	Url string `gorm:"column:url" json:"url"`
	// 书签图标
	Icon string `gorm:"column:icon" json:"icon"`
	// 备注
	Remark string `gorm:"column:remark" json:"remark"`
	// 文件夹ID
	FolderId int64 `gorm:"column:folder_id" json:"folderId"`
	// 创建时间
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	// 更新时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

func (obj *Bookmark) TableName() string {
	return "bookmark"
}

type BookmarkQueryReq struct {
	Page     int    `json:"page"`     // 页码
	Size     int    `json:"size"`     // 每页大小
	Name     string `json:"name"`     // 名称
	FolderId int64  `json:"folderId"` // 文件夹ID
}

type BookmarkQueryRes struct {
	List  []*Bookmark `json:"list"`  // 列表
	Total int64       `json:"total"` // 总数
}

type BookmarkCreateReq struct {
	Name     string `json:"name"`     // 名称
	Url      string `json:"url"`      // 书签URL
	Icon     string `json:"icon"`     // 书签图标
	Remark   string `json:"remark"`   // 备注
	FolderId int64  `json:"folderId"` // 文件夹ID
}

type BookmarkCreateRes struct {
	Id int64 `json:"id"` // ID
}

type BookmarkUpdateReq struct {
	Id       int64  `json:"id"`       // ID
	Name     string `json:"name"`     // 名称
	Url      string `json:"url"`      // 书签URL
	Icon     string `json:"icon"`     // 书签图标
	Remark   string `json:"remark"`   // 备注
	FolderId int64  `json:"folderId"` // 文件夹ID
}

type BookmarkUpdateRes struct {
}

type BookmarkDeleteReq struct {
	Id int64 `json:"id"` // ID
}

type BookmarkDeleteRes struct {
}
