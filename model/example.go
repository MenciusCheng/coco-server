package model

import "time"

// 示例
type Example struct {
	// 主键
	Id int64 `gorm:"column:id" json:"id" xlsx:"主键"`
	// 配置名称
	ConfName string `gorm:"column:conf_name" json:"confName" xlsx:"配置名称"`
	// 活动ID
	ActId int32 `gorm:"column:act_id" json:"actId" xlsx:"活动ID"`
	// 创建时间
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt" xlsx:"创建时间"`
	// 更新时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt" xlsx:"更新时间"`
}

// get table name of Example
func (obj *Example) TableName() string {
	return "example"
}

/*
CREATE TABLE `example`
(
   `id` int(11) AUTO_INCREMENT COMMENT '主键',
   `conf_name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '配置名称',
   `act_id` INT(11) NOT NULL DEFAULT '0' COMMENT '活动id',
   `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
   `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
   PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8mb4 COMMENT='示例';
*/
