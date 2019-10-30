#### Database

##### Product Table

```text
DROP TABLE IF EXISTS `product_info`;
CREATE TABLE `product_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `product_key` varchar(64) NOT NULL,
  `product_name` varchar(64) NOT NULL,
  `product_secret` varchar(64) NOT NULL,
  `description` varchar(256) NOT NULL DEFAULT '',
  `access_points` varchar(256) NOT NULL,
  `device_count` int(32) NOT NULL DEFAULT 0,
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `delete_flag` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0-valid, 1-deleted',
  PRIMARY KEY (`id`),
  UNIQUE KEY `product` (`product_key`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;
```


##### Device Info Table

```text
DROP TABLE IF EXISTS `device_info`;
CREATE TABLE `device_info` (
  `id` int(32) NOT NULL AUTO_INCREMENT,
  `product_key` varchar(64) COLLATE utf8_bin NOT NULL COMMENT '产品KEY',
  `device_secret` varchar(64) COLLATE utf8_bin NOT NULL COMMENT '设备密钥',
  `device_name` varchar(64) COLLATE utf8_bin NOT NULL COMMENT '设备名',
  `model` varchar(128) COLLATE utf8_bin NOT NULL DEFAULT '0' COMMENT '型号',
  `product_version` varchar(16) COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '产品版本',
  `sdk_version` varchar(32) COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT 'sdk版本',
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `last_active_at` timestamp NULL DEFAULT NULL COMMENT '上次活跃时间',
  `apply_id` varchar(64) COLLATE utf8_bin DEFAULT NULL COMMENT '申请ID',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '0-offline,1-oneline',
  `delete_flag` tinyint(1) NOT NULL DEFAULT 0 COMMENT '0-valid,1-deleted',
  PRIMARY KEY (`id`),
  UNIQUE KEY `device_name` (`device_name`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;
```

##### Topic Subscribe
```text
DROP TABLE IF EXISTS `topic_subscription`;
CREATE TABLE `topic_subscription` (
  `id` int(32) NOT NULL AUTO_INCREMENT=1,
  `product_key` varchar(64) COLLATE utf8_bin NOT NULL COMMENT '产品KEY',
  `guid` int(32) NOT NULL COMMENT 'device guid',
  `topic_filter` varchar(128) NOT NULL,
  `qos` tinyint(4) NOT NULL DEFAULT 0 COMMENT '0-qos0, 1-qos1',
  `topic_type` tinyint(4) NOT NULL DEFAULT 0 COMMENT '0-custom, 1-system, 2-ota',
  `subscribe_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `delete_flag` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0-valid, 1-deleted',
  PRIMARY KEY (`id`),
  UNIQUE KEY `device_topic_id` (`topic_filter`, `guid`),
  KEY `device_guid` (`guid`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;
```