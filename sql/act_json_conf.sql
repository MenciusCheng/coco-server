CREATE TABLE `act_json_conf`
(
    `id`         int(11) AUTO_INCREMENT COMMENT 'ID',
    `name`       varchar(255) NOT NULL DEFAULT '' COMMENT '名称',
    `extend`     json COMMENT '扩展配置',
    `created_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='活动JSON配置';
