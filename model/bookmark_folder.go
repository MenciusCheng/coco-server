package model

import (
	"time"
)

// 书签文件夹信息
type BookmarkFolder struct {
	// ID
	Id int64 `gorm:"primary_key;column:id" json:"id"`
	// 文件夹名称
	Name string `gorm:"column:name" json:"name"`
	// 父文件夹ID
	ParentId int64 `gorm:"column:parent_id" json:"parentId"`
	// 创建时间
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	// 更新时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

func (obj *BookmarkFolder) TableName() string {
	return "bookmark_folder"
}

type BookmarkFolderQueryReq struct {
	Page     int    `json:"page"`     // 页码
	Size     int    `json:"size"`     // 每页大小
	Name     string `json:"name"`     // 文件夹名称
	ParentId int64  `json:"parentId"` // 父文件夹ID
}

type BookmarkFolderQueryRes struct {
	List  []*BookmarkFolder `json:"list"`  // 列表
	Total int64             `json:"total"` // 总数
}

type BookmarkFolderCreateReq struct {
	Name     string `json:"name"`     // 文件夹名称
	ParentId int64  `json:"parentId"` // 父文件夹ID
}

type BookmarkFolderCreateRes struct {
	Id int64 `json:"id"` // ID
}

type BookmarkFolderUpdateReq struct {
	Id       int64  `json:"id"`       // ID
	Name     string `json:"name"`     // 文件夹名称
	ParentId int64  `json:"parentId"` // 父文件夹ID
}

type BookmarkFolderUpdateRes struct {
}

type BookmarkFolderDeleteReq struct {
	Id int64 `json:"id"` // ID
}

type BookmarkFolderDeleteRes struct {
}
