
SET FOREIGN_KEY_CHECKS=0;

CREATE TABLE `tb_article_master` (
  `id` char(36) NOT NULL,
  `created_at` datetime(3) NOT NULL,
  `updated_at` datetime(3) NOT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `title` varchar(256) NOT NULL,
  `description` varchar(512) NOT NULL,
  `content` text NOT NULL,
  `picture` varchar(512) NOT NULL,
  `sort` int(11) NOT NULL,
  `pv` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_tb_article_master_deleted_at` (`deleted_at`),
  KEY `idx_tb_article_master_sort` (`sort`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `tb_auth_permission` (
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(32) NOT NULL,
  `description` varchar(100) NOT NULL,
  `ref_method` varchar(255) NOT NULL,
  `ref_path` varchar(255) NOT NULL,
  `sort` int(11) NOT NULL,
  PRIMARY KEY (`name`),
  KEY `idx_tb_auth_permission_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `tb_auth_role` (
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(32) NOT NULL,
  `description` varchar(100) NOT NULL,
  `main_permission_name` varchar(32) NOT NULL,
  PRIMARY KEY (`name`),
  KEY `idx_tb_auth_role_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `tb_auth_role_permission` (
  `id` char(36) NOT NULL,
  `created_at` datetime(3) NOT NULL,
  `updated_at` datetime(3) NOT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `role_name` varchar(32) NOT NULL,
  `permission_name` varchar(32) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_tb_auth_role_permission_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `tb_auth_user_role` (
  `id` char(36) NOT NULL,
  `created_at` datetime(3) NOT NULL,
  `updated_at` datetime(3) NOT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `role_name` varchar(32) NOT NULL,
  `user_id` char(36) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_tb_auth_user_role_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `tb_carousel_master` (
  `id` char(36) NOT NULL,
  `created_at` datetime(3) NOT NULL,
  `updated_at` datetime(3) NOT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `message` varchar(128) NOT NULL,
  `url` varchar(512) NOT NULL,
  `is_on_show` tinyint(1) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_tb_carousel_master_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `tb_catchword` (
  `id` char(36) NOT NULL,
  `created_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `updated_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `deleted_at` datetime(3) DEFAULT NULL,
  `title` varchar(128) NOT NULL DEFAULT '',
  `content` varchar(2048) NOT NULL DEFAULT '',
  `sort` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_tb_catchword_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `tb_config_master` (
  `name` varchar(255) NOT NULL,
  `create_time` datetime NOT NULL,
  `value` text NOT NULL,
  PRIMARY KEY (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='配置表';

CREATE TABLE `tb_customer_account` (
  `id` char(36) NOT NULL,
  `created_at` datetime(3) NOT NULL,
  `updated_at` datetime(3) NOT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `customer_trace_id` char(36) NOT NULL,
  `login_name` varchar(256) NOT NULL,
  `password` varchar(256) NOT NULL,
  `last_access_ip` varchar(191) NOT NULL,
  `last_access_addr` varchar(191) NOT NULL,
  `mobile` varchar(32) NOT NULL,
  `email` varchar(32) NOT NULL,
  `other` varchar(1024) NOT NULL,
  `access_times` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_tb_customer_account_customer_trace_id` (`customer_trace_id`),
  KEY `idx_tb_customer_account_password` (`password`),
  KEY `idx_tb_customer_account_last_access_ip` (`last_access_ip`),
  KEY `idx_tb_customer_account_last_access_addr` (`last_access_addr`),
  KEY `idx_tb_customer_account_mobile` (`mobile`),
  KEY `idx_tb_customer_account_email` (`email`),
  KEY `idx_tb_customer_account_deleted_at` (`deleted_at`),
  KEY `idx_tb_customer_account_login_name` (`login_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `tb_dynamic_share` (
  `id` char(36) NOT NULL,
  `created_at` datetime(3) NOT NULL,
  `updated_at` datetime(3) NOT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `content` text NOT NULL,
  `sort` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_tb_dynamic_share_sort` (`sort`),
  KEY `idx_tb_dynamic_share_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `tb_dynamic_share_label` (
  `id` char(36) NOT NULL,
  `created_at` datetime(3) NOT NULL,
  `updated_at` datetime(3) NOT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `label,type:nvarchar(36)` varchar(191) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_tb_dynamic_share_label_deleted_at` (`deleted_at`),
  KEY `idx_tb_dynamic_share_label_label` (`label,type:nvarchar(36)`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `tb_file_master` (
  `id` char(36) NOT NULL,
  `created_at` datetime(3) NOT NULL,
  `updated_at` datetime(3) NOT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `key_name` varchar(200) NOT NULL,
  `bucket_name` varchar(200) NOT NULL,
  `url` varchar(200) NOT NULL,
  `user_id` char(36) NOT NULL,
  `content_type` varchar(50) NOT NULL,
  `biz_type` varchar(50) NOT NULL,
  `size` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_tb_file_master_key_name` (`key_name`),
  UNIQUE KEY `idx_tb_file_master_url` (`url`),
  KEY `idx_tb_file_master_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `tb_like_dislike` (
  `id` char(36) NOT NULL,
  `created_at` datetime(3) NOT NULL,
  `updated_at` datetime(3) NOT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` char(36) NOT NULL,
  `business_type` varchar(16) CHARACTER SET utf8 NOT NULL,
  `business_id` char(36) NOT NULL,
  `give_type` varchar(16) CHARACTER SET utf8 NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_tb_like_dislike_give_type` (`give_type`),
  KEY `idx_tb_like_dislike_deleted_at` (`deleted_at`),
  KEY `idx_tb_like_dislike_user_id` (`user_id`),
  KEY `idx_tb_like_dislike_business_type` (`business_type`),
  KEY `idx_tb_like_dislike_business_id` (`business_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `tb_message_email` (
  `id` char(36) NOT NULL,
  `created_at` datetime(3) NOT NULL,
  `updated_at` datetime(3) NOT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `send_time` datetime NOT NULL,
  `from` varchar(256) NOT NULL,
  `to` varchar(256) NOT NULL,
  `subject` varchar(256) NOT NULL,
  `content` varchar(1024) NOT NULL,
  `is_success` tinyint(1) NOT NULL,
  `other` varchar(1024) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_tb_message_email_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `tb_message_sms` (
  `id` char(36) NOT NULL,
  `created_at` datetime(3) NOT NULL,
  `updated_at` datetime(3) NOT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `send_time` datetime NOT NULL,
  `name` varchar(32) NOT NULL,
  `to` varchar(32) NOT NULL,
  `message` varchar(512) NOT NULL,
  `is_success` tinyint(1) NOT NULL,
  `other` varchar(1024) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_tb_message_sms_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `tb_message_valid_code` (
  `id` char(36) NOT NULL,
  `created_at` datetime(3) NOT NULL,
  `updated_at` datetime(3) NOT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `send_time` datetime NOT NULL,
  `phone_num` varchar(32) NOT NULL,
  `code` varchar(16) NOT NULL,
  `type` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_tb_message_valid_code_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `tb_stat_article_visitor_logs` (
  `id` char(36) NOT NULL,
  `created_at` datetime(3) NOT NULL,
  `updated_at` datetime(3) NOT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `access_time` datetime(3) NOT NULL,
  `ip` varchar(64) CHARACTER SET utf8 NOT NULL,
  `customer_trace_id` varchar(64) CHARACTER SET utf8 NOT NULL,
  `article_id` varchar(36) CHARACTER SET utf8 NOT NULL,
  `article_title` varchar(256) CHARACTER SET utf8 NOT NULL,
  `region_name` varchar(128) CHARACTER SET utf8 NOT NULL,
  `city` varchar(128) CHARACTER SET utf8 NOT NULL,
  `memo` varchar(512) CHARACTER SET utf8 NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_tb_stat_article_visitor_logs_ip` (`ip`),
  KEY `idx_tb_stat_article_visitor_logs_customer_trace_id` (`customer_trace_id`),
  KEY `idx_tb_stat_article_visitor_logs_region_name` (`region_name`),
  KEY `idx_tb_stat_article_visitor_logs_city` (`city`),
  KEY `idx_tb_stat_article_visitor_logs_deleted_at` (`deleted_at`),
  KEY `idx_tb_stat_article_visitor_logs_access_time` (`access_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `tb_stat_visitor_logs` (
  `id` char(36) NOT NULL,
  `created_at` datetime(3) NOT NULL,
  `updated_at` datetime(3) NOT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `access_time` datetime(3) NOT NULL,
  `ip` varchar(64) CHARACTER SET utf8 NOT NULL,
  `trace_id` varchar(64) CHARACTER SET utf8 NOT NULL,
  `region_name` varchar(128) CHARACTER SET utf8 NOT NULL,
  `city` varchar(128) CHARACTER SET utf8 NOT NULL,
  `memo` varchar(512) CHARACTER SET utf8 NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_tb_stat_visitor_logs_city` (`city`),
  KEY `idx_tb_stat_visitor_logs_deleted_at` (`deleted_at`),
  KEY `idx_tb_stat_visitor_logs_access_time` (`access_time`),
  KEY `idx_tb_stat_visitor_logs_ip` (`ip`),
  KEY `idx_tb_stat_visitor_logs_trace_id` (`trace_id`),
  KEY `idx_tb_stat_visitor_logs_region_name` (`region_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `tb_user_master` (
  `id` char(36) NOT NULL,
  `created_at` datetime(3) NOT NULL,
  `updated_at` datetime(3) NOT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `login_name` varchar(50) NOT NULL,
  `phone_num` varchar(20) NOT NULL,
  `email` varchar(100) NOT NULL,
  `password` varchar(64) NOT NULL,
  `picture` varchar(500) NOT NULL,
  `salt` varchar(8) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_tb_user_master_phone_num` (`phone_num`),
  KEY `idx_tb_user_master_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
