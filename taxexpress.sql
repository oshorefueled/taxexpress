-- MySQL Script generated by MySQL Workbench
-- Tue 26 Feb 2019 17:45:50 WAT
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL,ALLOW_INVALID_DATES';

-- -----------------------------------------------------
-- Schema taxexpress
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema taxexpress
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `taxexpress` DEFAULT CHARACTER SET utf8 ;
USE `taxexpress` ;

-- -----------------------------------------------------
-- Table `taxexpress`.`admins`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `taxexpress`.`admins` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` VARCHAR(64) NULL,
  `email` VARCHAR(45) NULL,
  `token` VARCHAR(255) NULL,
  `password` VARCHAR(255) NOT NULL,
  `created_at` TIMESTAMP NULL,
  `updated_at` TIMESTAMP NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC),
  UNIQUE INDEX `token_UNIQUE` (`token` ASC),
  UNIQUE INDEX `email_UNIQUE` (`email` ASC))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `taxexpress`.`businesses`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `taxexpress`.`businesses` (
  `id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NULL,
  `rc_number` INT(8) NULL,
  `business_desc` VARCHAR(255) NULL,
  `email` VARCHAR(45) NOT NULL,
  `total_revenue` FLOAT NULL,
  `tax_status` INT NOT NULL,
  `created_at` TIMESTAMP NULL,
  `updated_at` TIMESTAMP NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC),
  UNIQUE INDEX `email_UNIQUE` (`email` ASC))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `taxexpress`.`tax`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `taxexpress`.`tax` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `business_id` INT(11) UNSIGNED NOT NULL,
  `tax_period` DATE NOT NULL,
  `revenue` FLOAT NOT NULL,
  `tax_paid` FLOAT NOT NULL DEFAULT 0.00,
  `date_paid` DATETIME NULL,
  `created_at` TIMESTAMP NULL,
  `updated_at` TIMESTAMP NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC),
  INDEX `fk_tax_1_idx` (`business_id` ASC),
  CONSTRAINT `fk_tax_1`
    FOREIGN KEY (`business_id`)
    REFERENCES `taxexpress`.`businesses` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `taxexpress`.`notifications`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `taxexpress`.`notifications` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `business_id` INT(11) UNSIGNED NOT NULL,
  `message` VARCHAR(255) NOT NULL,
  `to_business` TINYINT(1) NOT NULL,
  `to_authority` TINYINT(1) NOT NULL,
  `created_at` TIMESTAMP NULL,
  `updated_at` TIMESTAMP NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC),
  INDEX `fk_notifications_1_idx` (`business_id` ASC),
  CONSTRAINT `fk_notifications_1`
    FOREIGN KEY (`business_id`)
    REFERENCES `taxexpress`.`businesses` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `taxexpress`.`tax_status`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `taxexpress`.`tax_status` (
  `id` INT NOT NULL,
  `status` VARCHAR(11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `status_UNIQUE` (`status` ASC))
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;

-- -----------------------------------------------------
-- Data for table `taxexpress`.`tax_status`
-- -----------------------------------------------------
START TRANSACTION;
USE `taxexpress`;
INSERT INTO `taxexpress`.`tax_status` (`id`, `status`) VALUES (1, 'paid');
INSERT INTO `taxexpress`.`tax_status` (`id`, `status`) VALUES (2, 'pending');

COMMIT;

