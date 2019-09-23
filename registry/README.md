#### Database

```
DROP TABLE IF EXISTS `device_info`;
CREATE TABLE `device_info` (
  `id` int NOT NULL AUTO_INCREMENT,
  `product_key` varchar(64) COLLATE utf8_bin NOT NULL COMMENT '产品KEY',
  `device_name` varchar(64) COLLATE utf8_bin NOT NULL COMMENT '设备名',
  `device_secret` varchar(64) COLLATE utf8_bin NOT NULL COMMENT '设备密钥',
  `device_model` varchar(128) COLLATE utf8_bin NOT NULL DEFAULT '0' COMMENT '型号',
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `last_active_at` timestamp NULL DEFAULT NULL COMMENT '上次活跃时间',
  `apply_id` varchar(64) COLLATE utf8_bin DEFAULT NULL COMMENT '申请ID',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '0-offline,1-online',
  `delete_flag` tinyint(1) NOT NULL DEFAULT 0 COMMENT '0-valid,1-deleted',
  PRIMARY KEY (`id`),
  UNIQUE KEY `device_name` (`device_name`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
```
