/*
Navicat MySQL Data Transfer

Source Server         : 惠居测试数据库
Source Server Version : 50173
Source Host           : 101.201.150.0:3306
Source Database       : doc

Target Server Type    : MYSQL
Target Server Version : 50173
File Encoding         : 65001

Date: 2016-01-20 14:55:36
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for account
-- ----------------------------
DROP TABLE IF EXISTS `account`;
CREATE TABLE `account` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `account` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `nick_name` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for catalogue
-- ----------------------------
DROP TABLE IF EXISTS `catalogue`;
CREATE TABLE `catalogue` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) NOT NULL COMMENT '父目录，没有为0',
  `name` varchar(255) NOT NULL COMMENT '目录名称',
  `serial_number` varchar(255) NOT NULL COMMENT '序号',
  `project_id` int(11) NOT NULL DEFAULT '0' COMMENT '所属项目',
  `version_id` int(11) NOT NULL DEFAULT '0' COMMENT '所属版本号',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8 COMMENT='文档目录';

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
  `project_id` int(11) NOT NULL DEFAULT '0' COMMENT '所属项目ID',
  `version_id` int(11) NOT NULL DEFAULT '0' COMMENT '所属版本号',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=47 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for err_code
-- ----------------------------
DROP TABLE IF EXISTS `err_code`;
CREATE TABLE `err_code` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `code` varchar(255) NOT NULL COMMENT '错误代码',
  `description_text` text NOT NULL COMMENT '描述',
  `project_id` int(11) NOT NULL DEFAULT '0' COMMENT '所属项目ID',
  `version_id` int(11) NOT NULL DEFAULT '0' COMMENT '所属版本号',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for explain
-- ----------------------------
DROP TABLE IF EXISTS `explain`;
CREATE TABLE `explain` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL COMMENT '标题',
  `description_text` text NOT NULL COMMENT '描述',
  `user_name` varchar(255) NOT NULL COMMENT '发布人',
  `created_at` datetime NOT NULL,
  `project_id` int(11) NOT NULL DEFAULT '0' COMMENT '所属项目ID',
  `version_id` int(11) NOT NULL COMMENT '版本ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8;

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
  `project_id` int(11) NOT NULL DEFAULT '0' COMMENT '所属项目ID',
  `version_id` int(11) NOT NULL DEFAULT '0' COMMENT '所属版本号',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=237 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for project
-- ----------------------------
DROP TABLE IF EXISTS `project`;
CREATE TABLE `project` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL COMMENT '项目名称',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for update_log
-- ----------------------------
DROP TABLE IF EXISTS `update_log`;
CREATE TABLE `update_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `description_text` text NOT NULL COMMENT '更新说明',
  `created_at` datetime NOT NULL,
  `project_id` int(11) NOT NULL DEFAULT '0' COMMENT '所属项目ID',
  `version_id` int(11) NOT NULL COMMENT '版本的ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=75 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for version
-- ----------------------------
DROP TABLE IF EXISTS `version`;
CREATE TABLE `version` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `version` varchar(255) NOT NULL COMMENT '版本号',
  `project_id` int(11) NOT NULL COMMENT '项目的ID',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
