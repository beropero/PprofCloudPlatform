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
CREATE TABLE projects (
    project_id INT PRIMARY KEY AUTO_INCREMENT,
    creator_id INT,
    project_name VARCHAR(255) NOT NULL,
    project_description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (creator_id) REFERENCES users(user_id)
);

-- 微服务表
CREATE TABLE microservices (
    microservice_id INT PRIMARY KEY AUTO_INCREMENT,
    project_id INT,
    microservice_name VARCHAR(255) NOT NULL,
    ip VARCHAR(15) NOT NULL,
    port INT NOT NULL,
    microservice_description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (project_id) REFERENCES projects(project_id)
);

-- 测试记录表
CREATE TABLE test_records (
    record_id INT PRIMARY KEY AUTO_INCREMENT,
    cpu_usage FLOAT,
    memory_usage FLOAT,
    disk_usage FLOAT,
    network_usage FLOAT,
    microservice_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (microservice_id) REFERENCES microservices(microservice_id)
);