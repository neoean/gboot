# SET PERSIST EXPLICIT_DEFAULTS_FOR_TIMESTAMP=OFF;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`          BIGINT(11)   NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `user_name`   VARCHAR(32)  NOT NULL COMMENT '自定义用户名',
    `nick_name`   VARCHAR(32)  NOT NULL COMMENT '微信用户名',
    `real_name`   VARCHAR(32)  NOT NULL DEFAULT '' COMMENT '真实姓名',
    `gender`      TINYINT(6)   NOT NULL DEFAULT 0 COMMENT '性别 0-未知, 1-男, 2-女',
    `avatar_url`  VARCHAR(256) NOT NULL DEFAULT '' COMMENT '头像地址',
    `open_id`     VARCHAR(128) NOT NULL DEFAULT '' COMMENT 'open_id',
    `union_id`    VARCHAR(128) NOT NULL DEFAULT '' COMMENT 'union_id',
    `session_key` VARCHAR(128) NOT NULL DEFAULT '' COMMENT 'session_key',
    `country`     VARCHAR(32)  NOT NULL DEFAULT '' COMMENT 'country',
    `province`    VARCHAR(32)  NOT NULL DEFAULT '' COMMENT 'province',
    `city`        VARCHAR(32)  NOT NULL DEFAULT '' COMMENT 'city',
    `coin`        BIGINT(20)   NOT NULL DEFAULT 0 COMMENT '积分',
    `balance`     BIGINT(20)   NOT NULL DEFAULT 0 COMMENT '余额 单位分',
    `settled`     BIGINT(20)   NOT NULL DEFAULT 0 COMMENT '已结算金额 单位分',
    `status`      TINYINT(6)   NOT NULL DEFAULT 1 COMMENT '状态 1-启用, 2-停用',
    `created_at`  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `delete_at`   DATETIME     NOT NULL DEFAULT '2000-01-01 00:00:00' COMMENT '删除时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_open_id` (`open_id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 100000
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户表';
