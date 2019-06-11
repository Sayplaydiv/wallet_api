/*
 Navicat MySQL Data Transfer

 Source Server         : new_audit
 Source Server Version : 50721
 Source Host           : 52.52.45.114
 Source Database       : new_audit

 Target Server Version : 50721
 File Encoding         : utf-8

 Date: 05/16/2019 17:23:56 PM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `depwit_status`
-- ----------------------------
DROP TABLE IF EXISTS `depwit_status`;
CREATE TABLE `depwit_status` (
  `id` int(32) NOT NULL AUTO_INCREMENT,
  `asset` varchar(18) CHARACTER SET utf8mb4 NOT NULL,
  `address` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
  `user_id` int(18) DEFAULT NULL,
  `amount` decimal(64,18) NOT NULL,
  `tx_hash` varchar(128) NOT NULL,
  `block_height` int(18) NOT NULL,
  `tx_time` datetime DEFAULT NULL,
  `confirm` int(18) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=59 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `depwit_status`
-- ----------------------------
BEGIN;
INSERT INTO `depwit_status` VALUES ('28', 'NANO', 'xrb_1xhpbmqqh44b5t9rm44cuidk7cqf63eicifms1u73phdc86jnromruw67qqo', '10557173', '0.000000000005076944', 'E183BE9A3C15734016330FAD1FECF92A5ADE8EDE8EA604477D53D53AF5D85544', '3830165', '2019-05-16 08:36:50', '1'), ('30', 'ETP', 'MEdDu4bixh3e3SN3XKySo5q1bFy11kiHh3', '10557173', '0.010000000000000000', '', '0', '2019-05-16 08:36:50', '2'), ('58', 'ETP', 'MEdDu4bixh3e3SN3XKySo5q1bFy11kiHh3', '10557266', '0.010000000000000000', '157619f56881a2a159720f150e117ff6be0f59cf5833d844e65baf6143529ec0', '2234706', '2019-05-16 09:23:53', '1');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
