package dao

import (
	"coco-server/middleware/db"
	"coco-server/model"
	"context"
	"fmt"
)

type genStreamConfDao struct {
	table string
}

var GenStreamConfDB = &genStreamConfDao{
	table: (&model.GenStreamConf{}).TableName(),
}

type QueryGenStreamConfFilter struct {
	ConfNameLike string `json:"confNameLike"`
	Offset       int    `json:"offset"`
	Limit        int    `json:"limit"`
}

func (dao *genStreamConfDao) Query(ctx context.Context, filter *QueryGenStreamConfFilter) ([]*model.GenStreamConf, int64, error) {
	var (
		records = make([]*model.GenStreamConf, 0)
		count   int64
	)

	query := db.MySQLCon.Table(dao.table)
	if filter.ConfNameLike != "" {
		query = query.Where("conf_name like ?", "%"+filter.ConfNameLike+"%")
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

func (dao *genStreamConfDao) Create(ctx context.Context, v *model.GenStreamConf) (int64, error) {
	err := db.MySQLCon.Omit("created_at", "updated_at").Create(&v).Error
	return v.Id, err
}

func (dao *genStreamConfDao) Save(ctx context.Context, v *model.GenStreamConf) (int64, error) {
	err := db.MySQLCon.Omit("created_at", "updated_at").Save(&v).Error
	return v.Id, err
}

func (dao *genStreamConfDao) Update(ctx context.Context, v *model.GenStreamConf) error {
	updates := map[string]interface{}{
		"conf_name": v.ConfName,
		"extend":    v.Extend,
	}
	err := db.MySQLCon.Table(dao.table).
		Where("id=?", v.Id).
		UpdateColumns(updates).
		Error
	return err
}

func (dao *genStreamConfDao) Delete(ctx context.Context, id int64) error {
	sql := fmt.Sprintf("DELETE FROM %s WHERE `id`=?", dao.table)
	return db.MySQLCon.Exec(sql, id).Error
}
