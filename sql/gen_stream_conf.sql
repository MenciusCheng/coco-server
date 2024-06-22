CREATE TABLE `gen_stream_conf`
(
    `id`         int(11) AUTO_INCREMENT COMMENT 'ID',
    `conf_name`  varchar(255) NOT NULL DEFAULT '' COMMENT '配置名称',
    `extend`     json COMMENT '扩展配置',
    `created_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='流式生成器配置';
