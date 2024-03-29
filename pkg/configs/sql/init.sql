CREATE TABLE `user`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `name`  varchar(128) NOT NULL DEFAULT '' COMMENT 'UserName',
    `password`   varchar(128) NOT NULL DEFAULT '' COMMENT 'Password',
    `title`   varchar(512) NOT NULL DEFAULT '' COMMENT 'Title',
    `avatar_url`  varchar(2048) NOT NULL DEFAULT '' COMMENT 'AVATAR IMAGE URL',
    `background_image`  varchar(2048) NOT NULL DEFAULT '' COMMENT 'Background IMAGE URL',
    `signature`  varchar(512) NOT NULL DEFAULT '' COMMENT 'SIGNATURE',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'User account create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'User account update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'User account delete time',
    PRIMARY KEY (`id`),
    UNIQUE KEY          `idx_user_name` (`name`) COMMENT 'UserName index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='User account table';

CREATE TABLE `video`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `user_id`    bigint unsigned NOT NULL DEFAULT 0 COMMENT 'UserID',
    `play_url`      varchar(2048) NOT NULL DEFAULT '' COMMENT 'Video Play URL',
    `title`      varchar(2048) NOT NULL DEFAULT '' COMMENT 'Video Title',
    `cover_url`      varchar(2048) NOT NULL DEFAULT '' COMMENT 'Video Cover URL',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Video create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Video update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'Video delete time',
    PRIMARY KEY (`id`),
    KEY `idx_user_id` (`user_id`) COMMENT 'UserID  index',
    KEY `idx_create_time` (`created_at` DESC) COMMENT 'CreateTime index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Video table';

CREATE TABLE `comments`
(
    `id`    bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'comment id',
    `video_id`    bigint unsigned NOT NULL DEFAULT 0 COMMENT 'belong video id',
    `user_id`   bigint unsigned NOT NULL DEFAULT 0 COMMENT 'Comment UserID',
    `content`    TEXT NOT NULL COMMENT 'Content',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Comment create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Comment update time',
    PRIMARY KEY (`id`),
    KEY `idx_video_id` (`video_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Comment table';

CREATE TABLE `follow`
(
    `id`    bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'follow id',
    `follow_id`    bigint unsigned NOT NULL DEFAULT 0 COMMENT 'Follow User Id',
    `follower_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT 'Follower User Id, be follow',
    `created_at` timestamp NOT NULL DEFAULT CURENT_TIMESTAMP COMMENT 'Follow create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Follow update time',
    PRIMARY KEY (`id`),
    KEY `idx_follow_id` (`follow_id`),
    KEY `idx_follower_id` (`follower_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Follow table';

CREATE TABLE `favoriate`
(
    `id`    bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'favoriate id',
    `video_id`    bigint unsigned NOT NULL DEFAULT 0 COMMENT 'Video ID',
    `user_id`    bigint unsigned NOT NULL DEFAULT 0 COMMENT 'Follower ID',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Favirate create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Favirate update time',
    PRIMARY KEY (`id`),
    KEY `idx_video_id` (`video_id`),
    KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Favoriate table';

CREATE TABLE `friend`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `primary_friend_id`  bigint unsigned NOT NULL DEFAULT 0 COMMENT 'Primary Friend USER ID',
    `second_friend_id`    bigint unsigned NOT NULL DEFAULT 0 COMMENT 'Second Friend USER ID',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Friend create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Friend update time',
    PRIMARY KEY (`id`),
    KEY `idx_pri_id` (`primary_friend_id`),
    KEY `idx_sec_id` (`second_friend_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Friend table';

CREATE TABLE `chat_message`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `to_user_id`    bigint unsigned NOT NULL DEFAULT 0 COMMENT 'TO  USER ID',
    `from_user_id`    bigint unsigned NOT NULL DEFAULT 0 COMMENT 'FROM USER ID',
    `content`    TEXT NULL COMMENT 'Content',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'chat create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Chat update time',
    PRIMARY KEY (`id`),
    KEY `idx_user_id` (`to_user_id`, `from_user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Chat Msg table';
