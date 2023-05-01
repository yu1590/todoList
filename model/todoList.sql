CREATE TABLE todo_list (
                           `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键 id',
                           `account_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '账户 id',
                           `time`  timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '日程时间',
                           `month` tinyint  NOT NULL DEFAULT 0 COMMENT '月份',
                           `extra` text  COMMENT '日程详细信息',
                           `status` TINYINT  NOT NULL DEFAULT '0' COMMENT '状态 0:未完成 1:已完成',
                           PRIMARY KEY (`id`)
)