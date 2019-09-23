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
  UNIQUE KEY `product` (`product_key`),
  UNIQUE KEY `tenant_product_name` (`tenant_id`, `product_name`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;
```
