CREATE TABLE `user`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `name`  varchar(128) NOT NULL DEFAULT '' COMMENT 'UserName',
    `password`   varchar(128) NOT NULL DEFAULT '' COMMENT 'Password',
    `title`   varchar(128) NOT NULL DEFAULT '' COMMENT 'Title',
    `avatar_url`  varchar(256) NOT NULL DEFAULT '' COMMENT 'AVATAR IMAGE URL',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'User account create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'User account update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'User account delete time',
    PRIMARY KEY (`id`),
    KEY          `idx_user_name` (`user_name`) COMMENT 'UserName index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='User account table';

CREATE TABLE `video`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `user_id`    int(64) NOT NULL DEFAULT 0 COMMENT 'UserID',
    `play_url`      varchar(256) NOT NULL DEFAULT '' COMMENT 'Video Play URL',
    `cover_url`      varchar(256) NOT NULL DEFAULT '' COMMENT 'Video Cover URL',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Video create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Video update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'Video delete time',
    PRIMARY KEY (`id`),
    KEY `idx_user_id` (`user_id`) COMMENT 'UserID  index'
    KEY `idx_create_time` (`create_at` DESC) COMMENT 'CreateTime index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Video table';

CREATE TABLE `comment`
(
    `id`    int(64) NOT NULL DEFAULT 0 COMMENT 'comment id',
    `video_id`    int(64) NOT NULL DEFAULT 0 COMMENT 'belong video id',
    `comment_user_id`    int(64) NOT NULL DEFAULT 0 COMMENT 'Comment UserID',
    `content`    TEXT NULL COMMENT 'Content',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Comment create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Comment update time',
    PRIMARY KEY (`id`)
    KEY `idx_video_id` (`video_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Comment table';

CREATE TABLE `follow`
(
    `follow_id`    int(64) NOT NULL DEFAULT 0 COMMENT 'Follow id',
    `follower_id`    int(64) NOT NULL DEFAULT 0 COMMENT 'Follower id, be follow',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Follow create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Follow update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'follow delete time',
    PRIMARY KEY (`follow_id`, `follower_id`)
    KEY `idx_follow_id` (`follow_id`)
    KEY `idx_follower_id` (`follower_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Follow table';

CREATE TABLE `favoriate`
(
    `video_id`    int(64) NOT NULL DEFAULT 0 COMMENT 'Video ID',
    `user_id`    int(64) NOT NULL DEFAULT 0 COMMENT 'Follower ID',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Favirate create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Favirate update time',
    PRIMARY KEY (`video_id`, `user_id`)
    KEY `idx_video_id` (`video_id`)
    KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Favoriate table';

CREATE TABLE `friend`
(
    `pri_friend_id`    int(64) NOT NULL DEFAULT 0 COMMENT 'Primary Friend USER ID',
    `sec_video_id`    int(64) NOT NULL DEFAULT 0 COMMENT 'Second Friend USER ID',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Friend create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Friend update time',
    PRIMARY KEY (`pri_friend_id`, `sec_friend_id`)
    KEY `idx_pri_id` (`pri_friend_id`)
    KEY `idx_sec_id` (`sec_friend_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Friend table';
