/*
Navicat MySQL Data Transfer

Source Server         : phpStudy
Source Server Version : 50553
Source Host           : localhost:3306
Source Database       : gocms

Target Server Type    : MYSQL
Target Server Version : 50553
File Encoding         : 65001

Date: 2019-08-23 16:03:39
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for gms_backend_user
-- ----------------------------
DROP TABLE IF EXISTS `gms_backend_user`;
CREATE TABLE `gms_backend_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `real_name` varchar(255) NOT NULL DEFAULT '',
  `user_name` varchar(255) NOT NULL DEFAULT '',
  `user_pwd` varchar(255) NOT NULL DEFAULT '',
  `is_super` tinyint(1) NOT NULL DEFAULT '0',
  `status` int(11) NOT NULL DEFAULT '0',
  `mobile` varchar(16) NOT NULL DEFAULT '',
  `email` varchar(256) NOT NULL DEFAULT '',
  `avatar` varchar(256) NOT NULL DEFAULT '',
  `school_id` int(11) DEFAULT '0' COMMENT '校区ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Records of gms_backend_user
-- ----------------------------
INSERT INTO `gms_backend_user` VALUES ('1', 'cuua', 'admin', 'e10adc3949ba59abbe56e057f20f883e', '1', '1', '10000', '7590547807@qq.com', '/static/upload/200.png', '10');
INSERT INTO `gms_backend_user` VALUES ('3', '张三', 'zhang3', 'e10adc3949ba59abbe56e057f20f883e', '0', '1', '', '', '', '11');
INSERT INTO `gms_backend_user` VALUES ('5', '李四', 'lisi', 'e10adc3949ba59abbe56e057f20f883e', '0', '1', '', '', '', '10');
INSERT INTO `gms_backend_user` VALUES ('6', '崔新', '开鲁路', 'c8837b23ff8aaa8a2dde915473ce0991', '0', '1', '13333333333', 'cx2018@qq.com', '', '12');
INSERT INTO `gms_backend_user` VALUES ('7', 'huni', 'huni', 'e10adc3949ba59abbe56e057f20f883e', '0', '1', '', '', '', '10');

-- ----------------------------
-- Table structure for gms_backend_user_rms_roles
-- ----------------------------
DROP TABLE IF EXISTS `gms_backend_user_rms_roles`;
CREATE TABLE `gms_backend_user_rms_roles` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `rms_backend_user_id` int(11) NOT NULL,
  `rms_role_id` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Records of gms_backend_user_rms_roles
-- ----------------------------

-- ----------------------------
-- Table structure for gms_banner
-- ----------------------------
DROP TABLE IF EXISTS `gms_banner`;
CREATE TABLE `gms_banner` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL COMMENT '名称',
  `url` varchar(100) DEFAULT NULL COMMENT '跳转地址',
  `image_url` varchar(150) NOT NULL COMMENT '图片地址',
  `sort` tinyint(4) DEFAULT '0',
  `state` tinyint(4) DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of gms_banner
-- ----------------------------
INSERT INTO `gms_banner` VALUES ('2', '暑期来了', '', '/static/upload/6fd67927-b00c-4084-9875-65a025c55d12_0.png', '3', '0');
INSERT INTO `gms_banner` VALUES ('3', '墨微教育', '', '/static/upload/c870e48e-f466-4560-853c-7149b0974a2f_0.png', '2', '0');
INSERT INTO `gms_banner` VALUES ('4', '团建', '', '/static/upload/344f992c-c735-4b4c-831f-506b504ec640_0.png', '1', '0');

-- ----------------------------
-- Table structure for gms_class_record
-- ----------------------------
DROP TABLE IF EXISTS `gms_class_record`;
CREATE TABLE `gms_class_record` (
  `id` int(16) unsigned NOT NULL AUTO_INCREMENT,
  `teacher_id` smallint(10) NOT NULL DEFAULT '0',
  `student_id` smallint(10) NOT NULL DEFAULT '0',
  `subject_id` smallint(10) NOT NULL DEFAULT '0',
  `end_time` char(10) NOT NULL DEFAULT '',
  `length` decimal(8,2) NOT NULL DEFAULT '2.00',
  `status` tinyint(5) NOT NULL DEFAULT '0' COMMENT '-1取消0待确认1确认',
  `date` date DEFAULT NULL COMMENT '日期',
  `time` varchar(20) DEFAULT '' COMMENT '上课时间',
  `grade` tinyint(4) DEFAULT '0' COMMENT '年级',
  `price` int(10) DEFAULT '0' COMMENT '教师单价',
  `amount` decimal(8,2) DEFAULT '0.00' COMMENT '总金额',
  `type` tinyint(4) DEFAULT '0' COMMENT '上课类型0一对一1小班2托班',
  `contract_id` varchar(50) NOT NULL DEFAULT '' COMMENT '合同ID',
  `student_price` varchar(50) DEFAULT '' COMMENT '学生消耗的单价多单价',
  `student_amount` decimal(8,2) DEFAULT '0.00' COMMENT '学生扣费总额',
  `created_at` varchar(30) DEFAULT NULL,
  `msg_state` int(2) DEFAULT '0' COMMENT '0未发送1已发送',
  `student_surplus` decimal(8,2) DEFAULT '0.00' COMMENT '学生余额',
  `student_quantity_surplus` float(8,2) DEFAULT '0.00' COMMENT '剩余时长',
  `school_name` varchar(30) NOT NULL COMMENT '校区名',
  `school_id` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `tid_sid_time` (`teacher_id`,`student_id`,`end_time`) USING BTREE,
  KEY `student_id` (`student_id`) USING BTREE,
  KEY `teacher_id` (`teacher_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=247 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Records of gms_class_record
-- ----------------------------
INSERT INTO `gms_class_record` VALUES ('1', '4', '6', '8', '', '2.00', '0', '2019-07-29', '6-8', '0', '0', '0.00', '0', '', '', '0.00', '2019-07-28 08:51:00', '0', '0.00', '0.00', '', '10');
INSERT INTO `gms_class_record` VALUES ('4', '6', '13', '8', '', '2.00', '0', '2019-08-03', '8-10', '0', '0', '0.00', '0', '', '', '0.00', '2019-07-28 08:51:00', '0', '0.00', '0.00', '', '10');
INSERT INTO `gms_class_record` VALUES ('5', '6', '2', '8', '', '2.00', '0', '2019-08-04', '8-10', '0', '0', '0.00', '0', '', '', '0.00', '2019-07-28 08:51:00', '0', '0.00', '0.00', '', '10');
INSERT INTO `gms_class_record` VALUES ('6', '6', '15', '8', '', '2.00', '0', '2019-08-03', '10-12', '0', '0', '0.00', '0', '', '', '0.00', '2019-07-28 08:51:00', '0', '0.00', '0.00', '', '10');
INSERT INTO `gms_class_record` VALUES ('7', '6', '16', '8', '', '2.00', '0', '2019-08-01', '13-15', '0', '0', '0.00', '0', '', '', '0.00', '2019-07-28 08:51:00', '0', '0.00', '0.00', '', '10');
INSERT INTO `gms_class_record` VALUES ('8', '6', '12', '8', '', '2.00', '0', '2019-08-02', '13-15', '0', '0', '0.00', '0', '', '', '0.00', '2019-07-28 08:51:00', '0', '0.00', '0.00', '', '10');
INSERT INTO `gms_class_record` VALUES ('45', '12', '4', '19', '', '2.00', '0', '2019-08-02', '17-19', '0', '0', '0.00', '2', '', '', '0.00', '2019-07-28 08:51:00', '0', '0.00', '0.00', '', '10');
INSERT INTO `gms_class_record` VALUES ('49', '13', '6', '19', '', '2.00', '0', '2019-07-31', '8-10', '0', '0', '0.00', '2', '', '', '0.00', '2019-07-28 08:51:00', '0', '0.00', '0.00', '', '10');
INSERT INTO `gms_class_record` VALUES ('52', '13', '7', '19', '', '2.00', '0', '2019-07-31', '10-12', '0', '0', '0.00', '2', '', '', '0.00', '2019-07-28 08:51:00', '0', '0.00', '0.00', '', '10');
INSERT INTO `gms_class_record` VALUES ('81', '8', '6', '9', '', '2.00', '0', '2019-07-29', '6-8', '0', '0', '0.00', '0', '', '', '0.00', '2019-07-28 09:09:35', '0', '0.00', '0.00', '', '10');
INSERT INTO `gms_class_record` VALUES ('82', '2', '2', '7', '', '1.00', '0', '2019-07-04', '8-10', '0', '0', '0.00', '0', '', '', '0.00', '2019-07-28 09:50:14', '0', '0.00', '0.00', '', '0');
INSERT INTO `gms_class_record` VALUES ('83', '1', '2', '7', '', '1.00', '0', '2019-07-12', '6-8', '0', '0', '0.00', '0', '', '', '0.00', '2019-07-28 10:25:05', '0', '0.00', '0.00', '', '0');
INSERT INTO `gms_class_record` VALUES ('84', '3', '2', '9', '', '0.00', '0', '2019-08-02', '10-12', '0', '0', '0.00', '0', '', '', '0.00', '2019-08-04 07:57:55', '0', '0.00', '0.00', '', '0');
INSERT INTO `gms_class_record` VALUES ('85', '4', '6', '8', '', '2.00', '0', '2019-08-05', '6-8', '0', '0', '0.00', '0', '', '', '0.00', '2019-08-04 10:30:02', '0', '0.00', '0.00', '', '10');
INSERT INTO `gms_class_record` VALUES ('86', '8', '6', '9', '', '2.00', '0', '2019-08-05', '6-8', '0', '0', '0.00', '0', '', '', '0.00', '2019-08-04 10:30:02', '0', '0.00', '0.00', '', '10');
INSERT INTO `gms_class_record` VALUES ('89', '6', '13', '8', '', '2.00', '0', '2019-08-10', '8-10', '0', '0', '0.00', '0', '', '', '0.00', '2019-08-04 10:30:02', '0', '0.00', '0.00', '', '10');
INSERT INTO `gms_class_record` VALUES ('90', '6', '2', '8', '', '2.00', '0', '2019-08-11', '8-10', '0', '0', '0.00', '0', '', '', '0.00', '2019-08-04 10:30:02', '0', '0.00', '0.00', '', '10');
INSERT INTO `gms_class_record` VALUES ('91', '6', '15', '8', '', '2.00', '0', '2019-08-10', '10-12', '0', '0', '0.00', '0', '', '', '0.00', '2019-08-04 10:30:02', '0', '0.00', '0.00', '', '10');
INSERT INTO `gms_class_record` VALUES ('92', '6', '16', '8', '', '2.00', '0', '2019-08-08', '13-15', '0', '0', '0.00', '0', '', '', '0.00', '2019-08-04 10:30:02', '0', '0.00', '0.00', '', '10');
INSERT INTO `gms_class_record` VALUES ('93', '6', '12', '8', '', '2.00', '0', '2019-08-09', '13-15', '0', '0', '0.00', '0', '', '', '0.00', '2019-08-04 10:30:02', '0', '0.00', '0.00', '', '10');
INSERT INTO `gms_class_record` VALUES ('130', '12', '4', '19', '', '2.00', '0', '2019-08-09', '17-19', '0', '0', '0.00', '2', '', '', '0.00', '2019-08-04 10:30:02', '0', '0.00', '0.00', '', '10');
INSERT INTO `gms_class_record` VALUES ('134', '13', '6', '19', '', '2.00', '0', '2019-08-07', '8-10', '0', '0', '0.00', '2', '', '', '0.00', '2019-08-04 10:30:02', '0', '0.00', '0.00', '', '10');
INSERT INTO `gms_class_record` VALUES ('137', '13', '7', '19', '', '2.00', '0', '2019-08-07', '10-12', '0', '0', '0.00', '2', '', '', '0.00', '2019-08-04 10:30:02', '0', '0.00', '0.00', '', '10');
INSERT INTO `gms_class_record` VALUES ('166', '4', '6', '8', '', '2.00', '0', '2019-08-12', '6-8', '0', '0', '0.00', '0', '', '', '0.00', '2019-08-11 10:30:01', '0', '0.00', '0.00', '', '10');
INSERT INTO `gms_class_record` VALUES ('167', '8', '6', '9', '', '2.00', '0', '2019-08-12', '6-8', '0', '0', '0.00', '0', '', '', '0.00', '2019-08-11 10:30:01', '0', '0.00', '0.00', '', '10');
INSERT INTO `gms_class_record` VALUES ('170', '6', '13', '8', '', '2.00', '0', '2019-08-17', '8-10', '0', '0', '0.00', '0', '', '', '0.00', '2019-08-11 10:30:01', '0', '0.00', '0.00', '', '10');
INSERT INTO `gms_class_record` VALUES ('171', '6', '2', '8', '', '2.00', '0', '2019-08-18', '8-10', '0', '0', '0.00', '0', '', '', '0.00', '2019-08-11 10:30:01', '0', '0.00', '0.00', '', '10');
INSERT INTO `gms_class_record` VALUES ('172', '6', '15', '8', '', '2.00', '0', '2019-08-17', '10-12', '0', '0', '0.00', '0', '', '', '0.00', '2019-08-11 10:30:01', '0', '0.00', '0.00', '', '10');
INSERT INTO `gms_class_record` VALUES ('173', '6', '16', '8', '', '2.00', '0', '2019-08-15', '13-15', '0', '0', '0.00', '0', '', '', '0.00', '2019-08-11 10:30:01', '0', '0.00', '0.00', '', '10');
INSERT INTO `gms_class_record` VALUES ('174', '6', '12', '8', '', '2.00', '0', '2019-08-16', '13-15', '0', '0', '0.00', '0', '', '', '0.00', '2019-08-11 10:30:01', '0', '0.00', '0.00', '', '10');
INSERT INTO `gms_class_record` VALUES ('211', '12', '4', '19', '', '2.00', '0', '2019-08-16', '17-19', '0', '0', '0.00', '2', '', '', '0.00', '2019-08-11 10:30:01', '0', '0.00', '0.00', '', '10');
INSERT INTO `gms_class_record` VALUES ('215', '13', '6', '19', '', '2.00', '0', '2019-08-14', '8-10', '0', '0', '0.00', '2', '', '', '0.00', '2019-08-11 10:30:01', '0', '0.00', '0.00', '', '10');
INSERT INTO `gms_class_record` VALUES ('218', '13', '7', '19', '', '2.00', '0', '2019-08-14', '10-12', '0', '0', '0.00', '2', '', '', '0.00', '2019-08-11 10:30:01', '0', '0.00', '0.00', '', '10');

-- ----------------------------
-- Table structure for gms_contract
-- ----------------------------
DROP TABLE IF EXISTS `gms_contract`;
CREATE TABLE `gms_contract` (
  `id` int(32) unsigned NOT NULL AUTO_INCREMENT,
  `student_id` smallint(10) NOT NULL DEFAULT '0' COMMENT '学生编号',
  `type` tinyint(4) NOT NULL COMMENT '类型0一对一1小班2托班',
  `price` int(11) NOT NULL COMMENT '单价',
  `quantity` decimal(8,2) NOT NULL DEFAULT '0.00' COMMENT '时长',
  `surplus_quantity` decimal(8,2) unsigned DEFAULT '0.00' COMMENT '剩余时长',
  `amount` decimal(8,2) DEFAULT '0.00' COMMENT '总金额',
  `surplus` decimal(8,2) DEFAULT '0.00',
  `payment` varchar(20) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime DEFAULT NULL,
  `status` tinyint(5) NOT NULL DEFAULT '0',
  `start_at` date DEFAULT NULL,
  `end_at` date DEFAULT NULL,
  `update_at` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=840 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Records of gms_contract
-- ----------------------------
INSERT INTO `gms_contract` VALUES ('1', '2', '0', '160', '71.00', '0.00', '11360.00', '0.00', 'pos', '2018-05-20 22:54:32', null, '0', '2018-05-20', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('2', '6', '2', '250', '1.00', '1.00', '250.00', '250.00', 'pos', '2018-05-20 22:57:43', '2018-05-20 23:00:15', '-1', '1970-01-01', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('3', '7', '2', '250', '1.00', '1.00', '250.00', '250.00', 'pos', '2018-05-20 22:58:07', '2018-05-20 23:00:24', '-1', '1970-01-01', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('4', '1', '1', '100', '3.00', '3.00', '300.00', '300.00', 'pos', '2018-05-20 22:58:37', '2018-05-20 23:04:43', '-1', '2018-05-20', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('5', '3', '1', '150', '3.00', '3.00', '450.00', '450.00', 'pos', '2018-05-20 22:59:08', '2018-05-20 23:00:35', '-1', '2018-05-20', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('6', '8', '2', '375', '1.00', '0.00', '375.00', '0.00', 'pos', '2018-05-20 22:59:23', null, '0', '2018-05-20', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('7', '6', '2', '250', '1.00', '0.00', '250.00', '0.00', 'pos', '2018-05-20 23:05:48', null, '0', '2018-05-20', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('8', '7', '2', '250', '1.00', '0.00', '250.00', '0.00', 'pos', '2018-05-20 23:06:19', null, '0', '2018-05-20', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('9', '9', '2', '375', '2.00', '0.00', '750.00', '0.00', 'pos', '2018-05-20 23:07:29', null, '0', '2018-05-20', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('10', '10', '2', '375', '2.00', '1.00', '750.00', '375.00', 'pos', '2018-05-20 23:07:52', '2019-02-18 12:28:29', '2', '2018-05-20', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('11', '11', '0', '200', '10.00', '0.00', '2000.00', '0.00', 'pos', '2018-05-20 23:09:00', null, '0', '2018-05-20', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('12', '12', '0', '200', '5.00', '1.00', '1000.00', '200.00', 'pos', '2018-05-20 23:09:49', '2018-06-10 18:14:50', '2', '2018-05-20', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('13', '12', '1', '140', '116.00', '0.00', '16240.00', '0.00', 'pos', '2018-05-20 23:11:13', null, '0', '2018-05-20', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('14', '13', '0', '200', '78.00', '56.00', '15600.00', '11200.00', 'pos', '2018-05-20 23:12:20', null, '0', '2018-05-20', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('15', '15', '0', '200', '24.00', '0.00', '4800.00', '0.00', 'pos', '2018-05-20 23:15:16', null, '0', '2018-05-20', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('16', '16', '0', '280', '10.00', '0.00', '2800.00', '0.00', 'pos', '2018-05-20 23:16:00', null, '0', '2018-05-20', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('51', '7', '2', '250', '5.00', '0.00', '1250.00', '0.00', 'pos', '2018-05-21 18:02:45', null, '0', '2018-05-21', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('52', '6', '2', '250', '4.00', '0.00', '1000.00', '0.00', 'pos', '2018-05-21 18:03:00', null, '0', '2018-05-21', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('62', '13', '0', '200', '50.00', '0.00', '10000.00', '0.00', 'cash', '2018-05-26 19:42:27', null, '0', '2018-05-26', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('64', '1', '0', '100', '5.00', '3.00', '500.00', '300.00', 'pos', '2018-05-27 15:52:34', '2018-08-12 14:14:20', '-1', '2018-05-27', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('66', '6', '2', '250', '1.00', '0.00', '250.00', '0.00', 'cash', '2018-05-30 01:38:23', null, '0', '2018-05-30', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('70', '9', '0', '180', '16.00', '0.00', '2880.00', '0.00', 'cash', '2018-05-30 01:45:26', null, '0', '2018-05-30', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('72', '10', '2', '1', '375.00', '375.00', '375.00', '375.00', 'pos', '2018-05-31 19:47:15', '2018-06-03 18:37:52', '-1', '2018-05-31', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('73', '9', '1', '375', '2.00', '2.00', '750.00', '750.00', 'pos', '2018-05-31 20:00:42', '2018-06-03 18:35:23', '-1', '2018-05-31', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('79', '9', '2', '375', '2.00', '0.00', '750.00', '0.00', 'cash', '2018-06-03 18:36:33', null, '0', '2018-06-03', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('80', '10', '2', '375', '2.00', '2.00', '750.00', '750.00', 'cash', '2018-06-03 18:38:54', '2018-06-03 18:39:53', '-1', '2018-06-03', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('81', '10', '2', '375', '1.00', '1.00', '375.00', '375.00', 'cash', '2018-06-03 18:40:11', '2019-02-18 12:28:35', '2', '2018-06-03', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('82', '9', '1', '105', '48.00', '10.00', '5040.00', '1050.00', 'cash', '2018-06-07 14:06:39', '2018-10-20 11:41:35', '2', '2018-06-07', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('92', '12', '0', '185', '119.00', '117.00', '22015.00', '21645.00', 'cash', '2018-06-10 15:29:24', '2018-06-10 18:11:01', '2', '2018-06-10', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('93', '12', '0', '160', '180.00', '57.00', '28800.00', '9120.00', 'cash', '2018-06-10 18:21:40', null, '0', '2018-06-10', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('100', '15', '0', '200', '80.00', '0.00', '16000.00', '0.00', 'cash', '2018-06-15 20:21:50', null, '0', '2018-06-15', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('116', '11', '0', '200', '32.00', '1.00', '6400.00', '200.00', 'cash', '2018-06-28 16:20:51', '2018-08-30 15:39:34', '2', '2018-06-28', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('120', '16', '0', '280', '8.00', '0.00', '2240.00', '0.00', 'cash', '2018-07-01 16:55:53', null, '0', '2018-07-01', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('142', '16', '0', '280', '16.00', '0.00', '4480.00', '0.00', 'cash', '2018-07-31 12:38:31', null, '0', '2018-07-31', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('143', '2', '0', '160', '30.00', '0.00', '4800.00', '0.00', 'cash', '2018-07-31 12:43:33', null, '0', '2018-07-31', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('160', '2', '0', '160', '130.00', '0.00', '20800.00', '0.00', 'pos', '2018-08-04 16:11:55', null, '0', '2018-08-04', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('169', '7', '2', '150', '3.00', '0.00', '450.00', '0.00', 'cash', '2018-08-13 16:45:08', null, '0', '2018-08-13', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('170', '6', '2', '150', '3.00', '0.00', '450.00', '0.00', 'cash', '2018-08-13 16:45:24', null, '0', '2018-08-13', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('171', '4', '2', '400', '2.00', '0.00', '800.00', '0.00', 'cash', '2018-08-13 18:42:25', null, '0', '2018-08-13', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('181', '15', '0', '200', '84.00', '5.00', '16800.00', '1000.00', 'pos', '2018-08-19 13:44:37', null, '0', '2018-08-19', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('198', '6', '2', '300', '10.00', '0.00', '3000.00', '0.00', 'cash', '2018-08-28 13:55:19', null, '0', '2018-08-28', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('199', '7', '2', '300', '10.00', '0.00', '3000.00', '0.00', 'cash', '2018-08-28 13:55:32', null, '0', '2018-08-28', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('203', '4', '1', '70', '40.00', '40.00', '2800.00', '2800.00', 'pos', '2018-08-29 17:04:05', '2018-11-12 12:17:12', '2', '2018-08-29', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('204', '4', '2', '300', '4.00', '0.00', '1200.00', '0.00', 'pos', '2018-08-29 17:04:16', null, '0', '2018-08-29', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('210', '11', '1', '80', '40.00', '6.00', '3200.00', '480.00', 'cash', '2018-08-30 15:40:39', null, '0', '2018-08-30', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('215', '9', '2', '300', '4.00', '1.00', '1200.00', '300.00', 'cash', '2018-08-31 12:42:40', '2018-10-20 11:41:51', '2', '2018-08-31', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('217', '4', '1', '70', '40.00', '0.00', '2800.00', '0.00', 'pos', '2018-08-31 17:04:27', null, '0', '2018-08-31', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('220', '16', '0', '280', '18.00', '0.00', '5040.00', '0.00', 'cash', '2018-09-01 15:03:49', null, '0', '2018-09-01', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('221', '16', '1', '100', '80.00', '4.00', '8000.00', '400.00', 'cash', '2018-09-01 15:03:59', null, '0', '2018-09-01', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('254', '11', '0', '200', '40.00', '4.00', '8000.00', '800.00', 'cash', '2018-09-09 17:13:16', null, '0', '2018-09-09', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('268', '16', '0', '280', '24.00', '0.00', '6720.00', '0.00', 'cash', '2018-09-23 15:08:45', null, '0', '2018-09-23', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('290', '4', '2', '300', '4.00', '0.00', '1200.00', '0.00', 'pos', '2018-10-08 18:55:41', null, '0', '2018-10-08', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('297', '12', '2', '300', '1.00', '0.00', '300.00', '0.00', 'cash', '2018-10-09 20:33:22', null, '0', '2018-10-09', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('299', '13', '0', '200', '100.00', '0.00', '20000.00', '0.00', 'cash', '2018-10-10 19:13:34', null, '0', '2018-10-10', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('317', '16', '0', '280', '24.00', '0.00', '6720.00', '0.00', 'cash', '2018-10-23 19:14:44', null, '0', '2018-10-23', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('319', '15', '0', '200', '28.00', '0.00', '5600.00', '0.00', 'cash', '2018-10-24 18:31:29', null, '0', '2018-10-24', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('326', '4', '1', '300', '4.00', '0.00', '1200.00', '0.00', 'pos', '2018-10-29 19:18:13', '2018-11-03 19:46:33', '2', '2018-10-29', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('330', '12', '2', '220', '1.00', '0.00', '220.00', '0.00', 'cash', '2018-11-02 14:21:02', null, '0', '2018-11-02', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('336', '4', '2', '300', '4.00', '0.00', '1200.00', '0.00', 'pos', '2018-11-03 19:46:50', null, '0', '2018-11-03', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('354', '4', '2', '70', '36.00', '36.00', '2520.00', '2520.00', 'cash', '2018-11-12 12:18:00', '2018-11-12 12:18:33', '-1', '2018-11-12', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('355', '4', '1', '70', '36.00', '0.00', '2520.00', '0.00', 'cash', '2018-11-12 12:18:54', null, '0', '2018-11-12', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('364', '7', '2', '300', '10.00', '0.00', '3000.00', '0.00', 'cash', '2018-11-16 16:03:27', null, '0', '2018-11-16', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('365', '6', '2', '300', '10.00', '0.00', '3000.00', '0.00', 'cash', '2018-11-16 16:03:39', null, '0', '2018-11-16', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('375', '16', '0', '300', '24.00', '0.00', '7200.00', '0.00', 'cash', '2018-11-23 19:06:12', null, '0', '2018-11-23', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('386', '4', '2', '300', '7.00', '0.00', '2100.00', '0.00', 'pos', '2018-11-29 19:17:39', null, '0', '2018-11-29', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('422', '2', '0', '160', '100.00', '25.50', '16000.00', '4080.00', 'cash', '2018-12-24 19:55:06', null, '0', '2018-12-24', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('428', '16', '0', '300', '24.00', '1.00', '7200.00', '300.00', 'cash', '2018-12-27 17:47:07', null, '0', '2018-12-27', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('432', '15', '0', '200', '40.00', '1.00', '8000.00', '200.00', 'pos', '2018-12-28 20:24:00', null, '0', '2018-12-28', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('487', '4', '1', '70', '50.00', '0.00', '3500.00', '0.00', 'pos', '2019-01-21 20:08:56', null, '0', '2019-01-21', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('488', '4', '2', '400', '2.00', '0.00', '800.00', '0.00', 'pos', '2019-01-21 20:09:05', null, '0', '2019-01-21', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('489', '4', '0', '14', '210.00', '200.00', '2940.00', '2800.00', 'pos', '2019-01-21 20:09:24', '2019-01-29 15:50:22', '2', '2019-01-21', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('492', '16', '0', '300', '24.00', '0.00', '7200.00', '0.00', 'cash', '2019-01-22 11:13:49', null, '0', '2019-01-22', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('493', '11', '0', '200', '20.00', '0.00', '4000.00', '0.00', 'cash', '2019-01-22 13:56:39', null, '0', '2019-01-22', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('494', '11', '1', '80', '40.00', '2.00', '3200.00', '160.00', 'cash', '2019-01-22 13:56:50', null, '0', '2019-01-22', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('521', '4', '0', '210', '4.00', '0.00', '840.00', '0.00', 'pos', '2019-01-29 15:50:42', null, '0', '2019-01-29', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('524', '16', '1', '120', '80.00', '10.00', '9600.00', '1200.00', 'cash', '2019-02-10 17:44:58', null, '0', '2019-02-10', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('533', '2', '0', '160', '44.50', '1.00', '7120.00', '160.00', 'cash', '2019-02-11 19:22:00', null, '0', '2019-02-11', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('534', '2', '0', '160', '20.00', '0.00', '3200.00', '0.00', 'cash', '2019-02-12 00:03:23', null, '0', '2019-02-12', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('546', '15', '0', '200', '114.00', '26.00', '22800.00', '5200.00', 'pos', '2019-02-15 17:47:29', null, '0', '2019-02-15', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('547', '13', '1', '90', '40.00', '20.00', '3600.00', '1800.00', 'cash', '2019-02-16 18:02:34', null, '0', '2019-02-16', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('561', '4', '2', '375', '10.00', '1.00', '3750.00', '375.00', 'pos', '2019-02-19 20:42:27', null, '0', '2019-02-19', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('590', '12', '1', '130', '180.00', '150.00', '23400.00', '19500.00', 'pos', '2019-02-23 11:15:44', null, '0', '2019-02-23', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('605', '6', '2', '300', '10.00', '0.00', '3000.00', '0.00', 'cash', '2019-02-25 14:28:17', null, '0', '2019-02-25', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('606', '7', '2', '300', '10.00', '0.00', '3000.00', '0.00', 'cash', '2019-02-25 14:30:05', null, '0', '2019-02-25', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('608', '16', '0', '300', '24.00', '0.00', '7200.00', '0.00', 'cash', '2019-02-28 17:25:20', null, '0', '2019-02-28', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('645', '4', '2', '0', '1.00', '0.00', '0.00', '0.00', 'cash', '2019-03-18 18:15:05', null, '0', '2019-03-18', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('670', '16', '0', '300', '24.00', '0.00', '7200.00', '0.00', 'cash', '2019-03-29 14:29:28', null, '0', '2019-03-29', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('671', '11', '0', '230', '24.00', '8.00', '5520.00', '1840.00', 'cash', '2019-03-29 15:12:22', null, '0', '2019-03-29', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('672', '11', '1', '80', '48.00', '17.50', '3840.00', '1400.00', 'cash', '2019-03-29 15:12:41', null, '0', '2019-03-29', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('692', '12', '2', '375', '1.00', '0.00', '375.00', '0.00', 'cash', '2019-04-08 19:46:25', null, '0', '2019-04-08', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('700', '12', '2', '375', '1.00', '0.00', '375.00', '0.00', 'cash', '2019-04-15 18:20:29', null, '0', '2019-04-15', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('713', '12', '2', '375', '1.00', '0.00', '375.00', '0.00', 'cash', '2019-04-22 14:57:55', null, '0', '2019-04-22', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('718', '4', '2', '0', '1.00', '0.00', '0.00', '0.00', 'cash', '2019-04-23 19:27:49', null, '0', '2019-04-23', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('743', '4', '1', '70', '32.00', '16.00', '2240.00', '1120.00', 'pos', '2019-04-28 18:36:29', null, '0', '2019-04-28', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('744', '4', '2', '375', '4.00', '2.00', '1500.00', '750.00', 'pos', '2019-04-28 18:36:37', null, '0', '2019-04-28', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('756', '16', '0', '300', '24.00', '4.00', '7200.00', '1200.00', 'cash', '2019-05-02 15:20:33', null, '0', '2019-05-02', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('759', '12', '2', '375', '1.00', '0.00', '375.00', '0.00', 'cash', '2019-05-06 14:55:12', null, '0', '2019-05-06', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('775', '4', '0', '210', '20.00', '16.00', '4200.00', '3360.00', 'pos', '2019-05-08 20:42:45', null, '0', '2019-05-08', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('776', '6', '2', '300', '6.00', '4.00', '1800.00', '1200.00', 'cash', '2019-05-08 20:54:15', null, '0', '2019-05-08', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('777', '7', '2', '300', '6.00', '4.00', '1800.00', '1200.00', 'cash', '2019-05-08 20:54:32', null, '0', '2019-05-08', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('789', '12', '2', '375', '1.00', '0.00', '375.00', '0.00', 'cash', '2019-05-13 14:44:20', null, '0', '2019-05-13', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('793', '4', '2', '0', '1.00', '0.00', '0.00', '0.00', 'cash', '2019-05-13 19:48:34', null, '0', '2019-05-13', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('815', '12', '2', '375', '1.00', '0.00', '375.00', '0.00', 'cash', '2019-05-20 15:40:27', null, '0', '2019-05-20', '1970-01-01', '');
INSERT INTO `gms_contract` VALUES ('838', '12', '2', '375', '1.00', '1.00', '375.00', '375.00', 'cash', '2019-05-28 17:25:06', null, '0', '2019-05-28', '1970-01-01', '');

-- ----------------------------
-- Table structure for gms_contract_price
-- ----------------------------
DROP TABLE IF EXISTS `gms_contract_price`;
CREATE TABLE `gms_contract_price` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `type` tinyint(4) NOT NULL COMMENT '合同类型0对一1小班2托班',
  `grade` mediumint(10) NOT NULL COMMENT '年级',
  `school_id` mediumint(10) NOT NULL COMMENT '校区ID',
  `lower` int(11) NOT NULL COMMENT '下限',
  `upper` int(11) NOT NULL COMMENT '上限',
  `grade120` int(11) DEFAULT '0',
  `grade110` int(11) DEFAULT '0',
  `grade100` int(11) DEFAULT '0',
  `grade90` int(11) DEFAULT '0',
  `grade80` int(11) DEFAULT '0',
  `grade70` int(11) DEFAULT '0',
  `grade60` int(11) DEFAULT '0',
  `grade50` int(11) DEFAULT '0',
  `grade40` int(11) DEFAULT '0',
  `grade30` int(11) DEFAULT '0',
  `grade20` int(11) DEFAULT '0',
  `grade10` int(11) DEFAULT '0',
  `grade3` int(11) DEFAULT '0',
  `grade2` int(11) DEFAULT '0',
  `grade1` int(11) DEFAULT '0',
  `status120` tinyint(4) DEFAULT '0' COMMENT '0不锁定1锁定',
  `status110` tinyint(4) DEFAULT '0' COMMENT '0不锁定1锁定',
  `status100` tinyint(4) DEFAULT '0' COMMENT '0不锁定1锁定',
  `status90` tinyint(4) DEFAULT '0' COMMENT '0不锁定1锁定',
  `status80` tinyint(4) DEFAULT '0' COMMENT '0不锁定1锁定',
  `status70` tinyint(4) DEFAULT '0' COMMENT '0不锁定1锁定',
  `status60` tinyint(4) DEFAULT '0' COMMENT '0不锁定1锁定',
  `status50` tinyint(4) DEFAULT '0' COMMENT '0不锁定1锁定',
  `status40` tinyint(4) DEFAULT '0' COMMENT '0不锁定1锁定',
  `status30` tinyint(4) DEFAULT '0' COMMENT '0不锁定1锁定',
  `status20` tinyint(4) DEFAULT '0' COMMENT '0不锁定1锁定',
  `status10` tinyint(4) DEFAULT '0' COMMENT '0不锁定1锁定',
  `status3` tinyint(4) DEFAULT '0' COMMENT '0不锁定1锁定',
  `status2` tinyint(4) DEFAULT '0' COMMENT '0不锁定1锁定',
  `status1` tinyint(4) DEFAULT '0' COMMENT '0不锁定1锁定',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of gms_contract_price
-- ----------------------------
INSERT INTO `gms_contract_price` VALUES ('1', '0', '0', '10', '0', '59', '400', '400', '380', '350', '320', '270', '270', '240', '210', '210', '210', '210', '210', '210', '210', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0');
INSERT INTO `gms_contract_price` VALUES ('2', '0', '0', '10', '60', '119', '380', '380', '360', '330', '305', '255', '255', '230', '200', '200', '200', '200', '200', '200', '200', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0');
INSERT INTO `gms_contract_price` VALUES ('3', '0', '0', '10', '120', '999', '360', '360', '340', '300', '275', '275', '230', '210', '180', '180', '180', '180', '180', '180', '180', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0');
INSERT INTO `gms_contract_price` VALUES ('4', '1', '0', '10', '0', '59', '999', '999', '999', '120', '110', '90', '90', '80', '70', '70', '70', '70', '70', '70', '70', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0');
INSERT INTO `gms_contract_price` VALUES ('5', '1', '0', '10', '60', '119', '999', '999', '999', '120', '110', '90', '90', '80', '70', '70', '70', '70', '70', '70', '70', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0');
INSERT INTO `gms_contract_price` VALUES ('6', '1', '0', '10', '120', '999', '999', '999', '999', '120', '110', '90', '90', '80', '70', '70', '70', '70', '70', '70', '70', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0');
INSERT INTO `gms_contract_price` VALUES ('13', '0', '0', '11', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0');
INSERT INTO `gms_contract_price` VALUES ('14', '0', '0', '11', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0');
INSERT INTO `gms_contract_price` VALUES ('15', '0', '0', '11', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0');
INSERT INTO `gms_contract_price` VALUES ('16', '1', '0', '11', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0');
INSERT INTO `gms_contract_price` VALUES ('17', '1', '0', '11', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0');
INSERT INTO `gms_contract_price` VALUES ('18', '1', '0', '11', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0');
INSERT INTO `gms_contract_price` VALUES ('19', '0', '0', '12', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0');
INSERT INTO `gms_contract_price` VALUES ('20', '0', '0', '12', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0');
INSERT INTO `gms_contract_price` VALUES ('21', '0', '0', '12', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0');
INSERT INTO `gms_contract_price` VALUES ('22', '1', '0', '12', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0');
INSERT INTO `gms_contract_price` VALUES ('23', '1', '0', '12', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0');
INSERT INTO `gms_contract_price` VALUES ('24', '1', '0', '12', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0');

-- ----------------------------
-- Table structure for gms_expend
-- ----------------------------
DROP TABLE IF EXISTS `gms_expend`;
CREATE TABLE `gms_expend` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) NOT NULL,
  `amount` int(11) NOT NULL,
  `month` date DEFAULT NULL COMMENT '月份',
  `school_id` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=28 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of gms_expend
-- ----------------------------
INSERT INTO `gms_expend` VALUES ('3', '测试哦', '1000', '2019-05-01', '0');
INSERT INTO `gms_expend` VALUES ('4', 'en', '2000', '2019-04-01', '0');
INSERT INTO `gms_expend` VALUES ('5', 'en', '2000', '2019-05-01', '0');
INSERT INTO `gms_expend` VALUES ('6', 'en', '2000', '2019-05-01', '0');
INSERT INTO `gms_expend` VALUES ('7', 'en', '2000', '2019-05-01', '0');
INSERT INTO `gms_expend` VALUES ('8', 'en', '2000', '2019-05-01', '0');
INSERT INTO `gms_expend` VALUES ('9', 'en', '2000', '2019-05-01', '0');
INSERT INTO `gms_expend` VALUES ('10', '测试哦', '1000', '2019-06-01', '0');
INSERT INTO `gms_expend` VALUES ('11', 'en', '2000', '2019-06-01', '0');
INSERT INTO `gms_expend` VALUES ('12', 'en', '2000', '2019-06-01', '0');
INSERT INTO `gms_expend` VALUES ('13', 'en', '2000', '2019-06-01', '0');
INSERT INTO `gms_expend` VALUES ('14', 'en', '2000', '2019-06-01', '0');
INSERT INTO `gms_expend` VALUES ('15', 'en', '2000', '2019-06-01', '0');
INSERT INTO `gms_expend` VALUES ('16', '测试哦', '1000', '2019-07-01', '0');
INSERT INTO `gms_expend` VALUES ('17', 'en', '2000', '2019-07-01', '0');
INSERT INTO `gms_expend` VALUES ('18', 'en', '2000', '2019-07-01', '0');
INSERT INTO `gms_expend` VALUES ('19', 'en', '2000', '2019-07-01', '0');
INSERT INTO `gms_expend` VALUES ('20', 'en', '2000', '2019-07-01', '0');
INSERT INTO `gms_expend` VALUES ('21', 'en', '2000', '2019-07-01', '0');
INSERT INTO `gms_expend` VALUES ('22', '测试哦', '1000', '2019-08-01', '0');
INSERT INTO `gms_expend` VALUES ('23', 'en', '2000', '2019-08-01', '0');
INSERT INTO `gms_expend` VALUES ('24', 'en', '2000', '2019-08-01', '0');
INSERT INTO `gms_expend` VALUES ('25', 'en', '2000', '2019-08-01', '0');
INSERT INTO `gms_expend` VALUES ('26', 'en', '2000', '2019-08-01', '0');
INSERT INTO `gms_expend` VALUES ('27', 'en', '2000', '2019-08-01', '0');

-- ----------------------------
-- Table structure for gms_history
-- ----------------------------
DROP TABLE IF EXISTS `gms_history`;
CREATE TABLE `gms_history` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '',
  `content` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `sort` int(11) NOT NULL DEFAULT '0',
  `date` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of gms_history
-- ----------------------------

-- ----------------------------
-- Table structure for gms_news
-- ----------------------------
DROP TABLE IF EXISTS `gms_news`;
CREATE TABLE `gms_news` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '',
  `content` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `sort` int(11) NOT NULL DEFAULT '0',
  `image_url` varchar(100) NOT NULL DEFAULT '' COMMENT '封面',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of gms_news
-- ----------------------------

-- ----------------------------
-- Table structure for gms_project
-- ----------------------------
DROP TABLE IF EXISTS `gms_project`;
CREATE TABLE `gms_project` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(45) NOT NULL DEFAULT '',
  `city` varchar(45) NOT NULL COMMENT '城市',
  `school_id` tinyint(4) NOT NULL DEFAULT '0',
  `content` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `sort` int(11) DEFAULT '0',
  `school_name` varchar(45) DEFAULT '',
  `image_url` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of gms_project
-- ----------------------------

-- ----------------------------
-- Table structure for gms_recruit
-- ----------------------------
DROP TABLE IF EXISTS `gms_recruit`;
CREATE TABLE `gms_recruit` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '',
  `content` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `sort` int(11) NOT NULL DEFAULT '0',
  `image_url` varchar(100) NOT NULL DEFAULT '' COMMENT '封面',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of gms_recruit
-- ----------------------------

-- ----------------------------
-- Table structure for gms_resource
-- ----------------------------
DROP TABLE IF EXISTS `gms_resource`;
CREATE TABLE `gms_resource` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `rtype` int(11) NOT NULL DEFAULT '0',
  `name` varchar(64) NOT NULL DEFAULT '',
  `parent_id` int(11) DEFAULT NULL,
  `seq` int(11) NOT NULL DEFAULT '0',
  `icon` varchar(32) NOT NULL DEFAULT '',
  `url_for` varchar(256) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=98 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Records of gms_resource
-- ----------------------------
INSERT INTO `gms_resource` VALUES ('7', '1', '权限管理', '8', '100', 'fa fa-balance-scale', '');
INSERT INTO `gms_resource` VALUES ('8', '0', '系统菜单', null, '200', '', '');
INSERT INTO `gms_resource` VALUES ('9', '1', '资源管理', '7', '100', '', 'ResourceController.Index');
INSERT INTO `gms_resource` VALUES ('12', '1', '角色管理', '7', '100', '', 'RoleController.Index');
INSERT INTO `gms_resource` VALUES ('13', '1', '用户管理', '7', '100', '', 'BackendUserController.Index');
INSERT INTO `gms_resource` VALUES ('14', '1', '系统管理', '8', '90', 'fa fa-gears', '');
INSERT INTO `gms_resource` VALUES ('21', '0', '业务菜单', null, '170', '', '');
INSERT INTO `gms_resource` VALUES ('22', '1', '学生中心', '21', '100', 'fa fa-user', '');
INSERT INTO `gms_resource` VALUES ('23', '1', '日志管理(空)', '14', '100', '', '');
INSERT INTO `gms_resource` VALUES ('25', '2', '编辑', '9', '100', 'fa fa-pencil', 'ResourceController.Edit');
INSERT INTO `gms_resource` VALUES ('26', '2', '编辑', '13', '100', 'fa fa-pencil', 'BackendUserController.Edit');
INSERT INTO `gms_resource` VALUES ('27', '2', '删除', '9', '100', 'fa fa-trash', 'ResourceController.Delete');
INSERT INTO `gms_resource` VALUES ('29', '2', '删除', '13', '100', 'fa fa-trash', 'BackendUserController.Delete');
INSERT INTO `gms_resource` VALUES ('30', '2', '编辑', '12', '100', 'fa fa-pencil', 'RoleController.Edit');
INSERT INTO `gms_resource` VALUES ('31', '2', '删除', '12', '100', 'fa fa-trash', 'RoleController.Delete');
INSERT INTO `gms_resource` VALUES ('32', '2', '分配资源', '12', '100', 'fa fa-th', 'RoleController.Allocate');
INSERT INTO `gms_resource` VALUES ('35', '1', ' 首页', null, '100', 'fa fa-dashboard', 'HomeController.Index');
INSERT INTO `gms_resource` VALUES ('36', '1', '学生信息', '22', '100', '', 'StudentController.Index');
INSERT INTO `gms_resource` VALUES ('37', '1', '班组管理', '22', '100', '', 'StudentController.Team');
INSERT INTO `gms_resource` VALUES ('38', '1', '学生地区', '22', '100', '', 'StudentAreaController.Index');
INSERT INTO `gms_resource` VALUES ('39', '1', '学生学校', '22', '100', '', 'StudentSchoolController.Index');
INSERT INTO `gms_resource` VALUES ('40', '1', '教师中心', '21', '100', 'fa fa-user-secret', '');
INSERT INTO `gms_resource` VALUES ('41', '1', '教师信息', '40', '100', '', 'TeacherController.Index');
INSERT INTO `gms_resource` VALUES ('42', '1', '校区信息', '40', '100', '', 'SchoolController.Index');
INSERT INTO `gms_resource` VALUES ('43', '1', '教学科目', '40', '100', '', 'SubjectController.Index');
INSERT INTO `gms_resource` VALUES ('44', '1', '教师课表', '40', '100', '', 'ScheduleController.Index');
INSERT INTO `gms_resource` VALUES ('45', '1', '合同中心', '21', '100', 'fa fa-file-text-o', '');
INSERT INTO `gms_resource` VALUES ('46', '1', '合同列表', '45', '100', '', 'ContractController.Index');
INSERT INTO `gms_resource` VALUES ('47', '1', '课程中心', '21', '100', 'fa fa-server', '');
INSERT INTO `gms_resource` VALUES ('48', '1', '上课记录', '47', '100', '', 'ClassRecordController.Index');
INSERT INTO `gms_resource` VALUES ('49', '1', '默认单价', '45', '100', '', 'ContractPriceController.Index');
INSERT INTO `gms_resource` VALUES ('50', '3', '班组数据', '37', '100', '', 'StudentController.TeamDataGrid');
INSERT INTO `gms_resource` VALUES ('51', '3', '学生分页数据', '36', '100', '', 'StudentController.DataGrid');
INSERT INTO `gms_resource` VALUES ('52', '3', '校区列表', '40', '100', '', 'SchoolController.DataList');
INSERT INTO `gms_resource` VALUES ('53', '3', '学生学校列表', '39', '100', '', 'StudentSchoolController.DataList');
INSERT INTO `gms_resource` VALUES ('54', '3', '学生地区列表', '38', '100', '', 'StudentAreaController.DataList');
INSERT INTO `gms_resource` VALUES ('55', '3', '学生地区分页', '38', '100', '', 'StudentAreaController.DataGrid');
INSERT INTO `gms_resource` VALUES ('56', '3', '学生学校分页', '39', '100', '', 'StudentSchoolController.DataGrid');
INSERT INTO `gms_resource` VALUES ('57', '2', '添加', '36', '100', '', 'StudentController.Edit');
INSERT INTO `gms_resource` VALUES ('58', '3', '上课记录分页', '48', '100', '', 'ClassRecordController.DataGrid');
INSERT INTO `gms_resource` VALUES ('59', '3', '教师信息分页', '41', '100', '', 'TeacherController.DataGrid');
INSERT INTO `gms_resource` VALUES ('60', '2', '添加', '37', '100', '', 'StudentController.TeamEdit');
INSERT INTO `gms_resource` VALUES ('61', '3', '学生/班组列表', '35', '100', '', 'StudentController.DataList');
INSERT INTO `gms_resource` VALUES ('62', '3', '教师列表', '40', '100', '', 'TeacherController.DataList');
INSERT INTO `gms_resource` VALUES ('63', '2', '排课表编辑', '44', '100', '', 'ScheduleController.Edit');
INSERT INTO `gms_resource` VALUES ('64', '1', '排课明细', '47', '100', '', 'ClassRecordController.Handle');
INSERT INTO `gms_resource` VALUES ('65', '1', '排课管理', '47', '100', '', 'ClassRecordController.Single');
INSERT INTO `gms_resource` VALUES ('66', '1', '数据中心', '21', '100', 'fa fa-list-alt', '');
INSERT INTO `gms_resource` VALUES ('67', '1', '教师数据', '66', '100', '', '');
INSERT INTO `gms_resource` VALUES ('68', '3', '教师工资明细', '67', '100', '', 'TeacherController.SalaryDetailGrid');
INSERT INTO `gms_resource` VALUES ('69', '1', '教师工资', '67', '100', '', 'TeacherController.Salary');
INSERT INTO `gms_resource` VALUES ('70', '1', '教师课耗', '67', '100', '', 'TeacherController.Useup');
INSERT INTO `gms_resource` VALUES ('71', '1', '教师预排课', '67', '100', '', 'TeacherController.Userate');
INSERT INTO `gms_resource` VALUES ('72', '1', '学生数据', '66', '100', '', '');
INSERT INTO `gms_resource` VALUES ('73', '1', '学生统计', '72', '100', '', 'StudentController.Total');
INSERT INTO `gms_resource` VALUES ('74', '1', '新生统计', '72', '100', '', 'StudentController.New');
INSERT INTO `gms_resource` VALUES ('75', '1', '合同数据', '66', '100', '', '');
INSERT INTO `gms_resource` VALUES ('76', '1', '合同统计', '75', '100', '', 'ContractController.Total');
INSERT INTO `gms_resource` VALUES ('77', '1', '合同周统计', '75', '100', '', 'ContractController.New');
INSERT INTO `gms_resource` VALUES ('78', '1', '合同余额', '75', '100', '', 'ContractController.Balance');
INSERT INTO `gms_resource` VALUES ('79', '1', '课程数据', '66', '100', '', '');
INSERT INTO `gms_resource` VALUES ('80', '1', '类型课耗统计', '79', '100', '', 'ClassRecordController.Class');
INSERT INTO `gms_resource` VALUES ('81', '1', '年级课耗统计', '79', '100', '', 'ClassRecordController.ClassGrade');
INSERT INTO `gms_resource` VALUES ('82', '1', '课耗周统计', '79', '100', '', 'ClassRecordController.ClassWeek');
INSERT INTO `gms_resource` VALUES ('83', '1', '损益数据', '66', '100', '', '');
INSERT INTO `gms_resource` VALUES ('84', '1', '支出', '83', '100', '', 'ExpendController.Index');
INSERT INTO `gms_resource` VALUES ('85', '1', '利润', '83', '100', '', 'ExpendController.Balance');
INSERT INTO `gms_resource` VALUES ('86', '1', '现金结余', '83', '100', '', 'ExpendController.Profit');
INSERT INTO `gms_resource` VALUES ('87', '2', '添加', '41', '100', '', 'TeacherController.Edit');
INSERT INTO `gms_resource` VALUES ('88', '2', '教师科目', '41', '100', '', 'TeacherSubjectController.Edit');
INSERT INTO `gms_resource` VALUES ('89', '3', '科目列表', '35', '100', '', 'SubjectController.DataList');
INSERT INTO `gms_resource` VALUES ('90', '1', '官网管理', '8', '180', 'fa fa-home', '');
INSERT INTO `gms_resource` VALUES ('91', '1', '轮播图', '90', '100', '', 'BannerController.Index');
INSERT INTO `gms_resource` VALUES ('92', '1', '新闻管理', '90', '100', '', 'NewsController.Index');
INSERT INTO `gms_resource` VALUES ('93', '1', '培训项目', '90', '100', '', 'ProjectController.Index');
INSERT INTO `gms_resource` VALUES ('94', '1', '关于我们', '90', '100', '', 'UsController.Index');
INSERT INTO `gms_resource` VALUES ('95', '1', '优秀员工', '90', '100', '', 'StaffController.Index');
INSERT INTO `gms_resource` VALUES ('96', '1', '招聘管理', '90', '100', '', 'RecruitController.Index');
INSERT INTO `gms_resource` VALUES ('97', '1', '时间线', '90', '100', '', 'HistoryController.Index');

-- ----------------------------
-- Table structure for gms_role
-- ----------------------------
DROP TABLE IF EXISTS `gms_role`;
CREATE TABLE `gms_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `seq` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Records of gms_role
-- ----------------------------
INSERT INTO `gms_role` VALUES ('22', '超级管理员', '20');
INSERT INTO `gms_role` VALUES ('24', '教师组', '10');
INSERT INTO `gms_role` VALUES ('25', '教务管理员', '5');
INSERT INTO `gms_role` VALUES ('26', '店长', '0');

-- ----------------------------
-- Table structure for gms_role_backenduser_rel
-- ----------------------------
DROP TABLE IF EXISTS `gms_role_backenduser_rel`;
CREATE TABLE `gms_role_backenduser_rel` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_id` int(11) NOT NULL,
  `backend_user_id` int(11) NOT NULL,
  `created` datetime NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=80 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Records of gms_role_backenduser_rel
-- ----------------------------
INSERT INTO `gms_role_backenduser_rel` VALUES ('73', '25', '3', '2019-06-09 04:09:01');
INSERT INTO `gms_role_backenduser_rel` VALUES ('75', '24', '5', '2019-06-09 04:43:49');
INSERT INTO `gms_role_backenduser_rel` VALUES ('79', '22', '1', '2019-08-02 06:06:36');

-- ----------------------------
-- Table structure for gms_role_resource_rel
-- ----------------------------
DROP TABLE IF EXISTS `gms_role_resource_rel`;
CREATE TABLE `gms_role_resource_rel` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_id` int(11) NOT NULL,
  `resource_id` int(11) NOT NULL,
  `created` datetime NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1034 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Records of gms_role_resource_rel
-- ----------------------------
INSERT INTO `gms_role_resource_rel` VALUES ('431', '22', '35', '2017-12-19 06:40:07');
INSERT INTO `gms_role_resource_rel` VALUES ('432', '22', '21', '2017-12-19 06:40:07');
INSERT INTO `gms_role_resource_rel` VALUES ('433', '22', '22', '2017-12-19 06:40:07');
INSERT INTO `gms_role_resource_rel` VALUES ('434', '22', '8', '2017-12-19 06:40:07');
INSERT INTO `gms_role_resource_rel` VALUES ('435', '22', '14', '2017-12-19 06:40:07');
INSERT INTO `gms_role_resource_rel` VALUES ('436', '22', '23', '2017-12-19 06:40:07');
INSERT INTO `gms_role_resource_rel` VALUES ('437', '22', '7', '2017-12-19 06:40:07');
INSERT INTO `gms_role_resource_rel` VALUES ('438', '22', '9', '2017-12-19 06:40:07');
INSERT INTO `gms_role_resource_rel` VALUES ('439', '22', '25', '2017-12-19 06:40:07');
INSERT INTO `gms_role_resource_rel` VALUES ('440', '22', '27', '2017-12-19 06:40:07');
INSERT INTO `gms_role_resource_rel` VALUES ('441', '22', '12', '2017-12-19 06:40:07');
INSERT INTO `gms_role_resource_rel` VALUES ('442', '22', '30', '2017-12-19 06:40:07');
INSERT INTO `gms_role_resource_rel` VALUES ('443', '22', '31', '2017-12-19 06:40:07');
INSERT INTO `gms_role_resource_rel` VALUES ('444', '22', '32', '2017-12-19 06:40:07');
INSERT INTO `gms_role_resource_rel` VALUES ('445', '22', '13', '2017-12-19 06:40:07');
INSERT INTO `gms_role_resource_rel` VALUES ('446', '22', '26', '2017-12-19 06:40:07');
INSERT INTO `gms_role_resource_rel` VALUES ('447', '22', '29', '2017-12-19 06:40:07');
INSERT INTO `gms_role_resource_rel` VALUES ('871', '25', '35', '2019-04-21 03:39:19');
INSERT INTO `gms_role_resource_rel` VALUES ('872', '25', '21', '2019-04-21 03:39:19');
INSERT INTO `gms_role_resource_rel` VALUES ('873', '25', '22', '2019-04-21 03:39:19');
INSERT INTO `gms_role_resource_rel` VALUES ('874', '25', '36', '2019-04-21 03:39:19');
INSERT INTO `gms_role_resource_rel` VALUES ('875', '25', '51', '2019-04-21 03:39:19');
INSERT INTO `gms_role_resource_rel` VALUES ('876', '25', '57', '2019-04-21 03:39:19');
INSERT INTO `gms_role_resource_rel` VALUES ('877', '25', '61', '2019-04-21 03:39:19');
INSERT INTO `gms_role_resource_rel` VALUES ('878', '25', '37', '2019-04-21 03:39:19');
INSERT INTO `gms_role_resource_rel` VALUES ('879', '25', '50', '2019-04-21 03:39:19');
INSERT INTO `gms_role_resource_rel` VALUES ('880', '25', '60', '2019-04-21 03:39:19');
INSERT INTO `gms_role_resource_rel` VALUES ('881', '25', '38', '2019-04-21 03:39:19');
INSERT INTO `gms_role_resource_rel` VALUES ('882', '25', '54', '2019-04-21 03:39:19');
INSERT INTO `gms_role_resource_rel` VALUES ('883', '25', '55', '2019-04-21 03:39:19');
INSERT INTO `gms_role_resource_rel` VALUES ('884', '25', '39', '2019-04-21 03:39:19');
INSERT INTO `gms_role_resource_rel` VALUES ('885', '25', '53', '2019-04-21 03:39:19');
INSERT INTO `gms_role_resource_rel` VALUES ('886', '25', '56', '2019-04-21 03:39:19');
INSERT INTO `gms_role_resource_rel` VALUES ('887', '25', '40', '2019-04-21 03:39:19');
INSERT INTO `gms_role_resource_rel` VALUES ('888', '25', '44', '2019-04-21 03:39:19');
INSERT INTO `gms_role_resource_rel` VALUES ('889', '25', '63', '2019-04-21 03:39:19');
INSERT INTO `gms_role_resource_rel` VALUES ('890', '25', '52', '2019-04-21 03:39:19');
INSERT INTO `gms_role_resource_rel` VALUES ('891', '25', '62', '2019-04-21 03:39:19');
INSERT INTO `gms_role_resource_rel` VALUES ('892', '25', '47', '2019-04-21 03:39:19');
INSERT INTO `gms_role_resource_rel` VALUES ('893', '25', '48', '2019-04-21 03:39:19');
INSERT INTO `gms_role_resource_rel` VALUES ('894', '25', '58', '2019-04-21 03:39:19');
INSERT INTO `gms_role_resource_rel` VALUES ('895', '26', '35', '2019-06-09 04:06:30');
INSERT INTO `gms_role_resource_rel` VALUES ('896', '26', '21', '2019-06-09 04:06:30');
INSERT INTO `gms_role_resource_rel` VALUES ('1016', '24', '35', '2019-06-09 23:44:58');
INSERT INTO `gms_role_resource_rel` VALUES ('1017', '24', '61', '2019-06-09 23:44:58');
INSERT INTO `gms_role_resource_rel` VALUES ('1018', '24', '89', '2019-06-09 23:44:58');
INSERT INTO `gms_role_resource_rel` VALUES ('1019', '24', '21', '2019-06-09 23:44:58');
INSERT INTO `gms_role_resource_rel` VALUES ('1020', '24', '40', '2019-06-09 23:44:58');
INSERT INTO `gms_role_resource_rel` VALUES ('1021', '24', '41', '2019-06-09 23:44:58');
INSERT INTO `gms_role_resource_rel` VALUES ('1022', '24', '59', '2019-06-09 23:44:58');
INSERT INTO `gms_role_resource_rel` VALUES ('1023', '24', '87', '2019-06-09 23:44:58');
INSERT INTO `gms_role_resource_rel` VALUES ('1024', '24', '88', '2019-06-09 23:44:58');
INSERT INTO `gms_role_resource_rel` VALUES ('1025', '24', '44', '2019-06-09 23:44:58');
INSERT INTO `gms_role_resource_rel` VALUES ('1026', '24', '63', '2019-06-09 23:44:58');
INSERT INTO `gms_role_resource_rel` VALUES ('1027', '24', '52', '2019-06-09 23:44:58');
INSERT INTO `gms_role_resource_rel` VALUES ('1028', '24', '62', '2019-06-09 23:44:58');
INSERT INTO `gms_role_resource_rel` VALUES ('1029', '24', '47', '2019-06-09 23:44:58');
INSERT INTO `gms_role_resource_rel` VALUES ('1030', '24', '48', '2019-06-09 23:44:58');
INSERT INTO `gms_role_resource_rel` VALUES ('1031', '24', '58', '2019-06-09 23:44:58');
INSERT INTO `gms_role_resource_rel` VALUES ('1032', '24', '64', '2019-06-09 23:44:58');
INSERT INTO `gms_role_resource_rel` VALUES ('1033', '24', '23', '2019-06-09 23:44:58');

-- ----------------------------
-- Table structure for gms_schedule
-- ----------------------------
DROP TABLE IF EXISTS `gms_schedule`;
CREATE TABLE `gms_schedule` (
  `id` int(16) unsigned NOT NULL AUTO_INCREMENT,
  `subject_id` smallint(10) NOT NULL,
  `teacher_id` smallint(10) NOT NULL,
  `school_id` smallint(10) NOT NULL COMMENT '校区ID',
  `class10` int(11) DEFAULT '0',
  `class20` int(11) DEFAULT '0',
  `class30` int(11) DEFAULT '0',
  `class40` int(11) DEFAULT '0',
  `class50` int(11) DEFAULT '0',
  `class60` int(11) DEFAULT '0',
  `class70` int(11) DEFAULT '0',
  `class80` int(11) DEFAULT '0',
  `class90` int(11) DEFAULT '0',
  `class100` int(11) DEFAULT '0',
  `class110` int(11) DEFAULT '0',
  `class120` int(11) DEFAULT '0',
  `class130` int(11) DEFAULT '0',
  `class140` int(11) DEFAULT '0',
  `class150` int(11) DEFAULT '0',
  `class160` int(11) DEFAULT '0',
  `class170` int(11) DEFAULT '0',
  `class180` int(11) DEFAULT '0',
  `class190` int(11) DEFAULT '0',
  `class200` int(11) DEFAULT '0',
  `class210` int(11) DEFAULT '0',
  `class220` int(11) DEFAULT '0',
  `class230` int(11) DEFAULT '0',
  `class240` int(11) DEFAULT '0',
  `class250` int(11) DEFAULT '0',
  `class260` int(11) DEFAULT '0',
  `class270` int(11) DEFAULT '0',
  `class280` int(11) DEFAULT '0',
  `class290` int(11) DEFAULT '0',
  `class300` int(11) DEFAULT '0',
  `class310` int(11) DEFAULT '0',
  `class320` int(11) DEFAULT '0',
  `class330` int(11) DEFAULT '0',
  `class340` int(11) DEFAULT '0',
  `class350` int(11) DEFAULT '0',
  `class360` int(11) DEFAULT '0',
  `class370` int(11) DEFAULT '0',
  `class380` int(11) DEFAULT '0',
  `class390` int(11) DEFAULT '0',
  `class400` int(11) DEFAULT '0',
  `class410` int(11) DEFAULT '0',
  `class420` int(11) DEFAULT '0',
  `class430` int(11) DEFAULT '0',
  `class440` int(11) DEFAULT '0',
  `class450` int(11) DEFAULT '0',
  `class460` int(11) DEFAULT '0',
  `class470` int(11) DEFAULT '0',
  `class480` int(11) DEFAULT '0',
  `class490` int(11) DEFAULT '0',
  `class500` int(11) DEFAULT '0',
  `class510` int(11) DEFAULT '0',
  `class520` int(11) DEFAULT '0',
  `class530` int(11) DEFAULT '0',
  `class540` int(11) DEFAULT '0',
  `class550` int(11) DEFAULT '0',
  `class560` int(11) DEFAULT '0',
  `class10_length` decimal(8,2) DEFAULT '0.00',
  `class20_length` decimal(8,2) DEFAULT '0.00',
  `class30_length` decimal(8,2) DEFAULT '0.00',
  `class40_length` decimal(8,2) DEFAULT '0.00',
  `class50_length` decimal(8,2) DEFAULT '0.00',
  `class60_length` decimal(8,2) DEFAULT '0.00',
  `class70_length` decimal(8,2) DEFAULT '0.00',
  `class80_length` decimal(8,2) DEFAULT '0.00',
  `class90_length` decimal(8,2) DEFAULT '0.00',
  `class100_length` decimal(8,2) DEFAULT '0.00',
  `class110_length` decimal(8,2) DEFAULT '0.00',
  `class120_length` decimal(8,2) DEFAULT '0.00',
  `class130_length` decimal(8,2) DEFAULT '0.00',
  `class140_length` decimal(8,2) DEFAULT '0.00',
  `class150_length` decimal(8,2) DEFAULT '0.00',
  `class160_length` decimal(8,2) DEFAULT '0.00',
  `class170_length` decimal(8,2) DEFAULT '0.00',
  `class180_length` decimal(8,2) DEFAULT '0.00',
  `class190_length` decimal(8,2) DEFAULT '0.00',
  `class200_length` decimal(8,2) DEFAULT '0.00',
  `class210_length` decimal(8,2) DEFAULT '0.00',
  `class220_length` decimal(8,2) DEFAULT '0.00',
  `class230_length` decimal(8,2) DEFAULT '0.00',
  `class240_length` decimal(8,2) DEFAULT '0.00',
  `class250_length` decimal(8,2) DEFAULT '0.00',
  `class260_length` decimal(8,2) DEFAULT '0.00',
  `class270_length` decimal(8,2) DEFAULT '0.00',
  `class280_length` decimal(8,2) DEFAULT '0.00',
  `class290_length` decimal(8,2) DEFAULT '0.00',
  `class300_length` decimal(8,2) DEFAULT '0.00',
  `class310_length` decimal(8,2) DEFAULT '0.00',
  `class320_length` decimal(8,2) DEFAULT '0.00',
  `class330_length` decimal(8,2) DEFAULT '0.00',
  `class340_length` decimal(8,2) DEFAULT '0.00',
  `class350_length` decimal(8,2) DEFAULT '0.00',
  `class360_length` decimal(8,2) DEFAULT '0.00',
  `class370_length` decimal(8,2) DEFAULT '0.00',
  `class380_length` decimal(8,2) DEFAULT '0.00',
  `class390_length` decimal(8,2) DEFAULT '0.00',
  `class400_length` decimal(8,2) DEFAULT '0.00',
  `class410_length` decimal(8,2) DEFAULT '0.00',
  `class420_length` decimal(8,2) DEFAULT '0.00',
  `class430_length` decimal(8,2) DEFAULT '0.00',
  `class440_length` decimal(8,2) DEFAULT '0.00',
  `class450_length` decimal(8,2) DEFAULT '0.00',
  `class460_length` decimal(8,2) DEFAULT '0.00',
  `class470_length` decimal(8,2) DEFAULT '0.00',
  `class480_length` decimal(8,2) DEFAULT '0.00',
  `class490_length` decimal(8,2) DEFAULT '0.00',
  `class500_length` decimal(8,2) DEFAULT '0.00',
  `class510_length` decimal(8,2) DEFAULT '0.00',
  `class520_length` decimal(8,2) DEFAULT '0.00',
  `class530_length` decimal(8,2) DEFAULT '0.00',
  `class540_length` decimal(8,2) DEFAULT '0.00',
  `class550_length` decimal(8,2) DEFAULT '0.00',
  `class560_length` decimal(8,2) DEFAULT '0.00',
  `created_at` datetime NOT NULL,
  `updated_at` datetime DEFAULT NULL,
  `status` tinyint(4) DEFAULT '0' COMMENT '状态0正常-1禁用',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `teacher_id-subject_id` (`teacher_id`,`subject_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Records of gms_schedule
-- ----------------------------
INSERT INTO `gms_schedule` VALUES ('2', '7', '5', '10', '4', '224', '6', '2', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '2', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '1.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '0.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '0.50', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '0.50', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '1.00', '0.00', '0.00', '0.00', '0000-00-00 00:00:00', '2019-04-13 09:01:21', '-1');
INSERT INTO `gms_schedule` VALUES ('3', '8', '4', '10', '6', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '2019-08-06 09:12:27', '2019-08-06 09:12:27', '0');
INSERT INTO `gms_schedule` VALUES ('4', '9', '8', '10', '6', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '1.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '2019-07-28 08:58:54', '2019-07-28 08:58:54', '0');
INSERT INTO `gms_schedule` VALUES ('5', '8', '6', '10', '0', '0', '0', '0', '0', '66', '68', '0', '0', '0', '0', '0', '13', '2', '0', '0', '0', '0', '0', '15', '0', '0', '0', '0', '16', '12', '38', '173', '0', '0', '0', '0', '22', '0', '0', '0', '0', '0', '0', '0', '0', '0', '40', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '2019-07-22 01:51:55', '2019-07-22 01:51:55', '0');
INSERT INTO `gms_schedule` VALUES ('11', '9', '15', '10', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '85', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '1.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0000-00-00 00:00:00', '2019-04-29 11:17:01', '0');
INSERT INTO `gms_schedule` VALUES ('13', '7', '16', '10', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '2018-05-13 23:37:32', '2018-09-05 21:02:00', '0');
INSERT INTO `gms_schedule` VALUES ('14', '8', '7', '10', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '60', '0', '0', '0', '0', '170', '25', '0', '0', '0', '0', '0', '196', '63', '0', '0', '0', '42', '0', '24', '0', '0', '0', '0', '0', '0', '164', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '2018-05-13 23:42:25', '2019-03-26 20:21:01', '0');
INSERT INTO `gms_schedule` VALUES ('15', '8', '1', '10', '0', '5', null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '2018-05-14 22:17:36', null, '0');
INSERT INTO `gms_schedule` VALUES ('16', '11', '1', '10', '4', '5', '4', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0000-00-00 00:00:00', '2019-04-12 10:51:53', '0');
INSERT INTO `gms_schedule` VALUES ('18', '19', '14', '10', '0', '0', '0', '24', '180', '0', '0', '0', '0', '0', '41', '55', '0', '0', '0', '0', '0', '105', '32', '0', '0', '0', '0', '0', '39', '40', '0', '0', '0', '0', '0', '25', '234', '0', '0', '0', '0', '0', '196', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '2018-05-17 19:58:35', '2019-04-10 16:18:49', '0');
INSERT INTO `gms_schedule` VALUES ('19', '19', '12', '10', '0', '0', '214', '186', '17', '0', '0', '0', '0', '70', '184', '26', '0', '0', '0', '0', '0', '174', '86', '0', '0', '0', '0', '122', '169', '181', '0', '0', '0', '0', '211', '120', '191', '0', '0', '0', '0', '0', '58', '4', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '2018-05-18 16:40:28', '2019-03-27 16:12:39', '0');
INSERT INTO `gms_schedule` VALUES ('20', '19', '13', '10', '0', '0', '224', '225', '125', '0', '0', '0', '0', '6', '198', '118', '0', '0', '0', '0', '7', '57', '177', '0', '0', '0', '0', '197', '207', '217', '0', '0', '0', '0', '101', '190', '223', '0', '0', '0', '0', '215', '28', '213', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '2018-05-20 23:44:18', '2019-03-27 16:11:40', '0');
INSERT INTO `gms_schedule` VALUES ('21', '9', '16', '10', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '2018-05-26 15:26:47', '2019-01-26 15:07:36', '0');
INSERT INTO `gms_schedule` VALUES ('22', '7', '17', '10', '0', '0', '0', '0', '0', '175', '105', '0', '0', '0', '0', '0', '156', '24', '0', '0', '0', '0', '165', '38', '110', '0', '0', '0', '0', '0', '51', '0', '0', '0', '0', '0', '0', '0', '189', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '2018-06-01 15:25:24', '2019-04-07 14:26:25', '0');
INSERT INTO `gms_schedule` VALUES ('23', '8', '18', '10', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '2018-06-07 14:50:55', '2018-10-16 15:51:25', '0');
INSERT INTO `gms_schedule` VALUES ('24', '8', '11', '10', '0', '0', '0', '0', '0', '110', '168', '0', '0', '0', '0', '0', '175', '41', '0', '0', '0', '0', '0', '0', '197', '0', '0', '0', '0', '0', '31', '28', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '2018-07-03 08:22:49', '2019-04-09 15:12:20', '0');
INSERT INTO `gms_schedule` VALUES ('25', '11', '7', '10', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '160', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '209', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '2018-07-11 15:19:24', '2019-03-26 20:33:23', '0');
INSERT INTO `gms_schedule` VALUES ('26', '10', '7', '10', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '166', '0', '0', '0', '0', '0', '0', '0', '166', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '2018-07-11 15:22:00', '2019-02-24 12:37:04', '0');
INSERT INTO `gms_schedule` VALUES ('27', '7', '20', '10', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '233', '0', '0', '0', '0', '0', '0', '22', '2', '0', '15', '16', '0', '0', '180', '155', '206', '0', '0', '0', '0', '13', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '2018-09-01 15:52:50', '2019-03-24 21:08:42', '0');
INSERT INTO `gms_schedule` VALUES ('28', '7', '21', '10', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '31', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '2018-09-30 20:35:01', '2018-10-07 17:55:47', '0');
INSERT INTO `gms_schedule` VALUES ('29', '9', '24', '10', '0', '0', '0', '0', '0', '157', '0', '0', '0', '0', '0', '0', '228', '0', '0', '0', '0', '0', '0', '31', '0', '0', '0', '0', '0', '222', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '28', '181', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '2019-01-20 13:34:31', '2019-04-06 17:55:06', '0');
INSERT INTO `gms_schedule` VALUES ('30', '8', '25', '11', '257', '255', '0', '0', '0', '0', '231', '0', '0', '0', '0', '0', '0', '36', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0.50', '1.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '2.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00', '0000-00-00 00:00:00', '2019-04-21 11:58:15', '0');

-- ----------------------------
-- Table structure for gms_school
-- ----------------------------
DROP TABLE IF EXISTS `gms_school`;
CREATE TABLE `gms_school` (
  `id` smallint(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(40) NOT NULL DEFAULT '' COMMENT '学校名',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of gms_school
-- ----------------------------
INSERT INTO `gms_school` VALUES ('10', '黄冈中学');
INSERT INTO `gms_school` VALUES ('11', '师范中学');
INSERT INTO `gms_school` VALUES ('12', '成都七中');

-- ----------------------------
-- Table structure for gms_staff
-- ----------------------------
DROP TABLE IF EXISTS `gms_staff`;
CREATE TABLE `gms_staff` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `image_url` varchar(150) NOT NULL DEFAULT '',
  `name` varchar(45) NOT NULL DEFAULT '',
  `content` text NOT NULL,
  `created_at` datetime NOT NULL,
  `sort` int(11) NOT NULL DEFAULT '0',
  `subject_name` varchar(30) NOT NULL DEFAULT '',
  `subject` smallint(8) NOT NULL,
  `school` varchar(255) NOT NULL DEFAULT '',
  `school_name` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of gms_staff
-- ----------------------------

-- ----------------------------
-- Table structure for gms_student
-- ----------------------------
DROP TABLE IF EXISTS `gms_student`;
CREATE TABLE `gms_student` (
  `id` smallint(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(10) NOT NULL DEFAULT '' COMMENT '姓名',
  `code` varchar(20) DEFAULT '' COMMENT '学生编号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0单个学生1小班',
  `sex` tinyint(4) DEFAULT '0' COMMENT '1-男2-女0-未知',
  `guarder` varchar(30) DEFAULT '' COMMENT '监护人',
  `relate` varchar(10) DEFAULT '' COMMENT '监护人关系',
  `grade` smallint(10) NOT NULL DEFAULT '0' COMMENT '年级',
  `contact1` varchar(15) DEFAULT '' COMMENT '联系方式',
  `contact2` varchar(15) DEFAULT '' COMMENT '联系方式',
  `school_id` smallint(10) DEFAULT '0' COMMENT '学校编号',
  `student_school_id` smallint(10) DEFAULT '0' COMMENT '学生学校ID',
  `adress` varchar(50) DEFAULT '' COMMENT '地址',
  `note` varchar(100) DEFAULT '' COMMENT '备注',
  `created_at` datetime NOT NULL,
  `updated_at` datetime DEFAULT NULL,
  `status` tinyint(5) NOT NULL DEFAULT '0' COMMENT '状态0未报名1正常2休眠-1删除',
  `group_id` text COMMENT '班组学生编号',
  `group_name` varchar(100) DEFAULT '' COMMENT '班组学生名',
  `balance1` decimal(8,2) unsigned DEFAULT '0.00' COMMENT '一对一合同剩余金额',
  `balance1_length` decimal(8,2) unsigned DEFAULT '0.00' COMMENT '一对一剩余时长',
  `balance2` decimal(8,2) unsigned DEFAULT '0.00' COMMENT '小班剩余金额',
  `balance2_length` decimal(8,2) unsigned DEFAULT '0.00' COMMENT '小班剩余时长',
  `balance3` decimal(8,2) unsigned DEFAULT '0.00' COMMENT '托班剩余金额',
  `balance3_length` decimal(8,2) unsigned DEFAULT '0.00' COMMENT '托班剩余时长',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `status` (`status`) USING BTREE,
  KEY `type_schoolid` (`type`,`school_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Records of gms_student
-- ----------------------------
INSERT INTO `gms_student` VALUES ('1', '李阳', 'st000000', '0', '1', '李刚', '爸爸', '20', '10086', '', '10', '87', '明月居', '', '2018-03-19 10:22:47', '2019-08-23 07:42:04', '1', '', '', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00');
INSERT INTO `gms_student` VALUES ('2', '夏天', 'st000001', '0', '1', '李娟', '妈妈', '70', '10086', null, '10', '87', '保利花城', '', '2018-03-19 10:33:01', '2019-08-23 03:15:31', '1', '', '', '3600.00', '22.50', '0.00', '0.00', '0.00', '0.00');
INSERT INTO `gms_student` VALUES ('3', '胡妮', 'st000002', '0', '1', '胡妮爸', '爸爸', '40', '10086', '', '10', '87', '康泰园', '', '2018-03-20 11:55:43', '2019-08-23 07:41:57', '1', '', '', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00');
INSERT INTO `gms_student` VALUES ('4', '韩信', 'st000003', '0', '1', '舞阳', '妈妈', '30', '10086', null, '10', '80', '珠江帝景', '', '2018-03-22 11:08:57', '2019-08-23 03:15:27', '1', '', '', '2940.00', '14.00', '840.00', '12.00', '750.00', '2.00');
INSERT INTO `gms_student` VALUES ('6', '金鑫', 'st000005', '0', '1', '金金', '爸爸', '50', '10086', null, '10', '85', '明月居', '', '2018-04-08 21:34:02', '2019-08-23 07:41:43', '1', '', '', '0.00', '0.00', '0.00', '0.00', '900.00', '3.00');
INSERT INTO `gms_student` VALUES ('7', '金成', 'st000006', '0', '1', '金金', '爸爸', '40', '10086', null, '10', '93', '亚运城', '', '2018-04-08 21:34:54', '2019-08-23 07:41:34', '1', '', '', '0.00', '0.00', '0.00', '0.00', '900.00', '3.00');
INSERT INTO `gms_student` VALUES ('8', '沈万', 'st000007', '0', '1', '翁凡', '妈妈', '50', '10086', null, '10', '85', '珠江帝景', '', '2018-04-10 12:55:36', '2019-08-23 07:41:23', '1', '', '', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00');
INSERT INTO `gms_student` VALUES ('9', '陈阳', 'st000008', '0', '1', '王玉', '妈妈', '50', '10086', null, '10', '85', '珠江帝景', '', '2018-04-10 12:56:06', '2018-08-07 16:40:53', '1', '', '', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00');
INSERT INTO `gms_student` VALUES ('10', '吴福', 'st000009', '0', '1', '董玄', '妈妈', '30', '10086', null, '10', '80', '保利花城', '', '2018-04-10 12:56:43', '2019-08-23 03:15:03', '1', '', '', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00');
INSERT INTO `gms_student` VALUES ('11', '李永', 'st000010', '0', '1', '何燕', '妈妈', '50', '10086', null, '10', '80', '康泰园', '', '2018-04-10 12:57:25', '2019-08-23 03:14:59', '1', '', '', '2640.00', '12.00', '1920.00', '24.00', '0.00', '0.00');
INSERT INTO `gms_student` VALUES ('12', '沈昊', 'st000011', '0', '1', '邓云', '妈妈', '40', '10086', null, '10', '82', '亚运城', '', '2018-04-10 12:58:00', '2019-08-23 03:14:36', '1', '', '', '8480.00', '53.00', '19240.00', '148.00', '0.00', '0.00');
INSERT INTO `gms_student` VALUES ('13', '张起', 'st000012', '0', '1', '刘颖', '妈妈', '60', '10086', null, '10', '84', '珠江帝景', '', '2018-04-10 12:59:39', '2019-08-23 07:41:08', '1', '', '', '10800.00', '54.00', '1620.00', '18.00', '0.00', '0.00');
INSERT INTO `gms_student` VALUES ('14', '高振', 'st000013', '0', '1', '顾颖', '妈妈', '30', '10086', null, '10', '80', '明月居', '', '2018-04-10 13:00:17', '2019-08-23 07:40:56', '1', '', '', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00');
INSERT INTO `gms_student` VALUES ('15', '齐钰', 'st000014', '0', '1', '齐玉', '爸爸', '50', '10086', null, '10', '87', '东豪大厦', '', '2018-04-10 13:02:38', '2019-08-23 07:40:51', '1', '', '', '5600.00', '28.00', '0.00', '0.00', '0.00', '0.00');
INSERT INTO `gms_student` VALUES ('16', '王二', 'st000015', '0', '1', '王大', '爸爸', '90', '10086', null, '10', '80', '珠江帝景', '', '2018-04-10 13:11:10', '2019-08-23 07:40:45', '1', '', '', '300.00', '1.00', '1360.00', '12.00', '0.00', '0.00');
INSERT INTO `gms_student` VALUES ('17', '突击英语', 'sh0001', '1', '0', '', '', '0', '', '', '10', '0', '', '', '2019-08-23 07:43:37', null, '1', '[\"11\",\"12\"]', '李永|沈昊', '0.00', '0.00', '0.00', '0.00', '0.00', '0.00');

-- ----------------------------
-- Table structure for gms_student_area
-- ----------------------------
DROP TABLE IF EXISTS `gms_student_area`;
CREATE TABLE `gms_student_area` (
  `id` smallint(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Records of gms_student_area
-- ----------------------------
INSERT INTO `gms_student_area` VALUES ('2', '珠江帝景');
INSERT INTO `gms_student_area` VALUES ('3', '东豪大厦');
INSERT INTO `gms_student_area` VALUES ('6', '明月居');
INSERT INTO `gms_student_area` VALUES ('7', '康泰园');
INSERT INTO `gms_student_area` VALUES ('20', '亚运城');
INSERT INTO `gms_student_area` VALUES ('22', '保利花城');
INSERT INTO `gms_student_area` VALUES ('23', '珠江花园');

-- ----------------------------
-- Table structure for gms_student_school
-- ----------------------------
DROP TABLE IF EXISTS `gms_student_school`;
CREATE TABLE `gms_student_school` (
  `id` smallint(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '学校名',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=99 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Records of gms_student_school
-- ----------------------------
INSERT INTO `gms_student_school` VALUES ('80', '实验小学');
INSERT INTO `gms_student_school` VALUES ('82', '实验中学');
INSERT INTO `gms_student_school` VALUES ('84', '师范大学附属高中');
INSERT INTO `gms_student_school` VALUES ('85', '培英中学');
INSERT INTO `gms_student_school` VALUES ('87', '育才小学');
INSERT INTO `gms_student_school` VALUES ('93', '六一小学');

-- ----------------------------
-- Table structure for gms_subject
-- ----------------------------
DROP TABLE IF EXISTS `gms_subject`;
CREATE TABLE `gms_subject` (
  `id` smallint(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `name` (`name`) USING BTREE,
  KEY `name_2` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of gms_subject
-- ----------------------------
INSERT INTO `gms_subject` VALUES ('11', '化学');
INSERT INTO `gms_subject` VALUES ('13', '历史');
INSERT INTO `gms_subject` VALUES ('20', '口才');
INSERT INTO `gms_subject` VALUES ('14', '地理');
INSERT INTO `gms_subject` VALUES ('19', '托班');
INSERT INTO `gms_subject` VALUES ('12', '政治');
INSERT INTO `gms_subject` VALUES ('8', '数学');
INSERT INTO `gms_subject` VALUES ('10', '物理');
INSERT INTO `gms_subject` VALUES ('15', '生物');
INSERT INTO `gms_subject` VALUES ('9', '英语');
INSERT INTO `gms_subject` VALUES ('7', '语文');

-- ----------------------------
-- Table structure for gms_teacher
-- ----------------------------
DROP TABLE IF EXISTS `gms_teacher`;
CREATE TABLE `gms_teacher` (
  `id` smallint(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '教师姓名',
  `idcard` varchar(30) DEFAULT '' COMMENT '身份证',
  `cardno` varchar(30) DEFAULT '' COMMENT '卡号',
  `phone` varchar(20) DEFAULT '' COMMENT '联系电话',
  `school_id` smallint(10) NOT NULL COMMENT '校区ID',
  `department` varchar(30) DEFAULT '' COMMENT '部门',
  `contacter` varchar(15) DEFAULT NULL,
  `contacter_phone` varchar(15) DEFAULT NULL,
  `sex` tinyint(5) DEFAULT '0' COMMENT '性别0未知1男2女',
  `picture` varchar(80) DEFAULT '' COMMENT '照片',
  `birth` date DEFAULT NULL COMMENT '生日',
  `entry_time` date DEFAULT NULL COMMENT '入职时间',
  `contract_expiration_time` date DEFAULT NULL COMMENT '合同到期时间',
  `price` decimal(8,2) DEFAULT '0.00' COMMENT '课时费',
  `reward` decimal(8,2) DEFAULT '0.00' COMMENT '奖金',
  `address` varchar(80) DEFAULT '' COMMENT '地址',
  `created_at` datetime DEFAULT NULL,
  `note` varchar(50) DEFAULT '' COMMENT '备注',
  `updated_at` datetime DEFAULT NULL,
  `status` tinyint(5) NOT NULL DEFAULT '0' COMMENT '状态',
  `price10` int(11) DEFAULT '0',
  `price20` int(11) DEFAULT '0',
  `price30` int(11) DEFAULT '0',
  `price40` int(11) DEFAULT '0',
  `price50` int(11) DEFAULT '0',
  `price60` int(11) DEFAULT '0',
  `price70` int(11) DEFAULT '0',
  `price80` int(11) DEFAULT '0',
  `price90` int(11) DEFAULT '0',
  `price100` int(11) DEFAULT '0',
  `price110` int(11) DEFAULT '0',
  `price120` int(11) DEFAULT '0',
  `social_security` varchar(30) DEFAULT '' COMMENT '社保',
  `funds` varchar(30) DEFAULT '' COMMENT '公积金',
  `salary_card` varchar(40) DEFAULT '' COMMENT '工资卡号',
  `balance` decimal(8,2) unsigned DEFAULT '0.00',
  `price1` int(11) NOT NULL DEFAULT '0',
  `price2` int(11) NOT NULL DEFAULT '0',
  `price3` int(11) NOT NULL DEFAULT '0',
  `xprice1` int(11) NOT NULL,
  `xprice2` int(11) NOT NULL,
  `xprice3` int(11) NOT NULL,
  `xprice10` int(11) NOT NULL,
  `xprice20` int(11) NOT NULL,
  `xprice30` int(11) NOT NULL,
  `xprice40` int(11) NOT NULL,
  `xprice50` int(11) NOT NULL,
  `xprice60` int(11) NOT NULL,
  `xprice70` int(11) NOT NULL,
  `xprice80` int(11) NOT NULL,
  `xprice90` int(11) NOT NULL,
  `xprice100` int(11) NOT NULL,
  `xprice110` int(11) NOT NULL,
  `xprice120` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `cardno` (`cardno`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Records of gms_teacher
-- ----------------------------
INSERT INTO `gms_teacher` VALUES ('1', '华罗庚', '434222199905123212', 'tc000001', '10010', '10', '', '李强', '10010', '1', '', '2018-03-01', '2018-03-08', '2018-03-08', '100.00', '0.00', '', '0000-00-00 00:00:00', '', '0000-00-00 00:00:00', '0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '', '', '1111', '0.00', '1', '1', '1', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0');
INSERT INTO `gms_teacher` VALUES ('2', '陈独秀', '434222199905123212', 'tc000002', '10010', '10', '', '', '10010', '0', '', '2018-03-01', '2018-03-21', '2018-03-21', '0.00', '0.00', null, null, '', '2018-03-22 14:03:04', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '', '', '', '0.00', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0');
INSERT INTO `gms_teacher` VALUES ('3', '李阳', '434222199905123212', 'tc000003', '10086', '10', '', '', '10010', '0', '', '2018-03-23', '1970-01-01', '1970-01-01', '0.00', '0.00', null, null, '', '2018-03-22 14:02:59', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '', '', '', '0.00', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0');
INSERT INTO `gms_teacher` VALUES ('4', '吴刚', '434222199905123212', 'tc000004', '10086', '10', null, '崔凯', '10010', '1', '', '1993-06-19', '2017-11-20', '2023-11-19', '75.00', '0.00', '杨浦区世界路200号', null, '', '2018-08-31 20:01:13', '0', '50', '50', '50', '50', '50', '55', '55', '55', '60', '65', '65', '70', '', '181291472205', '6222031001006938901', '110.00', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0');
INSERT INTO `gms_teacher` VALUES ('5', '郑志平', '434222199905123212', 'tc000005', '10086', '10', null, '张若东', '10010', '2', '', '1994-04-23', '2017-11-20', '2023-11-19', '25.00', '0.00', '控江西三村129号14室', null, '', '2018-06-10 18:05:57', '0', '50', '50', '50', '50', '50', '55', '55', '55', '60', '65', '65', '70', '', '182447625205', '6222031001006940683', '225.00', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0');
INSERT INTO `gms_teacher` VALUES ('6', '石磊', '434222199905123212', 'tc000006', '10086', '10', null, '姜玉', '10010', '1', '', '1986-05-05', '2017-11-20', '2023-11-19', '45.00', '5.00', '翔殷路578弄', null, '', '2018-06-06 19:26:39', '0', '70', '55', '70', '70', '70', '75', '75', '75', '80', '85', '85', '90', '', '183299322205', '6222031001007320455', '59382.50', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0');
INSERT INTO `gms_teacher` VALUES ('7', '吴月', '434222199905123212', 'tc000007', '10086', '10', null, '范东艳', '10010', '2', '', '1994-06-21', '2017-11-20', '2023-11-19', '45.00', '0.00', '杨浦区长白三村7号508', null, '', null, '0', '70', '70', '70', '70', '70', '75', '75', '75', '80', '85', '85', '90', '', '188994348205', '6222081001020019562', '50590.00', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0');
INSERT INTO `gms_teacher` VALUES ('8', '刘强', '434222199905123212', 'tc000008', '10086', '10', null, '孙振林', '10010', '2', '', '1993-01-07', '2017-11-20', '2023-11-19', '50.00', '10.00', '上海市', null, '', null, '0', '80', '80', '80', '80', '80', '85', '85', '85', '90', '95', '95', '100', '', '177625545205', '6222031001006746205', '93292.50', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0');

-- ----------------------------
-- Table structure for gms_teacher_school
-- ----------------------------
DROP TABLE IF EXISTS `gms_teacher_school`;
CREATE TABLE `gms_teacher_school` (
  `id` smallint(10) unsigned NOT NULL AUTO_INCREMENT,
  `teacher_id` smallint(10) NOT NULL DEFAULT '0',
  `school_id` smallint(10) NOT NULL DEFAULT '0',
  `school_name` varchar(30) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of gms_teacher_school
-- ----------------------------
INSERT INTO `gms_teacher_school` VALUES ('7', '4', '10', '');
INSERT INTO `gms_teacher_school` VALUES ('8', '5', '10', '');
INSERT INTO `gms_teacher_school` VALUES ('9', '6', '10', '');
INSERT INTO `gms_teacher_school` VALUES ('10', '7', '10', '');
INSERT INTO `gms_teacher_school` VALUES ('11', '8', '10', '');
INSERT INTO `gms_teacher_school` VALUES ('12', '11', '10', '');
INSERT INTO `gms_teacher_school` VALUES ('13', '12', '10', '');
INSERT INTO `gms_teacher_school` VALUES ('14', '13', '10', '');
INSERT INTO `gms_teacher_school` VALUES ('15', '14', '10', '');
INSERT INTO `gms_teacher_school` VALUES ('16', '15', '10', '');
INSERT INTO `gms_teacher_school` VALUES ('17', '16', '10', '');
INSERT INTO `gms_teacher_school` VALUES ('18', '18', '10', '');
INSERT INTO `gms_teacher_school` VALUES ('19', '17', '10', '');

-- ----------------------------
-- Table structure for gms_teacher_subject
-- ----------------------------
DROP TABLE IF EXISTS `gms_teacher_subject`;
CREATE TABLE `gms_teacher_subject` (
  `id` smallint(10) unsigned NOT NULL AUTO_INCREMENT,
  `teacher_id` smallint(10) NOT NULL DEFAULT '0' COMMENT '教师ID',
  `subject_name` varchar(10) DEFAULT '' COMMENT '学科',
  `subject_id` smallint(10) NOT NULL DEFAULT '0' COMMENT '学科编号',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `teacher_id` (`teacher_id`) USING BTREE,
  KEY `subject_id` (`subject_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=45 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of gms_teacher_subject
-- ----------------------------
INSERT INTO `gms_teacher_subject` VALUES ('6', '8', '', '9');
INSERT INTO `gms_teacher_subject` VALUES ('9', '5', '', '7');
INSERT INTO `gms_teacher_subject` VALUES ('10', '6', '', '8');
INSERT INTO `gms_teacher_subject` VALUES ('11', '7', '', '8');
INSERT INTO `gms_teacher_subject` VALUES ('12', '7', '', '11');
INSERT INTO `gms_teacher_subject` VALUES ('13', '11', '', '8');
INSERT INTO `gms_teacher_subject` VALUES ('14', '12', '', '19');
INSERT INTO `gms_teacher_subject` VALUES ('15', '13', '', '19');
INSERT INTO `gms_teacher_subject` VALUES ('16', '14', '', '19');
INSERT INTO `gms_teacher_subject` VALUES ('17', '15', '', '9');
INSERT INTO `gms_teacher_subject` VALUES ('18', '16', '', '9');
INSERT INTO `gms_teacher_subject` VALUES ('20', '18', '', '8');
INSERT INTO `gms_teacher_subject` VALUES ('21', '17', '', '7');
INSERT INTO `gms_teacher_subject` VALUES ('22', '7', '', '10');
INSERT INTO `gms_teacher_subject` VALUES ('23', '19', '', '10');
INSERT INTO `gms_teacher_subject` VALUES ('24', '20', '', '7');
INSERT INTO `gms_teacher_subject` VALUES ('25', '21', '', '7');
INSERT INTO `gms_teacher_subject` VALUES ('26', '24', '', '9');
INSERT INTO `gms_teacher_subject` VALUES ('27', '25', '', '8');
INSERT INTO `gms_teacher_subject` VALUES ('30', '26', '', '7');
INSERT INTO `gms_teacher_subject` VALUES ('31', '26', '', '8');
INSERT INTO `gms_teacher_subject` VALUES ('32', '26', '', '9');
INSERT INTO `gms_teacher_subject` VALUES ('33', '26', '', '10');
INSERT INTO `gms_teacher_subject` VALUES ('34', '26', '', '12');
INSERT INTO `gms_teacher_subject` VALUES ('35', '26', '', '13');
INSERT INTO `gms_teacher_subject` VALUES ('36', '26', '', '14');
INSERT INTO `gms_teacher_subject` VALUES ('37', '28', '', '8');
INSERT INTO `gms_teacher_subject` VALUES ('38', '28', '', '7');
INSERT INTO `gms_teacher_subject` VALUES ('39', '28', '', '9');
INSERT INTO `gms_teacher_subject` VALUES ('40', '28', '', '10');
INSERT INTO `gms_teacher_subject` VALUES ('41', '28', '', '11');
INSERT INTO `gms_teacher_subject` VALUES ('42', '28', '', '14');
INSERT INTO `gms_teacher_subject` VALUES ('43', '1', '', '7');
INSERT INTO `gms_teacher_subject` VALUES ('44', '2', '', '7');

-- ----------------------------
-- Table structure for gms_us
-- ----------------------------
DROP TABLE IF EXISTS `gms_us`;
CREATE TABLE `gms_us` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '',
  `content` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `sort` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of gms_us
-- ----------------------------
INSERT INTO `gms_us` VALUES ('1', '公司介绍', '<p>公司介绍1</p>\n', '2019-07-10 23:26:01', '0');
INSERT INTO `gms_us` VALUES ('2', '公司简介', '<p><span style=\"font-size:20px\">&nbsp; &nbsp; &nbsp; &nbsp;上海墨微教育自创办以来，以良好的社会声誉赢得了家长的信赖，是学生校外成长的优质基地，是帮助家长解决后顾之忧的好助手。公司视打造优质的教育团队为发展根本，在拥有骨干教师的基础上，吸收了一批师范类高校的优秀毕业生担任辅导任务，他们精力充沛，教育理念先进，和学生沟通零距离。同时，为提高辅导水平，聘请名师对青年教师进行岗前培训，跟踪培养，形成了知识结构完整，年龄结构合理辅导团队。目前，主要开设文化课辅导、语言艺术培训、篮球训练营以及书法培训。我们本着&ldquo;教而有方，严而有格&rdquo;的教育理念对学生的发展科学指导，在把握好尺度的基础上对学生的学习习惯严格要求，为学生全面发展打下坚实基础。</span></p>\n', '2019-08-15 05:23:24', '0');
