-- ----------------------------
-- Table structure for device
-- ----------------------------
DROP TABLE IF EXISTS `device`;
CREATE TABLE `device`
(
    `id`             bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id`        varchar(512)  NOT NULL COMMENT '用户业务唯一id',
    `type`           tinyint(3) NOT NULL COMMENT '设备类型,1:Android；2：IOS；3：Windows; 4：MacOS；5：Web',
    `brand`          varchar(20) NOT NULL COMMENT '手机厂商',
    `model`          varchar(20) NOT NULL COMMENT '机型',
    `system_version` varchar(10) NOT NULL COMMENT '系统版本',
    `sdk_version`    varchar(10) NOT NULL COMMENT 'app版本',
    `status`         tinyint(3) NOT NULL DEFAULT '0' COMMENT '在线状态，0：离线；1：在线',
    `conn_addr`      varchar(25) NOT NULL COMMENT '连接层服务器地址',
    `client_addr`    varchar(25) NOT NULL COMMENT '客户端地址',
    `create_time`    datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`    datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY              `idx_user_id` (`user_id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='设备';

-- ----------------------------
-- Table structure for friend
-- ----------------------------
DROP TABLE IF EXISTS `friend`;
CREATE TABLE `friend`
(
    `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `user_id`     varchar(256) NOT NULL COMMENT '用户业务唯一id',
    `friend_id`   varchar(256) NOT NULL COMMENT '好友用户业务唯一id',
    `remarks`     varchar(20)   NOT NULL COMMENT '备注',
    `extra`       varchar(1024) NOT NULL COMMENT '附加属性',
    `status`      tinyint(4) NOT NULL COMMENT '状态，1：申请，2：同意',
    `create_time` datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_user_id_friend_id` (`user_id`, `friend_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='好友';

-- ----------------------------
-- Table structure for group
-- ----------------------------
DROP TABLE IF EXISTS `group`;
CREATE TABLE `group`
(
    `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `group_id`     varchar(255)  NOT NULL COMMENT '群组的业务id',
    `name`         varchar(50)   NOT NULL COMMENT '群组名称',
    `avatar_url`   varchar(255)  NOT NULL COMMENT '群组头像',
    `introduction` varchar(255)  NOT NULL COMMENT '群组简介',
    `user_num`     int(11) NOT NULL DEFAULT '0' COMMENT '群组人数',
    `extra`        varchar(1024) NOT NULL COMMENT '附加属性',
    `create_time`  datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`  datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `create_use_id` varchar(255)  NOT NULL default '' COMMENT '创建者id, 如果是系统创建就是空',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_group_id` (`group_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='群组';

-- ----------------------------
-- Table structure for group_user
-- ----------------------------
DROP TABLE IF EXISTS `group_user`;
CREATE TABLE `group_user`
(
    `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `group_id`    varchar(255)  NOT NULL COMMENT '群组的业务id',
    `user_id`      varchar(512)  NOT NULL COMMENT '用户业务唯一id',
    `member_type` tinyint(4) NOT NULL COMMENT '成员类型，1：管理员；2：普通成员',
    `remarks`     varchar(20)   NOT NULL COMMENT '备注',
    `extra`       varchar(1024) NOT NULL COMMENT '附加属性',
    `status`      tinyint(255) NOT NULL COMMENT '状态',
    `create_time` datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_group_id_user_id` (`group_id`, `user_id`) USING BTREE,
    KEY           `idx_user_id` (`user_id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='群组成员';

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `user_id`      varchar(512) NOT NULL COMMENT '用户业务唯一id',
    `nickname`     varchar(20)   NOT NULL COMMENT '昵称',
    `sex`          tinyint(4) NOT NULL COMMENT '性别，0:未知；1:男；2:女',
    `avatar_url`   varchar(256)  NOT NULL COMMENT '用户头像链接',
    `extra`        varchar(1024) NOT NULL COMMENT '附加属性',
    `source_code`  varchar(20) NOT NULL DEFAULT '' COMMENT '用户来源编码: 其他系统如果接入用户时候，定义产品对应的编码传过来  ',
    `phone`        varchar(100) COMMENT '手机号码',
    `create_time`  datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`  datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_user_id` (`user_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='用户';

CREATE TABLE `seq`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `object_type` tinyint  NOT NULL COMMENT '对象类型,1:用户；2：群组',
    `object_id`   varchar(256) NOT NULL COMMENT '对象id',
    `seq`         bigint unsigned NOT NULL COMMENT '序列号',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_object` (`object_type`,`object_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='序列号';


-- ----------------------------
-- Table structure for message
-- ----------------------------
DROP TABLE IF EXISTS `message`;
CREATE TABLE `message`
(
    `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `user_id`     varchar(512) NOT NULL COMMENT '用户业务唯一id',
    `request_id`  bigint(20) NOT NULL COMMENT '请求(消息)id',
    `code`        tinyint(4) NOT NULL COMMENT '消息类型',
    `content`     blob     NOT NULL COMMENT '消息内容',
    `seq`         bigint(20) unsigned NOT NULL COMMENT '消息序列号',
    `user_seq`    varchar(512)  COMMENT '消息唯一序列号',
    `send_time`   datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP (3) COMMENT '消息发送时间',
    `status`      tinyint(255) NOT NULL DEFAULT '0' COMMENT '消息状态，0：未处理1：消息撤回',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_user_id_seq` (`user_id`, `seq`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='消息';
