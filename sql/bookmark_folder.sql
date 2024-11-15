CREATE TABLE `bookmark_folder`
(
    `id`         int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `name`       varchar(255)     NOT NULL DEFAULT '' COMMENT '文件夹名称',
    `parent_id`  int(11)          NOT NULL DEFAULT '0' COMMENT '父文件夹ID',
    `created_at` datetime         NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime         NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) COMMENT ='书签文件夹信息';
