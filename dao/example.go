package dao

import (
	"coco-server/middleware/db"
	"coco-server/model"
	"context"
)

type exampleDao struct {
	table string
}

var ExampleDB = &exampleDao{table: (&model.Example{}).TableName()}

// query all data from mysql table example
// @param ctx context.Context
func (dao *exampleDao) QueryExampleAll(ctx context.Context) ([]*model.Example, error) {
	records := make([]*model.Example, 0)
	err := db.MySQLCon.Table(dao.table).Find(&records).Error
	return records, err
}

type QueryExampleFilter struct {
	ConfName string `json:"confName"`
	ActId    int32  `json:"actId"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

// query model Example records from mysql table example
// @param ctx context.Context
// @param filter query params
func (dao *exampleDao) QueryExampleRecords(ctx context.Context, filter *QueryExampleFilter) ([]*model.Example, int64) {
	var (
		records = make([]*model.Example, 0)
		count   int64
	)

	query := db.MySQLCon.Table(dao.table)
	if filter.ConfName != "" {
		query = query.Where("conf_name=?", filter.ConfName)
	}
	if filter.ActId > 0 {
		query = query.Where("act_id=?", filter.ActId)
	}

	query.Count(&count)
	query.Order("id desc").Offset(filter.Offset).Limit(filter.Limit).Find(&records)
	return records, count
}

// insert model Example to mysql table example
// @param ctx context.Context
func (dao *exampleDao) InsertExample(ctx context.Context, v *model.Example) error {
	sql := "INSERT INTO `example` (`conf_name`,`act_id`) VALUES (?,?) "
	return db.MySQLCon.Exec(sql, v.ConfName, v.ActId).Error
}

// update mysql table example use model Example
// @param ctx context.Context
func (dao *exampleDao) UpdateExample(ctx context.Context, v *model.Example) error {
	sql := "UPDATE `example` SET `conf_name`=?, `act_id`=? WHERE `id`=? "
	return db.MySQLCon.Exec(sql, v.ConfName, v.ActId, v.Id).Error
}

// insert model Example to mysql table example
// @param ctx context.Context
func (dao *exampleDao) InsertOrUpdateExample(ctx context.Context, v *model.Example) error {
	sql := "INSERT INTO `example` (`conf_name`,`act_id`) VALUES (?,?)  ON DUPLICATE KEY UPDATE " +
		"`conf_name`=VALUES(`conf_name`),`act_id`=VALUES(`act_id`)"
	return db.MySQLCon.Exec(sql, v.ConfName, v.ActId).Error
}

// delete mysql table example record
// @param ctx context.Context
// @param id int64
func (dao *exampleDao) DeleteExample(ctx context.Context, id int64) error {
	sql := "DELETE FROM `example` WHERE `id`=?"
	return db.MySQLCon.Exec(sql, id).Error
}
