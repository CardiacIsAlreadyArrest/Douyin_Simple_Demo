DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `user_name` varchar(255) NOT NULL COMMENT '用户名',
  `password` varchar(255) NOT NULL COMMENT '用户密码',
  `avatar` varchar(255) NOT NULL COMMENT '用户头像',
  `background_image` varchar(255) NOT NULL COMMENT '用户个人页顶部大图',
  `signature` varchar(255) NOT NULL COMMENT '个人简介',
  PRIMARY KEY (`id`),
--   KEY `user_name_password_idx` (`user_name`,`password`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8 COMMENT='用户表';