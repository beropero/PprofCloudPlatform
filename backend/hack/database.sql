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
    FOREIGN KEY (creator_id) REFERENCES users(user_id) ON DELETE CASCADE
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
    FOREIGN KEY (project_id) REFERENCES projects(project_id) ON DELETE CASCADE
);

-- 测试记录
-- Create profile table
CREATE TABLE profile (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    microservice_id INT,
    project_id INT,
    default_sample_type VARCHAR(255),
    doc_url TEXT,
    drop_frames VARCHAR(255),
    keep_frames VARCHAR(255),
    time_nanos BIGINT,
    duration_nanos BIGINT,
    period BIGINT,
    period_type_type VARCHAR(255),
    period_type_unit VARCHAR(255),

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (microservice_id) REFERENCES microservices(microservice_id) ON DELETE CASCADE
);

-- Create sample_type table
CREATE TABLE sample_type (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    profile_id BIGINT UNSIGNED NOT NULL,
    type VARCHAR(255) NOT NULL,
    unit VARCHAR(255) NOT NULL,
    FOREIGN KEY (profile_id) REFERENCES profile(id) ON DELETE CASCADE
);

-- Create sample table
CREATE TABLE sample (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    profile_id BIGINT UNSIGNED NOT NULL,
    sample_type_value varchar(255),
    sample_type_unit varchar(255),
    value VARCHAR(255) NOT NULL,
    FOREIGN KEY (profile_id) REFERENCES profile(id) ON DELETE CASCADE
);

-- Create mapping table
CREATE TABLE mapping (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    profile_id BIGINT UNSIGNED NOT NULL,
    start BIGINT UNSIGNED NOT NULL,
    `limit` BIGINT UNSIGNED NOT NULL,
    offset BIGINT UNSIGNED NOT NULL,
    file VARCHAR(255) NOT NULL,
    build_id VARCHAR(255),
    has_functions BOOLEAN NOT NULL,
    has_filenames BOOLEAN NOT NULL,
    has_line_numbers BOOLEAN NOT NULL,
    has_inline_frames BOOLEAN NOT NULL,
    FOREIGN KEY (profile_id) REFERENCES profile(id) ON DELETE CASCADE
);

-- Create location table
CREATE TABLE location (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    profile_id BIGINT UNSIGNED NOT NULL,
    mapping_id BIGINT UNSIGNED,
    address BIGINT UNSIGNED NOT NULL,
    is_folded BOOLEAN NOT NULL,
    FOREIGN KEY (profile_id) REFERENCES profile(id) ON DELETE CASCADE,
    FOREIGN KEY (mapping_id) REFERENCES mapping(id) ON DELETE SET NULL
);

-- Create sample_Value table for storing sample values
CREATE TABLE sample_Value (
    sample_id BIGINT UNSIGNED NOT NULL,
    value_index INT UNSIGNED NOT NULL,
    value BIGINT NOT NULL,
    PRIMARY KEY (sample_id, value_index),
    FOREIGN KEY (sample_id) REFERENCES sample(id) ON DELETE CASCADE
);

-- Create sample_Location junction table
CREATE TABLE sample_Location (
    sample_id BIGINT UNSIGNED NOT NULL,
    location_id BIGINT UNSIGNED NOT NULL,
    order_index INT UNSIGNED NOT NULL,
    PRIMARY KEY (sample_id, order_index),
    FOREIGN KEY (sample_id) REFERENCES sample(id) ON DELETE CASCADE,
    FOREIGN KEY (location_id) REFERENCES location(id) ON DELETE CASCADE
);

-- Create sample_Label table for string labels
CREATE TABLE sample_Label (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    sample_id BIGINT UNSIGNED NOT NULL,
    key_str VARCHAR(255) NOT NULL,
    value_str VARCHAR(255) NOT NULL,
    FOREIGN KEY (sample_id) REFERENCES sample(id) ON DELETE CASCADE
);

-- Create sample_numLabel table for numeric labels
CREATE TABLE sample_numLabel (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    sample_id BIGINT UNSIGNED NOT NULL,
    key_str VARCHAR(255) NOT NULL,
    value_index INT UNSIGNED NOT NULL,
    num_value BIGINT NOT NULL,
    unit_str VARCHAR(255),
    FOREIGN KEY (sample_id) REFERENCES sample(id) ON DELETE CASCADE
);

-- Create function table
CREATE TABLE `function` (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    profile_id BIGINT UNSIGNED NOT NULL,
    name VARCHAR(255) NOT NULL,
    system_name VARCHAR(255) NOT NULL,
    filename VARCHAR(255) NOT NULL,
    start_line BIGINT NOT NULL,
    FOREIGN KEY (profile_id) REFERENCES profile(id) ON DELETE CASCADE
);

-- Create line_entry table for location lines
CREATE TABLE line_entry (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    location_id BIGINT UNSIGNED NOT NULL,
    function_id BIGINT UNSIGNED,
    line BIGINT NOT NULL,
    `column` BIGINT NOT NULL,
    FOREIGN KEY (location_id) REFERENCES location(id) ON DELETE CASCADE,
    FOREIGN KEY (function_id) REFERENCES `function`(id) ON DELETE SET NULL
);

-- Create profile_comment table for comments
CREATE TABLE profile_comment (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    profile_id BIGINT UNSIGNED NOT NULL,
    comment TEXT NOT NULL,
    order_index INT UNSIGNED NOT NULL,
    FOREIGN KEY (profile_id) REFERENCES profile(id) ON DELETE CASCADE
);

-- Indexes for foreign keys to improve query performance
CREATE INDEX idx_sample_type_profile_id ON sample_type (profile_id);
CREATE INDEX idx_sample_profile_id ON sample (profile_id);
CREATE INDEX idx_sample_location_sample_id ON sample_Location (sample_id);
CREATE INDEX idx_sample_location_location_id ON sample_Location (location_id);
CREATE INDEX idx_sample_label_sample_id ON sample_Label (sample_id);
CREATE INDEX idx_sample_numlabel_sample_id ON sample_numLabel (sample_id);
CREATE INDEX idx_mapping_profile_id ON mapping (profile_id);
CREATE INDEX idx_location_profile_id ON location (profile_id);
CREATE INDEX idx_location_mapping_id ON location (mapping_id);
CREATE INDEX idx_lineentry_location_id ON line_entry (location_id);
CREATE INDEX idx_lineentry_function_id ON line_entry (function_id);
CREATE INDEX idx_function_profile_id ON `function` (profile_id);
CREATE INDEX idx_profile_comment_profile_id ON profile_comment (profile_id);
