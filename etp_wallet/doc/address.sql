/*
 Navicat MySQL Data Transfer

 Source Server         : 蚂蚁服务器数据库
 Source Server Version : 50725
 Source Host           : 18.144.38.221
 Source Database       : eth_wallet

 Target Server Version : 50725
 File Encoding         : utf-8

 Date: 03/13/2019 13:06:38 PM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `test`
-- ----------------------------
DROP TABLE IF EXISTS `test`;
CREATE TABLE `test` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `address` char(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `asset` char(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `inuse` int(11) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=latin1;

SET FOREIGN_KEY_CHECKS = 1;
