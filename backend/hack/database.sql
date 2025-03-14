DROP DATABASE IF EXISTS `pprof_cloud_platform`;
CREATE DATABASE pprof_cloud_platform;
USE pprof_cloud_platform;
-- 用户表
CREATE TABLE users (
    user_id INT PRIMARY KEY AUTO_INCREMENT, 
    email VARCHAR(255) NOT NULL UNIQUE,
    hashed_password VARCHAR(255) NOT NULL,
    permission ENUM('admin', 'user') DEFAULT 'user',
    status ENUM('active', 'inactive') DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

-- 项目表
CREATE TABLE project (
    project_id INT PRIMARY KEY AUTO_INCREMENT,
    creator_id INT,
    project_name VARCHAR(255) NOT NULL,
    project_description TEXT,
    project_token VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (creator_id) REFERENCES users(user_id) ON DELETE CASCADE
);

-- 微服务表
CREATE TABLE microservice (
    microservice_id INT PRIMARY KEY AUTO_INCREMENT,
    project_id INT,
    creator_id INT,
    microservice_name VARCHAR(255) NOT NULL,
    ip VARCHAR(15) NOT NULL,
    port INT NOT NULL,
    microservice_description TEXT,
    microservice_token VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (project_id) REFERENCES project(project_id) ON DELETE CASCADE
);

-- 测试记录
-- Create profile table
CREATE TABLE profile (
    id INT PRIMARY KEY AUTO_INCREMENT,
    microservice_id INT,
    project_id INT,
    ptype VARCHAR(255) NOT NULL,
    oss_path VARCHAR(255)  ,
    comment VARCHAR(255)  ,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (microservice_id) REFERENCES microservice(microservice_id)
);


-- 测试数据
-- 生成1个用户,密码123456
INSERT INTO users (email, hashed_password, permission, status) VALUES
('user1@example.com', '$2a$10$djutXo6qKRBfMY.X1YWrZOokzghdsUyNA1Zhgtsd5lHYGEU6Wl3e6', 'user', 'active');


-- 生成10个项目（用用户ID=1）
INSERT INTO project (creator_id, project_name, project_description, project_token) VALUES
(1, '电商平台', '分布式电商系统', 'prj_ec123'),
(1, 'IoT监控', '工业物联网监控平台', 'prj_iot456'),
(1, '在线教育', '实时互动教学系统', 'prj_edu789'),
(1, '医疗系统', '电子病历管理系统', 'prj_med101'),
(1, '社交网络', '去中心化社交平台', 'prj_soc202'),
(1, '物流跟踪', '智能物流追踪系统', 'prj_log303'),
(1, '金融服务', '微服务化支付系统', 'prj_fin404'),
(1, '游戏后台', 'MMORPG游戏服务器集群', 'prj_gam505'),
(1, '内容管理', '多租户CMS平台', 'prj_cms606'),
(1, '智能家居', '家庭自动化控制系统', 'prj_home707');

-- 为每个项目生成10个微服务（项目ID 1-10）
-- 项目1的微服务
INSERT INTO microservice (project_id, creator_id, microservice_name, ip, port, microservice_description, microservice_token) VALUES
-- 项目1 (ID=1)
(1, 1, '用户服务', '10.0.0.101', 8080, '用户注册登录管理','mis_111'),
(1, 1, '商品服务', '10.0.0.102', 8081, '商品信息管理','mis_112'),
(1, 1, '订单服务', '10.0.0.103', 8082, '订单处理系统','mis_113'),
(1, 1, '支付服务', '10.0.0.104', 8083, '集成支付网关','mis_114'),
(1, 1, '库存服务', '10.0.0.105', 8084, '实时库存管理','mis_115'),
(1, 1, '推荐服务', '10.0.0.106', 8085, '个性化推荐引擎','mis_116'),
(1, 1, '搜索服务', '10.0.0.107', 8086, '商品搜索引擎','mis_117'),
(1, 1, '日志服务', '10.0.0.108', 8087, '分布式日志采集','mis_118'),
(1, 1, '监控服务', '10.0.0.109', 8088, '系统健康监测','mis_119'),
(1, 1, '网关服务', '10.0.0.110', 8089, 'API网关路由','mis_1110'),

-- 项目2 (ID=2)
(2, 1, '设备管理', '192.168.1.101', 5000, 'IoT设备接入管理','mis_21'),
(2, 1, '数据分析', '192.168.1.102', 5001, '时序数据分析引擎','mis_22'),
(2, 1, '告警服务', '192.168.1.103', 5002, '实时异常检测','mis_23'),
(2, 1, '边缘计算', '192.168.1.104', 5003, '边缘节点计算','mis_24'),
(2, 1, '协议转换', '192.168.1.105', 5004, 'MQTT/CoAP协议转换','mis_251'),
(2, 1, '视频处理', '192.168.1.106', 5005, '实时视频流分析','mis_321'),
(2, 1, '位置服务', '192.168.1.107', 5006, '设备地理围栏','mis_1234'),
(2, 1, '固件管理', '192.168.1.108', 5007, 'OTA固件升级','mis_124'),
(2, 1, '规则引擎', '192.168.1.109', 5008, '业务规则编排','mis_165'),
(2, 1, '数据可视化', '192.168.1.110', 5009, '实时监控看板','mis_135'),

-- 项目3 (ID=3 在线教育系统)
(3, 1, '课程服务', '172.16.0.101', 3000, '在线课程管理模块','mis_21'),
(3, 1, '直播服务', '172.16.0.102', 3001, '实时视频互动教学','mis_21'),
(3, 1, '考试服务', '172.16.0.103', 3002, '在线考试与自动评分','mis_21'),
(3, 1, '资源服务', '172.16.0.104', 3003, '教学资源存储管理','mis_21'),
(3, 1, '讨论区服务', '172.16.0.105', 3004, '学习社区交流模块','mis_21'),
(3, 1, '认证服务', '172.16.0.106', 3005, '教育机构认证管理','mis_21'),
(3, 1, '排课服务', '172.16.0.107', 3006, '智能课程排期系统','mis_21'),
(3, 1, '消息推送', '172.16.0.108', 3007, '学习通知提醒服务','mis_21'),
(3, 1, '数据看板', '172.16.0.109', 3008, '学习行为分析系统','mis_21'),
(3, 1, '支付服务', '172.16.0.110', 3009, '课程购买支付接口','mis_21'),

-- 项目4 (ID=4 医疗系统)
(4, 1, '病历服务', '10.10.1.101', 4000, '电子病历管理系统','mis_21'),
(4, 1, '预约服务', '10.10.1.102', 4001, '门诊预约挂号系统','mis_21'),
(4, 1, '影像服务', '10.10.1.103', 4002, '医学影像存储分析','mis_21'),
(4, 1, '药房服务', '10.10.1.104', 4003, '药品库存管理系统','mis_21'),
(4, 1, '护理服务', '10.10.1.105', 4004, '护理工作流管理','mis_21'),
(4, 1, '检验服务', '10.10.1.106', 4005, '实验室检验报告系统','mis_21'),
(4, 1, '急救服务', '10.10.1.107', 4006, '急诊流程管理系统','mis_21'),
(4, 1, '移动查房', '10.10.1.108', 4007, '移动端病房管理系统','mis_21'),
(4, 1, '医保对接', '10.10.1.109', 4008, '医保结算接口服务','mis_21'),
(4, 1, '远程会诊', '10.10.1.110', 4009, '视频会诊协调服务','mis_21'),

-- 项目5 (ID=5 社交网络)
(5, 1, '动态服务', '192.168.2.101', 5000, '用户状态发布系统','mis_21'),
(5, 1, '好友服务', '192.168.2.102', 5001, '社交关系链管理','mis_21'),
(5, 1, '消息服务', '192.168.2.103', 5002, '即时通讯聊天系统','mis_21'),
(5, 1, '推荐服务', '192.168.2.104', 5003, '内容推荐引擎','mis_21'),
(5, 1, '内容审核', '192.168.2.105', 5004, 'AI内容安全过滤','mis_21'),
(5, 1, '直播服务', '192.168.2.106', 5005, '视频直播推流系统','mis_21'),
(5, 1, '存储服务', '192.168.2.107', 5006, '分布式文件存储','mis_21'),
(5, 1, '位置服务', '192.168.2.108', 5007, '地理围栏功能','mis_21'),
(5, 1, '活动服务', '192.168.2.109', 5008, '线上活动管理','mis_21'),
(5, 1, '支付服务', '192.168.2.110', 5009, '虚拟货币交易','mis_21'),

-- 项目6 (ID=6 物流跟踪)
(6, 1, '订单服务', '10.20.3.101', 6000, '物流订单管理系统','mis_21'),
(6, 1, '轨迹服务', '10.20.3.102', 6001, '实时位置追踪系统','mis_21'),
(6, 1, '仓储服务', '10.20.3.103', 6002, '智能仓储管理系统','mis_21'),
(6, 1, '运输服务', '10.20.3.104', 6003, '运输路线规划引擎','mis_21'),
(6, 1, '报关服务', '10.20.3.105', 6004, '跨境通关数据对接','mis_21'),
(6, 1, '结算服务', '10.20.3.106', 6005, '运费自动结算系统','mis_21'),
(6, 1, '预警服务', '10.20.3.107', 6006, '异常状态监测预警','mis_21'),
(6, 1, '设备服务', '10.20.3.108', 6007, '物联网设备管理','mis_21'),
(6, 1, '报表服务', '10.20.3.109', 6008, '物流数据分析报表','mis_21'),
(6, 1, '调度服务', '10.20.3.110', 6009, '智能车辆调度系统','mis_21'),

-- 项目7 (ID=7 金融服务)
(7, 1, '账户服务', '172.20.4.101', 7000, '资金账户管理系统','mis_21'),
(7, 1, '交易服务', '172.20.4.102', 7001, '支付交易处理核心','mis_21'),
(7, 1, '风控服务', '172.20.4.103', 7002, '实时风险控制系统','mis_21'),
(7, 1, '对账服务', '172.20.4.104', 7003, '财务自动对账系统','mis_21'),
(7, 1, '清算服务', '172.20.4.105', 7004, '资金清算结算系统','mis_21'),
(7, 1, '路由服务', '172.20.4.106', 7005, '支付渠道路由管理','mis_21'),
(7, 1, '通知服务', '172.20.4.107', 7006, '交易结果通知服务','mis_21'),
(7, 1, '商户服务', '172.20.4.108', 7007, '商户进件管理系统','mis_21'),
(7, 1, '账单服务', '172.20.4.109', 7008, '电子账单生成系统','mis_21'),
(7, 1, '加密服务', '172.20.4.110', 7009, '金融级加密模块','mis_21'),

-- 项目8 (ID=8 游戏后台)
(8, 1, '角色服务', '10.30.5.101', 8000, '玩家角色数据管理','mis_21'),
(8, 1, '战斗服务', '10.30.5.102', 8001, '实时战斗逻辑处理','mis_21'),
(8, 1, '匹配服务', '10.30.5.103', 8002, '游戏对战匹配系统','mis_21'),
(8, 1, '物品服务', '10.30.5.104', 8003, '虚拟物品管理系统','mis_21'),
(8, 1, '社交服务', '10.30.5.105', 8004, '游戏内社交系统','mis_21'),
(8, 1, '成就服务', '10.30.5.106', 8005, '成就系统管理','mis_21'),
(8, 1, '排行榜服务', '10.30.5.107', 8006, '实时排名计算','mis_21'),
(8, 1, '日志服务', '10.30.5.108', 8007, '玩家行为日志','mis_21'),
(8, 1, '支付服务', '10.30.5.109', 8008, '游戏内购支付','mis_21'),
(8, 1, '防作弊服务', '10.30.5.110', 8009, '反外挂检测系统','mis_21'),

-- 项目9 (ID=9 内容管理)
(9, 1, '文章服务', '192.168.6.101', 9000, '多格式内容编辑器','mis_21'),
(9, 1, '用户服务', '192.168.6.102', 9001, '多租户用户体系','mis_21'),
(9, 1, '审核服务', '192.168.6.103', 9002, '内容安全审核','mis_21'),
(9, 1, '发布服务', '192.168.6.104', 9003, '多渠道发布引擎','mis_21'),
(9, 1, '统计服务', '192.168.6.105', 9004, '访问数据分析','mis_21'),
(9, 1, '搜索服务', '192.168.6.106', 9005, '内容搜索引擎','mis_21'),
(9, 1, '版本服务', '192.168.6.107', 9006, '内容版本控制','mis_21'),
(9, 1, '模板服务', '192.168.6.108', 9007, '可视化模板管理','mis_21'),
(9, 1, '工作流服务', '192.168.6.109', 9008, '审批流程引擎','mis_21'),
(9, 1, '通知服务', '192.168.6.110', 9009, '消息通知中心','mis_21'),

-- 项目10 (ID=10 智能家居)
(10, 1, '设备服务', '10.100.7.101', 10000, '物联网设备管理','mis_21'),
(10, 1, '场景服务', '10.100.7.102', 10001, '智能场景模式管理','mis_21'),
(10, 1, '语音服务', '10.100.7.103', 10002, '语音指令解析','mis_21'),
(10, 1, '能源服务', '10.100.7.104', 10003, '能耗监测分析','mis_21'),
(10, 1, '安防服务', '10.100.7.105', 10004, '安全预警系统','mis_21'),
(10, 1, '环境服务', '10.100.7.106', 10005, '室内环境监测','mis_21'),
(10, 1, '联动服务', '10.100.7.107', 10006, '设备联动规则引擎','mis_21'),
(10, 1, '固件服务', '10.100.7.108', 10007, 'OTA固件升级','mis_21'),
(10, 1, '日志服务', '10.100.7.109', 10008, '操作日志记录','mis_21'),
(10, 1, 'API网关', '10.100.7.110', 10009, '统一接入网关','mis_21');


INSERT INTO profile (microservice_id, project_id, ptype, oss_path, comment) VALUES
(1, 1, 'cpu', 'oss://pprof/project1/ms1/profile_1.pprof', 'pprof profile test data 1'),
(1, 1, 'memory', 'oss://pprof/project1/ms1/profile_2.pprof', 'pprof profile test data 2'),
(1, 1, 'gorutine', 'oss://pprof/project1/ms1/profile_3.pprof', 'pprof profile test data 3'),
(1, 1, 'block', 'oss://pprof/project1/ms1/profile_4.pprof', 'pprof profile test data 4'),
(1, 1, 'mutex', 'oss://pprof/project1/ms1/profile_5.pprof', 'pprof profile test data 5'),
(1, 1, 'cpu', 'oss://pprof/project1/ms1/profile_6.pprof', 'pprof profile test data 6'),
(1, 1, 'memory', 'oss://pprof/project1/ms1/profile_7.pprof', 'pprof profile test data 7'),
(1, 1, 'gorutine', 'oss://pprof/project1/ms1/profile_8.pprof', 'pprof profile test data 8'),
(1, 1, 'block', 'oss://pprof/project1/ms1/profile_9.pprof', 'pprof profile test data 9'),
(1, 1, 'mutex', 'oss://pprof/project1/ms1/profile_10.pprof', 'pprof profile test data 10');