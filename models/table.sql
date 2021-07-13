DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`(
                       `id` bigint(20) NOT NULL AUTO_INCREMENT,
                       `user_id` bigint(20) NOT NULL,
                       `username` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
                       `password` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
                       `email` varchar(64) COLLATE utf8mb4_general_ci,
                       `gender` tinyint(4) NOT NULL DEFAULT 0,
                       `create_time` timestamp NULL DEFAULI CURRENT TIMESTAMP,
                       `update_time` timestamp NULL DEFAULT CURRENTTIMESTAMP ON UPDATECURRENT_TIMESTAMP,
                       PRIMARY KEY (`id`),
                       UNIQUE KEY `idx_username` (`username`) USING BTREE,
                       UNIQUE KEY `idx user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE-utf8mb4_general_ci;


DROP TABLE IF EXISTS "community";
CREATE TABLE `community`
(
    `id`             int(11) not null AUTO_INCREMENT,
    `community_id`   int(10) unsigned NOT NULL,
    `community_name` varchar(128) COLLATE utf8mb4_general_ci NOT NULL,
    `introduction`   varchar(256) COLLATE utf8mb4_general_ci NOT NULL,
    `create_time`    timestamp                               NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time`    timestamp                               NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`1d`),
    UNIQUE KEY `idx_comnunity_id` (`community_id`),
    UNIQUE KEY "idx_comnunity_name" (`comnunity_name`)
)ENGINE=InnoOB DEFAULT CHARSET=utf8mb4 COLLATE = utf8mb4_general_ci;