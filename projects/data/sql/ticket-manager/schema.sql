/*
-- Query: SELECT * FROM ticket.users
LIMIT 0, 1000

-- Date: 2021-05-24 23:49
*/

CREATE TABLE IF NOT EXISTS `ticket`.`users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` longtext COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `password` longtext COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `email` longtext COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `age` bigint(20) DEFAULT NULL,
  `birthday` datetime(3) DEFAULT NULL,
  `member_number` longtext COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `role` longtext COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


CREATE TABLE IF NOT EXISTS `ticket`.`movies` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` longtext COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `image` longtext COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `actors` longtext COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `mtype` longtext COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `minfo` longtext COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `mtime` longtext COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_movies_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


CREATE TABLE IF NOT EXISTS `ticket`.`infos` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` longtext COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `title` longtext COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `content` text COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `mtype` longtext COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `minfo` longtext COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `mtime` longtext COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `image` longtext COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_infos_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
SELECT * FROM ticket.shows;


CREATE TABLE IF NOT EXISTS `ticket`.duo`shows` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` longtext COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `image` longtext COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `actors` longtext COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `mtype` longtext COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `minfo` longtext COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `mtime` longtext COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_shows_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


