SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `user_name` varchar(255) NOT NULL COMMENT '用户名',
  `password` varchar(255) NOT NULL COMMENT '用户密码',
  `avatar` varchar(255) NOT NULL COMMENT '用户头像',
  `background_image` varchar(255) NOT NULL COMMENT '用户个人页顶部大图',
  `signature` varchar(255) NOT NULL COMMENT '个人简介',
  PRIMARY KEY (`id`),
  KEY `user_name_password_idx` (`user_name`,`password`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8 COMMENT='用户表';


DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '视频ID',
  `author_id` bigint NOT NULL COMMENT '视频作者id',
  `play_url` varchar(255) NOT NULL COMMENT '播放url',
  `cover_url` varchar(255) NOT NULL COMMENT '封面url',
  `publish_time` timestamp NOT NULL COMMENT '发布时间戳',
  `title` varchar(255) DEFAULT NULL COMMENT '视频名称',
  PRIMARY KEY (`id`),
  KEY `time` (`publish_time`) USING BTREE,
  KEY `author` (`author_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=115 DEFAULT CHARSET=utf8 COMMENT='视频表';