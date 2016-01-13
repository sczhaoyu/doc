/*
Navicat MySQL Data Transfer

Source Server         : hj
Source Server Version : 50616
Source Host           : 10.0.0.252:3306
Source Database       : doc

Target Server Type    : MYSQL
Target Server Version : 50616
File Encoding         : 65001

Date: 2016-01-13 18:17:29
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for catalogue
-- ----------------------------
DROP TABLE IF EXISTS `catalogue`;
CREATE TABLE `catalogue` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) NOT NULL COMMENT '父目录，没有为0',
  `name` varchar(255) NOT NULL COMMENT '目录名称',
  `serial_number` varchar(255) NOT NULL COMMENT '序号',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COMMENT='文档目录';

-- ----------------------------
-- Table structure for doc
-- ----------------------------
DROP TABLE IF EXISTS `doc`;
CREATE TABLE `doc` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `catalogue_id` int(11) NOT NULL COMMENT '所属目录',
  `path` varchar(255) NOT NULL COMMENT '接口请求路径',
  `description_text` text NOT NULL COMMENT '接口介绍',
  `input_demo` text NOT NULL COMMENT '输入示例',
  `out_demo` text NOT NULL COMMENT '输出示例',
  `name` varchar(255) NOT NULL,
  `serial_number` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for err_code
-- ----------------------------
DROP TABLE IF EXISTS `err_code`;
CREATE TABLE `err_code` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `code` varchar(255) NOT NULL COMMENT '错误代码',
  `description_text` text NOT NULL COMMENT '描述',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for parameters
-- ----------------------------
DROP TABLE IF EXISTS `parameters`;
CREATE TABLE `parameters` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL COMMENT '参数名称',
  `data_type` varchar(255) NOT NULL COMMENT '数据类型',
  `description_text` text NOT NULL COMMENT '参数描述',
  `required` tinyint(11) NOT NULL COMMENT '是否必选参数，0否，1是',
  `doc_id` int(11) NOT NULL COMMENT '所属文档,doc表外键',
  `prm_type` tinyint(4) NOT NULL COMMENT '0请求参数1响应参数',
  `length` varchar(11) NOT NULL,
  `serial_number` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=43 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for update_log
-- ----------------------------
DROP TABLE IF EXISTS `update_log`;
CREATE TABLE `update_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `description_text` text NOT NULL COMMENT '更新说明',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
