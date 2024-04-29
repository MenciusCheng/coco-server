package model

import "time"

// 数据转换
type DataConvert struct {
	// ID
	Id int64 `gorm:"column:id" json:"id" xlsx:"ID"`
	// 配置名称
	ConfName string `gorm:"column:conf_name" json:"confName" xlsx:"配置名称"`
	// 配置内容
	ConfContent string `gorm:"column:conf_content" json:"confContent" xlsx:"配置内容"`
	// 创建时间
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt" xlsx:"创建时间"`
	// 更新时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt" xlsx:"更新时间"`
}

// get table name of DataConvert
func (obj *DataConvert) TableName() string {
	return "data_convert"
}

/*
CREATE TABLE `data_convert` (
   `id` int(11) AUTO_INCREMENT COMMENT '主键',
   `conf_name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '配置名称',
   `conf_content` json DEFAULT NULL COMMENT '配置内容',
   `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
   `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
   PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8mb4 COMMENT='数据转换';
*/

type DataConvertGetListReq struct {
	ConfName string `json:"confName"` // 配置名称
	Page     int    `json:"page"`     // 页码
	Size     int    `json:"size"`     // 每页大小
}

type DataConvertGenReq struct {
	DataSourceType string `json:"dataSourceType"` // 数据类型
	DataSource     string `json:"dataSource"`     // 数据源
	Template       string `json:"template"`       // 模板内容
}

type DataConvertGenRes struct {
	List []DataConvertGenData `json:"list"`
}

type DataConvertGenData struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}
