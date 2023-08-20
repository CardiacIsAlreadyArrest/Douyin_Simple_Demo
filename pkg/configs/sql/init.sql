SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- 每张表结构还会加上 gorm.Model

-- ----------------------------
-- Table structure for users
-- ----------------------------
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


-- ----------------------------
-- Table structure for videos
-- ----------------------------
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


-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '评论 ID',
  `user_id` bigint NOT NULL COMMENT '评论发布用户ID',
  `video_id` bigint NOT NULL COMMENT '评论视频ID',
  `comment_text` varchar(255) NOT NULL COMMENT '评论内容',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '评论创建时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '评论删除时间',
  PRIMARY KEY (`id`),
  KEY `videoIdIdx` (`video_id`) USING BTREE COMMENT '评论 ID 索引'
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8 COMMENT='评论表';


-- ----------------------------
-- Table structure for likes
-- ----------------------------
DROP TABLE IF EXISTS `likes`;
CREATE TABLE `likes` (
 `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增主键',
 `user_id` bigint NOT NULL COMMENT '点赞用户id',
 `video_id` bigint NOT NULL COMMENT '被点赞的视频id',
 `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '点赞创建时间',
 `deleted_at` timestamp NULL DEFAULT NULL COMMENT '点赞删除时间',
 PRIMARY KEY (`id`),
 KEY `userIdtoVideoIdIdx` (`user_id`,`video_id`) USING BTREE,
 KEY `userIdIdx` (`user_id`) USING BTREE,
 KEY `videoIdx` (`video_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8 COMMENT='点赞表';
