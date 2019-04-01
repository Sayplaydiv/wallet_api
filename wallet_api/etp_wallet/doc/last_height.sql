/*
 Navicat MySQL Data Transfer

 Source Server         : 蚂蚁服务器数据库
 Source Server Version : 50725
 Source Host           : 18.144.38.221
 Source Database       : eth_wallet

 Target Server Version : 50725
 File Encoding         : utf-8

 Date: 03/13/2019 18:45:18 PM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `last_height`
-- ----------------------------
DROP TABLE IF EXISTS `last_height`;
CREATE TABLE `last_height` (
  `id` int(11) NOT NULL,
  `asset` char(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `height` int(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- ----------------------------
--  Records of `last_height`
-- ----------------------------
BEGIN;
INSERT INTO `last_height` VALUES ('1', 'etp', '1185354');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
