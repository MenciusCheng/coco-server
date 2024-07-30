package dao

import (
	"coco-server/middleware/db"
	"coco-server/model"
	"context"
	"fmt"
)

type actJsonConfDao struct {
	table string
}

var ActJsonConfDB = &actJsonConfDao{
	table: (&model.ActJsonConf{}).TableName(),
}

type QueryActJsonConfFilter struct {
	Name   string `json:"name"`
	Extend string `json:"extend"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
}

func (dao *actJsonConfDao) Query(ctx context.Context, filter *QueryActJsonConfFilter) ([]*model.ActJsonConf, int64, error) {
	var (
		records = make([]*model.ActJsonConf, 0)
		count   int64
	)

	query := db.MySQLCon.Table(dao.table)
	if filter.Name != "" {
		query = query.Where("name=?", filter.Name)
	}
	if filter.Extend != "" {
		query = query.Where("extend=?", filter.Extend)
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

func (dao *actJsonConfDao) Create(ctx context.Context, v *model.ActJsonConf) (int64, error) {
	err := db.MySQLCon.Omit("created_at", "updated_at").Create(&v).Error
	return v.Id, err
}

func (dao *actJsonConfDao) Save(ctx context.Context, v *model.ActJsonConf) (int64, error) {
	err := db.MySQLCon.Omit("created_at", "updated_at").Save(&v).Error
	return v.Id, err
}

func (dao *actJsonConfDao) Update(ctx context.Context, v *model.ActJsonConf) error {
	updates := map[string]interface{}{
		"name":   v.Name,
		"extend": v.Extend,
	}
	err := db.MySQLCon.Table(dao.table).
		Where("id=?", v.Id).
		UpdateColumns(updates).
		Error
	return err
}

func (dao *actJsonConfDao) Delete(ctx context.Context, id int64) error {
	sql := fmt.Sprintf("DELETE FROM %s WHERE `id`=?", dao.table)
	return db.MySQLCon.Exec(sql, id).Error
}
