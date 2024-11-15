package dao

import (
	"coco-server/middleware/db"
	"coco-server/model"
	"context"
	"fmt"
)

type bookmarkDao struct {
	table string
}

var BookmarkDB = &bookmarkDao{
	table: (&model.Bookmark{}).TableName(),
}

type QueryBookmarkFilter struct {
	Name     string `json:"name"`
	FolderId int64  `json:"folderId"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

func (dao *bookmarkDao) Query(ctx context.Context, filter *QueryBookmarkFilter) ([]*model.Bookmark, int64, error) {
	var (
		records = make([]*model.Bookmark, 0)
		count   int64
	)

	query := db.MySQLCon.Table(dao.table)
	if filter.Name != "" {
		query = query.Where("name=?", filter.Name)
	}
	if filter.FolderId > 0 {
		query = query.Where("folder_id=?", filter.FolderId)
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

func (dao *bookmarkDao) Create(ctx context.Context, v *model.Bookmark) (int64, error) {
	err := db.MySQLCon.Omit("created_at", "updated_at").Create(&v).Error
	return v.Id, err
}

func (dao *bookmarkDao) Save(ctx context.Context, v *model.Bookmark) (int64, error) {
	err := db.MySQLCon.Omit("created_at", "updated_at").Save(&v).Error
	return v.Id, err
}

func (dao *bookmarkDao) Update(ctx context.Context, v *model.Bookmark) error {
	updates := map[string]interface{}{
		"name":      v.Name,
		"url":       v.Url,
		"icon":      v.Icon,
		"remark":    v.Remark,
		"folder_id": v.FolderId,
	}
	err := db.MySQLCon.Table(dao.table).
		Where("id=?", v.Id).
		UpdateColumns(updates).
		Error
	return err
}

func (dao *bookmarkDao) Delete(ctx context.Context, id int64) error {
	sql := fmt.Sprintf("DELETE FROM %s WHERE `id`=?", dao.table)
	return db.MySQLCon.Exec(sql, id).Error
}
