CREATE TABLE `tb_comment`
(
    `id`            char(36)    NOT NULL,
    `created_at`    datetime(3) NOT NULL,
    `updated_at`    datetime(3) NOT NULL,
    `deleted_at`    datetime(3)          DEFAULT NULL,

    `user_id`       char(36)    not null default '',
    `comment_id`    char(36)    not null default '',
    `target_id`     char(36)    not null default '',

    `content`       nvarchar(256)   not null default '',
    `like_count`    int(11)     not null default 0,
    `dislike_count` int(11)     not null default 0,
    `is_check`      tinyint(1)  not null default 0,

    PRIMARY KEY (`id`),
    KEY `idx_tb_user_master_deleted_at` (`deleted_at`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
