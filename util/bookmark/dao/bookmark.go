package dao

import (
	"coco-server/util/bookmark/model"
	"gorm.io/gorm"
)

type BookmarkDAO struct {
	DB *gorm.DB
}

func NewBookmarkDAO(db *gorm.DB) *BookmarkDAO {
	return &BookmarkDAO{DB: db}
}

func (dao *BookmarkDAO) Create(bookmark *model.Bookmark) error {
	//忽略 created_at 和 updated_at 字段
	return dao.DB.Omit("CreatedAt", "UpdatedAt").Create(bookmark).Error
}

func (dao *BookmarkDAO) GetByID(id uint) (*model.Bookmark, error) {
	var bookmark model.Bookmark
	if err := dao.DB.First(&bookmark, id).Error; err != nil {
		return nil, err
	}
	return &bookmark, nil
}

func (dao *BookmarkDAO) Update(bookmark *model.Bookmark) error {
	// 忽略 created_at 和 updated_at 字段
	return dao.DB.Omit("CreatedAt", "UpdatedAt").Save(bookmark).Error
}

func (dao *BookmarkDAO) Delete(id uint) error {
	return dao.DB.Delete(&model.Bookmark{}, id).Error
}

func (dao *BookmarkDAO) GetBookmarksByFolderID(folderID uint) ([]*model.Bookmark, error) {
	var bookmarks []*model.Bookmark
	if err := dao.DB.Where("folder_id = ?", folderID).Find(&bookmarks).Error; err != nil {
		return nil, err
	}
	return bookmarks, nil
}

func (dao *BookmarkDAO) GetAllBookmarks() ([]*model.Bookmark, error) {
	var bookmarks []*model.Bookmark
	if err := dao.DB.Find(&bookmarks).Error; err != nil {
		return nil, err
	}
	return bookmarks, nil
}
