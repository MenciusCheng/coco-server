package model

import (
	"time"
)

type Folder struct {
	ID             uint      `gorm:"primaryKey;column:id" json:"id"`
	Name           string    `gorm:"column:name" json:"name"`
	ParentFolderID uint      `gorm:"column:parent_folder_id" json:"parentFolderId"`
	CreatedAt      time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

func (*Folder) TableName() string {
	return "folder"
}
