/* *****************************************************************************
// Setup the preferences
// ****************************************************************************/
SET NAMES utf8 COLLATE 'utf8_unicode_ci';
SET foreign_key_checks = 1;
SET time_zone = '+00:00';
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';
SET storage_engine = InnoDB;
SET CHARACTER SET utf8;

/* *****************************************************************************
// Remove old database
// ****************************************************************************/
DROP DATABASE IF EXISTS achievers;

/* *****************************************************************************
// Create new database
// ****************************************************************************/
CREATE DATABASE achievers DEFAULT CHARSET = utf8 COLLATE = utf8_unicode_ci;
USE achievers;

/* *****************************************************************************
// Create the tables
// ****************************************************************************/

CREATE TABLE user_status (
    id TINYINT(1) UNSIGNED NOT NULL AUTO_INCREMENT,
    
    status VARCHAR(25) NOT NULL,
    
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted TINYINT(1) UNSIGNED NOT NULL DEFAULT 0,
    
    PRIMARY KEY (id)
);

CREATE TABLE involvement (
    id TINYINT(1) UNSIGNED NOT NULL AUTO_INCREMENT,
    
    name VARCHAR(25) NOT NULL,
    
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT 0,
    
    PRIMARY KEY (id)
);

CREATE TABLE multimedia_type (
    id TINYINT(1) UNSIGNED NOT NULL AUTO_INCREMENT,
    
    name VARCHAR(25) NOT NULL,
    
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT 0,
    
    PRIMARY KEY (id)
);

CREATE TABLE user (
    id VARCHAR(36) NOT NULL,
    
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password CHAR(60) NOT NULL,
    
    status_id TINYINT(1) UNSIGNED NOT NULL DEFAULT 1,
    
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT 0,
    
    UNIQUE KEY (email),
    CONSTRAINT `f_user_status` FOREIGN KEY (`status_id`)
        REFERENCES `user_status` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
    
    PRIMARY KEY (id)
);

CREATE TABLE achievement (
    id VARCHAR(36) NOT NULL,
    
    title VARCHAR(50) NOT NULL,
    description VARCHAR(255) NOT NULL,
    picture_url VARCHAR(100) NOT NULL,
    
    involvement_id TINYINT(1) UNSIGNED NOT NULL,
    author_id      VARCHAR(36) NOT NULL,
    
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT 0,
    
    CONSTRAINT `f_achievement_involvement` FOREIGN KEY (`involvement_id`) 
        REFERENCES `involvement` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,

    CONSTRAINT `f_achievement_user` FOREIGN KEY (`author_id`) 
        REFERENCES `user` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
    
    PRIMARY KEY (id)
);

CREATE TABLE evidence (
    id VARCHAR(36) NOT NULL,
    
    description VARCHAR(255) NOT NULL,
    preview_url VARCHAR(255) NOT NULL,
    url         VARCHAR(255) NOT NULL,
    
    multimedia_type_id TINYINT(1) UNSIGNED NOT NULL,
    achievement_id     VARCHAR(36) NOT NULL,
    author_id          VARCHAR(36) NOT NULL,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT 0,
    
    CONSTRAINT `f_evidence_multimedia_type_id` FOREIGN KEY (`multimedia_type_id`)
        REFERENCES `multimedia_type` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,

    CONSTRAINT `f_evidence_achievement` FOREIGN KEY (`achievement_id`) 
        REFERENCES `achievement` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,

    CONSTRAINT `f_evidence_user` FOREIGN KEY (`author_id`) 
        REFERENCES `user` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
    
    PRIMARY KEY (id)
);

INSERT INTO `user_status` (`id`, `status`, `created_at`, `updated_at`, `deleted`) VALUES
(1, 'active',   CURRENT_TIMESTAMP,  CURRENT_TIMESTAMP,  0),
(2, 'inactive', CURRENT_TIMESTAMP,  CURRENT_TIMESTAMP,  0);

INSERT INTO `involvement` (`id`, `name`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'bronze',   CURRENT_TIMESTAMP,  CURRENT_TIMESTAMP,  0),
(2, 'silver', CURRENT_TIMESTAMP,  CURRENT_TIMESTAMP,  0),
(3, 'gold', CURRENT_TIMESTAMP,  CURRENT_TIMESTAMP,  0),
(4, 'platinum', CURRENT_TIMESTAMP,  CURRENT_TIMESTAMP,  0),
(5, 'diamond', CURRENT_TIMESTAMP,  CURRENT_TIMESTAMP,  0);

INSERT INTO `multimedia_type` (`id`, `name`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'photo',   CURRENT_TIMESTAMP,  CURRENT_TIMESTAMP,  0),
(2, 'video', CURRENT_TIMESTAMP,  CURRENT_TIMESTAMP,  0),
(3, 'voice', CURRENT_TIMESTAMP,  CURRENT_TIMESTAMP,  0);