/*
 Navicat Premium Data Transfer

 Source Server         : laragon
 Source Server Type    : MySQL
 Source Server Version : 50724
 Source Host           : localhost:3306
 Source Schema         : go

 Target Server Type    : MySQL
 Target Server Version : 50724
 File Encoding         : 65001

 Date: 19/05/2022 18:07:24
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for blog_article
-- ----------------------------
DROP TABLE IF EXISTS `blog_article`;
CREATE TABLE `blog_article`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '文章标题',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '文章简述',
  `cover_image_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '封面图片地址',
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '文章内容',
  `created_on` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '创建时间',
  `created_by` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '修改时间',
  `modified_by` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '删除时间',
  `is_del` tinyint(3) UNSIGNED NULL DEFAULT 0 COMMENT '是否删除 0 为未删除、1 为已删除',
  `state` tinyint(3) UNSIGNED NULL DEFAULT 1 COMMENT '状态 0 为禁用、1 为启用',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文章管理' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of blog_article
-- ----------------------------
INSERT INTO `blog_article` VALUES (1, '正夫口腔正畸总院长刘从华主任受邀出席“广东省民营牙科协会口腔医生正畸专业能力评估委员会成立大会”并聘任为委员专家', '为促进口腔正畸技术服务规范，提高口腔正畸治疗的质量和保证医疗安全，特别邀请正夫口腔正畸总院长刘从华主任出席大会成立仪式。', '/images/admin/upload/20220107/5.jpg', '蔡斌主任表示，由于正畸学科专业性强、涉及范围广、难度大，属于继续教育和研究生教育课程，使得接收过正规化教育的正畸医生较少，出现口腔正畸专业发展不平衡，水平参差不齐，为促进口腔正畸技术服务规范，提高口腔正畸治疗的质量和保证医疗安全，保障群众的口腔健康，特成立口腔医生正畸专业能力评估委员会。委员会采用宽进严出的审核制度，面向全省民营口腔有意从事正畸临床工作医师，发放正畸专业能力评估流程，接受全省民营口腔医师报名。', 1651029656, '刘粤新', 1651029656, '', 0, 0, 1);
INSERT INTO `blog_article` VALUES (2, '热烈欢迎 | 广东财经大学深圳校友会参访正夫口腔！', '2021年12月31日下午14:00，广东财经大学深圳校友会参访正夫口腔。此次活动采用参观与讲座相结合的形式，让大家更好地认知医疗口腔行业的市场现状及发展趋势，深度挖掘医疗行业的投资价值。', '/images/admin/upload/20220107/5.jpg', '2021年12月31日下午14:00，广东财经大学深圳校友会参访正夫口腔。此次活动采用参观与讲座相结合的形式，让大家更好地认知医疗口腔行业的市场现状及发展趋势，深度挖掘医疗行业的投资价值。\n\nhttp://www.zfck.net/images/admin/upload/20220104/微信图片_20220104113558(1).jpg\n\nhttp://www.zfck.net/images/admin/upload/20220104/微信图片_20220104113607(1).jpg\n\n校友们陆续签到入场后，由正夫俱乐部主席吴剑兴先生、正夫口腔中洲明星店总经理兼董事长助理黄艳女士带领校友们一起参观了中洲明星旗舰店，并逐一介绍了正夫口腔的发展历程、医护团队建设、医疗设备等情况，校友们对正夫口腔高品质高技术的医疗团队，专业的服务精神给予了极高的评价。', 1651029743, '刘粤新', 1651029743, '', 0, 0, 1);
INSERT INTO `blog_article` VALUES (3, '重磅消息 | 种植大咖张帅玉院长强势加盟正夫口腔！', '2021年12月31日，正夫口腔为深圳广大疑难复杂种植顾客带来福音，一位重量级种植口腔大咖强势加盟正夫口腔，他就是北京大学硕士研究生，从事口腔临床工作二十余年并深受患者信任与喜爱的张帅玉院长。', '/images/admin/upload/20220107/5.jpg', '张帅玉院长是北京大学硕士，从事口腔临床工作二十余年，成功完成超一万例种植案例，曾多次赴日本明海大学、朝日大学及德国法兰克福大学等知名院校学习深造口腔种植、口腔修复等技术，拥有国际先进的医疗理念。赴全国各地进行实操技能培训及讲座百余次，受益学员数千人，帮助数百名全科医师成功开展种植技术。擅长各类复杂牙缺失的美学种植修复技术及种植并发症处理。具有扎实的外科种植手术经验、牙周颌面学经验及精湛的修复技术，是深圳最早开展种植牙技术的资深种植专科医师。', 1651029815, '刘粤新', 1651029815, '', 0, 0, 1);
INSERT INTO `blog_article` VALUES (4, '正夫口腔受邀出席“第三届深圳新三板50强高峰论坛”！', '新三板是我国中小企业赖以发展的重要资本市场，在改善中小企业金融环境，大力推动创新、创业方面扮演着重要角色。深圳新三板企业作为中国产业的生力军，新生代创新力量，拥有独特的发展韧性。', '/images/admin/upload/20220107/5.jpg', '在圆桌论坛上，当主持人问及杨瑞锋董事长：正夫口腔从2014年11月注册成立，2015年1月开第一家口腔门店，到2017年8月上新三板挂牌，之后又让企鹅医生全资收购，正夫口腔受投资者青睐的原因是什么?\n\n杨瑞锋董事长表示：正夫口腔能从众多民营口腔机构企业中脱颖而出，短时间内获得投资者的青睐，一方面是因为正夫口腔拥有好的产品加好的模式双管齐下。在产品的选择上，正夫口腔以丰富的品种、过硬的品质、驰名的品牌、高标准的供给质量，满足人民美好生活日益增长的个性化、多元化消费需求。医疗安全是口腔管理的核心，正夫口腔不断完善健全医疗安全管理机制，旨在提高口腔服务质量。另一方面则离不开自身发展的优势，投资者所青睐的一定是具备“快、大、有未来”这三点的企业，而正夫口腔在短短几年间就在深圳开了30多家口腔连锁门诊，足见其强大的综合实力。同时，正夫口腔一直坚信重视专业人才的引进和培养是医疗质量的保障，随着高端医疗人才队伍不断扩大，拥有逐渐形成优质人才的资源合力。在医疗消费场景不断升级，老百姓对医疗健康的需求与日俱增，正夫口腔始终坚持创新，紧跟市场需求，不断成长，成为具有核心竞争力的企业。', 1651029854, '刘粤新', 1651029854, '', 1651031507, 1, 1);

-- ----------------------------
-- Table structure for blog_article_tag
-- ----------------------------
DROP TABLE IF EXISTS `blog_article_tag`;
CREATE TABLE `blog_article_tag`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `article_id` int(11) NOT NULL COMMENT '文章 ID',
  `tag_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '标签 ID',
  `created_on` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '创建时间',
  `created_by` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '修改时间',
  `modified_by` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '删除时间',
  `is_del` tinyint(3) UNSIGNED NULL DEFAULT 0 COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文章标签关联' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of blog_article_tag
-- ----------------------------

-- ----------------------------
-- Table structure for blog_auth
-- ----------------------------
DROP TABLE IF EXISTS `blog_auth`;
CREATE TABLE `blog_auth`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `app_key` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'Key',
  `app_secret` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'Secret',
  `created_on` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '创建时间',
  `created_by` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '修改时间',
  `modified_by` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '删除时间',
  `is_del` tinyint(3) UNSIGNED NULL DEFAULT 0 COMMENT '是否删除 0 为未删除、1 为已删除',
  `state` tinyint(3) UNSIGNED NULL DEFAULT 1 COMMENT '状态 0 为禁用、1 为启用',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '认证管理' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of blog_auth
-- ----------------------------
INSERT INTO `blog_auth` VALUES (1, 'admin', 'admin', 0, 'eddycjy', 0, '', 0, 0, 1);

-- ----------------------------
-- Table structure for blog_tag
-- ----------------------------
DROP TABLE IF EXISTS `blog_tag`;
CREATE TABLE `blog_tag`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '标签名称',
  `created_on` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '创建时间',
  `created_by` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '修改时间',
  `modified_by` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '删除时间',
  `is_del` tinyint(3) UNSIGNED NULL DEFAULT 0 COMMENT '是否删除 0 为未删除、1 为已删除',
  `state` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '状态 0 为禁用、1 为启用',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '标签管理' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of blog_tag
-- ----------------------------
INSERT INTO `blog_tag` VALUES (1, '种植牙', 1650963883, '刘粤新', 1650963883, '', 0, 0, 1);
INSERT INTO `blog_tag` VALUES (2, '正畸矫正', 1650963926, '刘粤新', 1650963926, '', 0, 0, 1);
INSERT INTO `blog_tag` VALUES (3, '拔牙', 1650963957, '刘粤新', 1650963957, '', 0, 0, 1);
INSERT INTO `blog_tag` VALUES (4, '补牙', 1650963961, '刘粤新', 1650963961, '', 0, 0, 1);
INSERT INTO `blog_tag` VALUES (5, '烤瓷牙', 1650963966, '刘粤新', 1650963966, '', 0, 0, 1);
INSERT INTO `blog_tag` VALUES (6, '根管治疗', 1650964002, '刘粤新', 1650964002, '', 0, 0, 1);
INSERT INTO `blog_tag` VALUES (7, '牙周治疗', 1650964036, '刘粤新', 1650964036, '', 0, 0, 1);
INSERT INTO `blog_tag` VALUES (8, '美学贴面修复', 1650964050, '刘粤新', 1650964050, '', 0, 0, 1);
INSERT INTO `blog_tag` VALUES (9, '美白', 1650964075, '刘粤新', 1650964075, '', 0, 0, 1);
INSERT INTO `blog_tag` VALUES (10, '测试', 1650964658, '刘粤新', 1651649930, '刘先生', 1650964896, 1, 1);

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
  `p_type` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  INDEX `IDX_casbin_rule_p_type`(`p_type`) USING BTREE,
  INDEX `IDX_casbin_rule_v0`(`v0`) USING BTREE,
  INDEX `IDX_casbin_rule_v1`(`v1`) USING BTREE,
  INDEX `IDX_casbin_rule_v2`(`v2`) USING BTREE,
  INDEX `IDX_casbin_rule_v3`(`v3`) USING BTREE,
  INDEX `IDX_casbin_rule_v4`(`v4`) USING BTREE,
  INDEX `IDX_casbin_rule_v5`(`v5`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/admin/api/v1/policy', 'GET', '', NULL, NULL);

SET FOREIGN_KEY_CHECKS = 1;
