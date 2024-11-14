package model

import (
	"time"
)

type Bookmark struct {
	ID        uint      `gorm:"primaryKey;column:id" json:"id"`
	Title     string    `gorm:"column:title" json:"title"`
	URL       string    `gorm:"column:url" json:"url"`
	Icon      string    `gorm:"column:icon" json:"icon"`
	FolderID  uint      `gorm:"column:folder_id" json:"folderId"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

func (*Bookmark) TableName() string {
	return "bookmark"
}
