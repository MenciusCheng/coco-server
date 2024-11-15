package dao

import (
	"coco-server/middleware/db"
	"coco-server/model"
	"context"
	"fmt"
)

type bookmarkFolderDao struct {
	table string
}

var BookmarkFolderDB = &bookmarkFolderDao{
	table: (&model.BookmarkFolder{}).TableName(),
}

type QueryBookmarkFolderFilter struct {
	Name     string `json:"name"`
	ParentId int64  `json:"parentId"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

func (dao *bookmarkFolderDao) Query(ctx context.Context, filter *QueryBookmarkFolderFilter) ([]*model.BookmarkFolder, int64, error) {
	var (
		records = make([]*model.BookmarkFolder, 0)
		count   int64
	)

	query := db.MySQLCon.Table(dao.table)
	if filter.Name != "" {
		query = query.Where("name=?", filter.Name)
	}
	if filter.ParentId > 0 {
		query = query.Where("parent_id=?", filter.ParentId)
	}

	err := query.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	err = query.Order("id desc").Offset(filter.Offset).Limit(filter.Limit).Scan(&records).Error
	if err != nil {
		return nil, 0, err
	}
	return records, count, nil
}

func (dao *bookmarkFolderDao) Create(ctx context.Context, v *model.BookmarkFolder) (int64, error) {
	err := db.MySQLCon.Omit("created_at", "updated_at").Create(&v).Error
	return v.Id, err
}

func (dao *bookmarkFolderDao) Save(ctx context.Context, v *model.BookmarkFolder) (int64, error) {
	err := db.MySQLCon.Omit("created_at", "updated_at").Save(&v).Error
	return v.Id, err
}

func (dao *bookmarkFolderDao) Update(ctx context.Context, v *model.BookmarkFolder) error {
	updates := map[string]interface{}{
		"name":      v.Name,
		"parent_id": v.ParentId,
	}
	err := db.MySQLCon.Table(dao.table).
		Where("id=?", v.Id).
		UpdateColumns(updates).
		Error
	return err
}

func (dao *bookmarkFolderDao) Delete(ctx context.Context, id int64) error {
	sql := fmt.Sprintf("DELETE FROM %s WHERE `id`=?", dao.table)
	return db.MySQLCon.Exec(sql, id).Error
}
