package dao

import (
	"coco-server/middleware/db"
	"coco-server/model"
	"context"
)

type dataConvertDao struct {
	table string
}

var DataConvertDB = &dataConvertDao{table: (&model.DataConvert{}).TableName()}

type QueryDataConvertFilter struct {
	ConfName string `json:"confName"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

// query model DataConvert records from mysql table data_convert
// @param ctx context.Context
// @param filter query params
func (dao *dataConvertDao) QueryDataConvertRecords(ctx context.Context, filter *QueryDataConvertFilter) ([]*model.DataConvert, int64) {
	var (
		records = make([]*model.DataConvert, 0)
		count   int64
	)

	query := db.MySQLCon.Table(dao.table)
	if filter.ConfName != "" {
		query = query.Where("conf_name=?", filter.ConfName)
	}

	query.Count(&count)
	query.Order("id desc").Offset(filter.Offset).Limit(filter.Limit).Find(&records)
	return records, count
}

// insert model DataConvert to mysql table data_convert
// @param ctx context.Context
func (dao *dataConvertDao) InsertDataConvert(ctx context.Context, v *model.DataConvert) error {
	sql := "INSERT INTO `data_convert` (`conf_name`,`conf_content`) VALUES (?,?)"
	return db.MySQLCon.Exec(sql, v.ConfName, v.ConfContent).Error
}

func (dao *dataConvertDao) Create(ctx context.Context, v *model.DataConvert) (int64, error) {
	err := db.MySQLCon.Omit("created_at", "updated_at").Create(&v).Error
	return v.Id, err
}

// update mysql table data_convert use model DataConvert
// @param ctx context.Context
func (dao *dataConvertDao) UpdateDataConvert(ctx context.Context, v *model.DataConvert) error {
	sql := "UPDATE `data_convert` SET `conf_name`=?, `conf_content`=? WHERE `id`=?"
	return db.MySQLCon.Exec(sql, v.ConfName, v.ConfContent, v.Id).Error
}

// delete mysql table data_convert record
// @param ctx context.Context
// @param id int64
func (dao *dataConvertDao) DeleteDataConvert(ctx context.Context, id int64) error {
	sql := "DELETE FROM `data_convert` WHERE `id`=?"
	return db.MySQLCon.Exec(sql, id).Error
}
