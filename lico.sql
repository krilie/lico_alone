/*
Navicat MySQL Data Transfer

Source Server         : ligo.ml
Source Server Version : 50726
Source Host           : ligo.ml:3306
Source Database       : lico

Target Server Type    : MYSQL
Target Server Version : 50726
File Encoding         : 65001

Date: 2019-05-24 08:59:05
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for tb_client_user_acc_token
-- ----------------------------
DROP TABLE IF EXISTS `tb_client_user_acc_token`;
CREATE TABLE `tb_client_user_acc_token` (
  `token` varchar(32) NOT NULL,
  `create_time` datetime NOT NULL,
  `expiration_time` datetime NOT NULL,
  `user_id` varchar(32) NOT NULL,
  `create_by` varchar(32) DEFAULT NULL,
  `is_valid` tinyint(1) DEFAULT NULL,
  `description` varchar(128) DEFAULT NULL,
  PRIMARY KEY (`token`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tb_client_user_acc_token
-- ----------------------------
INSERT INTO `tb_client_user_acc_token` VALUES ('1234', '2019-05-24 08:58:24', '2019-06-09 08:58:19', '2', '1', '1', 'client端flutter');

-- ----------------------------
-- Table structure for tb_permission
-- ----------------------------
DROP TABLE IF EXISTS `tb_permission`;
CREATE TABLE `tb_permission` (
  `id` varchar(32) NOT NULL,
  `create_time` datetime NOT NULL,
  `name` varchar(50) NOT NULL,
  `description` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uix_tb_permission_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tb_permission
-- ----------------------------

-- ----------------------------
-- Table structure for tb_role
-- ----------------------------
DROP TABLE IF EXISTS `tb_role`;
CREATE TABLE `tb_role` (
  `id` varchar(32) NOT NULL,
  `create_time` datetime NOT NULL,
  `name` varchar(50) NOT NULL,
  `description` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uix_tb_role_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tb_role
-- ----------------------------
INSERT INTO `tb_role` VALUES ('1', '2019-05-24 08:54:57', 'admin', '管理员');
INSERT INTO `tb_role` VALUES ('2', '2019-05-24 08:56:27', 'client', '客户端');

-- ----------------------------
-- Table structure for tb_role_permission
-- ----------------------------
DROP TABLE IF EXISTS `tb_role_permission`;
CREATE TABLE `tb_role_permission` (
  `role_id` varchar(32) NOT NULL,
  `permission_id` varchar(32) NOT NULL,
  PRIMARY KEY (`role_id`,`permission_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tb_role_permission
-- ----------------------------

-- ----------------------------
-- Table structure for tb_user
-- ----------------------------
DROP TABLE IF EXISTS `tb_user`;
CREATE TABLE `tb_user` (
  `id` varchar(32) NOT NULL,
  `create_time` datetime NOT NULL,
  `login_name` varchar(50) NOT NULL,
  `nick_name` varchar(50) DEFAULT NULL,
  `phone` varchar(20) DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL,
  `password` varchar(64) NOT NULL,
  `picture` varchar(500) DEFAULT NULL,
  `salt` varchar(8) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uix_tb_user_login_name` (`login_name`),
  UNIQUE KEY `uix_tb_user_phone` (`phone`),
  UNIQUE KEY `uix_tb_user_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tb_user
-- ----------------------------
INSERT INTO `tb_user` VALUES ('1', '2019-05-24 08:51:09', '1', '1', '18855566126', '1197829331@qq.com', '123', null, '123');

-- ----------------------------
-- Table structure for tb_user_role
-- ----------------------------
DROP TABLE IF EXISTS `tb_user_role`;
CREATE TABLE `tb_user_role` (
  `user_id` varchar(32) NOT NULL,
  `role_id` varchar(32) NOT NULL,
  PRIMARY KEY (`user_id`,`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tb_user_role
-- ----------------------------
INSERT INTO `tb_user_role` VALUES ('1', '1');
INSERT INTO `tb_user_role` VALUES ('2', '2');
