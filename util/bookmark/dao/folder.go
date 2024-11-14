package dao

import (
	"coco-server/util/bookmark/model"
	"gorm.io/gorm"
)

type FolderDAO struct {
	DB *gorm.DB
}

func NewFolderDAO(db *gorm.DB) *FolderDAO {
	return &FolderDAO{DB: db}
}

func (dao *FolderDAO) Create(folder *model.Folder) error {
	// 忽略 created_at 和 updated_at 字段
	return dao.DB.Omit("CreatedAt", "UpdatedAt").Create(folder).Error
}

func (dao *FolderDAO) GetByID(id uint) (*model.Folder, error) {
	var folder model.Folder
	if err := dao.DB.First(&folder, id).Error; err != nil {
		return nil, err
	}
	return &folder, nil
}

func (dao *FolderDAO) Update(folder *model.Folder) error {
	// 忽略 created_at 和 updated_at 字段
	return dao.DB.Omit("CreatedAt", "UpdatedAt").Save(folder).Error
}

func (dao *FolderDAO) Delete(id uint) error {
	return dao.DB.Delete(&model.Folder{}, id).Error
}

func (dao *FolderDAO) GetFoldersByParentID(parentID uint) ([]model.Folder, error) {
	var folders []model.Folder
	if err := dao.DB.Where("parent_folder_id = ?", parentID).Find(&folders).Error; err != nil {
		return nil, err
	}
	return folders, nil
}

func (dao *FolderDAO) GetAllFolders() ([]*model.Folder, error) {
	var folders []*model.Folder
	if err := dao.DB.Find(&folders).Error; err != nil {
		return nil, err
	}
	return folders, nil
}
