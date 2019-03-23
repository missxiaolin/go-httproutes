# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.7.24)
# Database: go_video
# Generation Time: 2019-03-23 05:34:48 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table comments
# ------------------------------------------------------------

DROP TABLE IF EXISTS `comments`;

CREATE TABLE `comments` (
  `id` varchar(64) NOT NULL DEFAULT '',
  `video_id` varchar(64) DEFAULT NULL,
  `author_id` int(10) unsigned DEFAULT NULL,
  `content` text,
  `time` date DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `comments` WRITE;
/*!40000 ALTER TABLE `comments` DISABLE KEYS */;

INSERT INTO `comments` (`id`, `video_id`, `author_id`, `content`, `time`)
VALUES
	('51958142-395f-473e-974b-6de77d523a9e','12345',1,'I like this video',NULL),
	('e9700e61-c7f5-47a6-b94c-49807ee5de66','12345',1,'I like this video',NULL);

/*!40000 ALTER TABLE `comments` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sessions
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sessions`;

CREATE TABLE `sessions` (
  `session_id` varchar(255) NOT NULL DEFAULT '',
  `TTL` varchar(255) DEFAULT NULL,
  `login_name` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`session_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `sessions` WRITE;
/*!40000 ALTER TABLE `sessions` DISABLE KEYS */;

INSERT INTO `sessions` (`session_id`, `TTL`, `login_name`)
VALUES
	('020d8a77-04a8-4979-a56d-01417ab32f26','1552881495400',''),
	('0e86f759-0ade-40e2-829e-40c0f40acb5a','1552879194363','xiaobei'),
	('14f9be82-76dc-4c80-8cfe-edc97816b850','1552881532948','xiaobei'),
	('ade46f92-59a9-40cc-9ae3-9af1fbfba800','1552751282787','xiaobei'),
	('b77e5a17-30bb-4434-87aa-ef541f39d0a5','1552751196670','xiaobei'),
	('ba9fafc2-7d95-48da-b5a0-befd514a5dc5','1552750874689','xiaobei'),
	('e006f91d-3251-42d1-8816-a94e18588830','1552891708996','xiaobei'),
	('efec3b2b-e8cf-494e-9dc9-5ae58e2afafc','1552881466219','xiaobei');

/*!40000 ALTER TABLE `sessions` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table users
# ------------------------------------------------------------

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `login_name` varchar(255) NOT NULL DEFAULT '' COMMENT '用户名称',
  `pwd` varchar(255) DEFAULT NULL COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;

INSERT INTO `users` (`id`, `login_name`, `pwd`)
VALUES
	(11,'xiaolin','123'),
	(12,'xiaobei','123'),
	(13,'xiaobei','123'),
	(14,'xiaobei','123'),
	(15,'xiaobei','123'),
	(16,'xiaobei','123'),
	(17,'xiaobei','123'),
	(18,'xiaobei','123'),
	(19,'xiaobei','123'),
	(20,'xiaobei','123'),
	(21,'xiaobei','123'),
	(22,'','123'),
	(23,'xiaobei','123'),
	(24,'xiaobei','123');

/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table video_del_rec
# ------------------------------------------------------------

DROP TABLE IF EXISTS `video_del_rec`;

CREATE TABLE `video_del_rec` (
  `video_id` varchar(64) NOT NULL DEFAULT '',
  PRIMARY KEY (`video_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table video_info
# ------------------------------------------------------------

DROP TABLE IF EXISTS `video_info`;

CREATE TABLE `video_info` (
  `id` varchar(64) NOT NULL DEFAULT '',
  `author_id` bigint(11) unsigned DEFAULT NULL,
  `name` text,
  `display_ctime` text,
  `create_time` date DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
