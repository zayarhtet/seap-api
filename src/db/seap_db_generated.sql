-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema mydb
-- -----------------------------------------------------
-- -----------------------------------------------------
-- Schema seap_db
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema seap_db
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `seap_db_test` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci ;
USE `seap_db_test` ;

-- -----------------------------------------------------
-- Table `seap_db`.`credential`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `seap_db`.`credential` (
  `credential_id` VARCHAR(255) NOT NULL,
  `password` VARCHAR(255) NOT NULL,
  `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `modified_at` TIMESTAMP NULL DEFAULT NULL,
  UNIQUE INDEX `credential_id` (`credential_id` ASC))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `seap_db`.`family`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `seap_db`.`family` (
  `family_id` VARCHAR(255) NOT NULL,
  `family_name` VARCHAR(255) NOT NULL,
  `family_info` VARCHAR(1000) NOT NULL,
  `family_icon` VARCHAR(300) NULL DEFAULT NULL,
  `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `modified_at` TIMESTAMP NULL DEFAULT NULL,
  PRIMARY KEY (`family_id`),
  UNIQUE INDEX `family_id` (`family_id` ASC) VISIBLE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `seap_db`.`duty`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `seap_db`.`duty` (
  `duty_id` VARCHAR(255) NOT NULL,
  `title` VARCHAR(255) NOT NULL,
  `instruction` VARCHAR(1000) NULL DEFAULT NULL,
  `publishing_date` TIMESTAMP NOT NULL,
  `deadline_date` TIMESTAMP NOT NULL,
  `closing_date` TIMESTAMP NOT NULL,
  `family_id` VARCHAR(255) NOT NULL,
  `point_system` TINYINT(1) NULL DEFAULT '1',
  `possible_points` DOUBLE NULL DEFAULT NULL,
  `multiple_submission` TINYINT(1) NULL DEFAULT '1',
  `plugin_name` VARCHAR(255) NULL DEFAULT NULL,
  PRIMARY KEY (`duty_id`),
  INDEX `FK_duty_family` (`family_id` ASC) VISIBLE,
  CONSTRAINT `FK_duty_family`
    FOREIGN KEY (`family_id`)
    REFERENCES `seap_db`.`family` (`family_id`)
    ON DELETE CASCADE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `seap_db`.`role`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `seap_db`.`role` (
  `role_id` INT NOT NULL,
  `name` VARCHAR(20) NOT NULL,
  `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `modified_at` TIMESTAMP NULL DEFAULT NULL,
  PRIMARY KEY (`role_id`),
  UNIQUE INDEX `role_id` (`role_id` ASC) VISIBLE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `seap_db`.`member`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `seap_db`.`member` (
  `first_name` VARCHAR(255) NULL DEFAULT NULL,
  `last_name` VARCHAR(255) NULL DEFAULT NULL,
  `username` VARCHAR(20) NOT NULL,
  `email` VARCHAR(50) NOT NULL,
  `credential_id` VARCHAR(255) NOT NULL,
  `role_id` INT NOT NULL,
  `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `modified_at` TIMESTAMP NULL DEFAULT NULL,
  PRIMARY KEY (`username`),
  UNIQUE INDEX `username` (`username` ASC) VISIBLE,
  UNIQUE INDEX `email` (`email` ASC) VISIBLE,
  UNIQUE INDEX `credential_id` (`credential_id` ASC) VISIBLE,
  INDEX `FK_memeber_role` (`role_id` ASC) VISIBLE,
  CONSTRAINT `fk_member_credential`
    FOREIGN KEY (`credential_id`)
    REFERENCES `seap_db`.`credential` (`credential_id`),
  CONSTRAINT `FK_memeber_role`
    FOREIGN KEY (`role_id`)
    REFERENCES `seap_db`.`role` (`role_id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `seap_db`.`family_member`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `seap_db`.`family_member` (
  `username` VARCHAR(255) NOT NULL,
  `family_id` VARCHAR(255) NOT NULL,
  `role_id` INT NOT NULL,
  `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `modified_at` TIMESTAMP NULL DEFAULT NULL,
  PRIMARY KEY (`username`, `family_id`),
  INDEX `FK_familymember_role` (`role_id` ASC) VISIBLE,
  INDEX `FK_familymember_family` (`family_id` ASC) VISIBLE,
  CONSTRAINT `FK_familymember_family`
    FOREIGN KEY (`family_id`)
    REFERENCES `seap_db`.`family` (`family_id`)
    ON DELETE CASCADE,
  CONSTRAINT `FK_familymember_member`
    FOREIGN KEY (`username`)
    REFERENCES `seap_db`.`member` (`username`)
    ON DELETE CASCADE,
  CONSTRAINT `FK_familymember_role`
    FOREIGN KEY (`role_id`)
    REFERENCES `seap_db`.`role` (`role_id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `seap_db`.`given_file`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `seap_db`.`given_file` (
  `file_id` VARCHAR(255) NOT NULL,
  `duty_id` VARCHAR(255) NOT NULL,
  `file_path` VARCHAR(1000) NOT NULL,
  PRIMARY KEY (`file_id`),
  INDEX `given_file_duty` (`duty_id` ASC) VISIBLE,
  CONSTRAINT `given_file_duty`
    FOREIGN KEY (`duty_id`)
    REFERENCES `seap_db`.`duty` (`duty_id`)
    ON DELETE CASCADE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `seap_db`.`grading`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `seap_db`.`grading` (
  `username` VARCHAR(255) NOT NULL,
  `duty_id` VARCHAR(255) NOT NULL,
  `family_id` VARCHAR(255) NOT NULL,
  `submitted` TINYINT(1) NULL DEFAULT '0',
  `points` DOUBLE NULL DEFAULT NULL,
  `is_late` TINYINT(1) NULL DEFAULT '0',
  `is_passed` TINYINT(1) NULL DEFAULT '0',
  `grade_comment` VARCHAR(1000) NULL DEFAULT NULL,
  `execution_comment` VARCHAR(2000) NULL DEFAULT NULL,
  `submitted_at` TIMESTAMP NULL DEFAULT NULL,
  `grading_id` VARCHAR(255) NOT NULL,
  `has_graded` TINYINT(1) NULL DEFAULT '0',
  PRIMARY KEY (`grading_id`),
  INDEX `FK_grading_member` (`username` ASC) VISIBLE,
  INDEX `FK_grading_family` (`family_id` ASC) VISIBLE,
  INDEX `FK_grading_duty` (`duty_id` ASC) VISIBLE,
  CONSTRAINT `FK_grading_duty`
    FOREIGN KEY (`duty_id`)
    REFERENCES `seap_db`.`duty` (`duty_id`)
    ON DELETE CASCADE,
  CONSTRAINT `FK_grading_family`
    FOREIGN KEY (`family_id`)
    REFERENCES `seap_db`.`family` (`family_id`)
    ON DELETE CASCADE,
  CONSTRAINT `FK_grading_member`
    FOREIGN KEY (`username`)
    REFERENCES `seap_db`.`member` (`username`)
    ON DELETE CASCADE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `seap_db`.`submitted_file`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `seap_db`.`submitted_file` (
  `file_id` VARCHAR(255) NOT NULL,
  `grading_id` VARCHAR(255) NOT NULL,
  `file_path` VARCHAR(1000) NOT NULL,
  `submitted_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`file_id`),
  INDEX `submitted_file_grading` (`grading_id` ASC) VISIBLE,
  CONSTRAINT `submitted_file_grading`
    FOREIGN KEY (`grading_id`)
    REFERENCES `seap_db`.`grading` (`grading_id`)
    ON DELETE CASCADE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;


insert into role(role_id,name) values(99, "admin");
insert into role(role_id, name) values(1, "tutor");
insert into role(role_id, name) values(2, "tutee");

insert into credential (credential_id, password) values ("e0f5a784-aa4b-4523-a3bd-b4c01a6ca7e6", "$2a$10$2y5U2e1P7pr8CrVjjMNOtOmcrB25e4JLZMDf.S5KwzFVbcFs8Wt3G");
insert into member (first_name, last_name, username, email, credential_id, role_id) values ("Admin", "Admin", "admin", "admin@seap.org", "e0f5a784-aa4b-4523-a3bd-b4c01a6ca7e6", 99);

insert into credential (credential_id, password) values ("91164b58-2eba-4a81-ba0c-dae18be822e0", "$2a$10$omsB.jQKURmZ8g3v7imIWOGhaOjiW4tt.9j2oSJTL7xtV9jEBiz9.");
insert into member (first_name, last_name, username, email, credential_id, role_id) values ("John", "Doe", "johndoe", "johndoe@seap.org", "f72d1529-327c-4f2d-8fe1-6bf777b7fd25", 1);

insert into credential (credential_id, password) values ("f72d1529-327c-4f2d-8fe1-6bf777b7fd25", "$2a$10$7L5D5yIUOkpbWyi90NL6rePzMhU3YhS.s5MfYJje1NqLokIy7PxVS");
insert into member (first_name, last_name, username, email, credential_id, role_id) values ("Jane", "Doe", "janedoe", "janedoe@seap.org", "91164b58-2eba-4a81-ba0c-dae18be822e0", 2);
