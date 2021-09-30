CREATE TABLE `bms`.`bms_user` (
  `id` INT NOT NULL AUTO_INCREMENT ,
  `name` VARCHAR(64) NULL,
  `authId` VARCHAR(256) NOT NULL,
  `mobile` VARCHAR(13) NULL,
  `email` VARCHAR(64) NOT NULL,
  `password` varchar(256) NOT NULL,
  `gender` VARCHAR(1) NULL,
  `createdOn` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
  `updatedOn` DATETIME NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`));

CREATE TABLE `bms`.`bms_movie` (
`id` INT NOT NULL AUTO_INCREMENT,
`name` VARCHAR(64) NULL,
`description` VARCHAR(256) NULL,
`ratings` VARCHAR(10) NULL,
`language` VARCHAR(45) NULL,
`photo` VARCHAR(256) NULL,
`genere` VARCHAR(45) NULL,
`createdOn` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
`updatedOn` DATETIME NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`id`));

CREATE TABLE `bms`.`bms_theater` (
`id` INT NOT NULL AUTO_INCREMENT,
`name` VARCHAR(45) NULL,
`address` VARCHAR(256) NULL,
`totalCapacity` INT(20) NULL,
`totalShows` INT(10) NULL,
`createdOn` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
`updatedOn` DATETIME NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`id`));

CREATE TABLE `bms`.`bms_shows` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `movieId` INT NOT NULL,
  `theaterId` INT NOT NULL,
  `showDate` DATE NULL,
  `showPattern` VARCHAR(64) NULL,
  `createdOn` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
  `updatedOn` DATETIME NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `fk_bms_shows_1_idx` (`movieId` ASC) VISIBLE,
  INDEX `fk_bms_shows_2_idx` (`theaterId` ASC) VISIBLE,
  CONSTRAINT `fk_bms_shows_1`
    FOREIGN KEY (`movieId`)
    REFERENCES `bms`.`bms_movie` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_bms_shows_2`
    FOREIGN KEY (`theaterId`)
    REFERENCES `bms`.`bms_theater` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION);

CREATE TABLE `bms`.`bms_bookings` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `userId` INT NOT NULL,
  `movieId` INT NOT NULL,
  `theaterId` INT NOT NULL,
  `showDate` DATE NULL,
  `showTime` VARCHAR(12) NULL,
  `totalSeats` INT NULL,
  `createdOn` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
  `updatedOn` DATETIME NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `fk_bms_bookings_1_idx` (`userId` ASC) VISIBLE,
  INDEX `fk_bms_bookings_2_idx` (`movieId` ASC) VISIBLE,
  INDEX `fk_bms_bookings_3_idx` (`theaterId` ASC) VISIBLE,
  CONSTRAINT `fk_bms_bookings_1`
    FOREIGN KEY (`userId`)
    REFERENCES `bms`.`bms_user` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_bms_bookings_2`
    FOREIGN KEY (`movieId`)
    REFERENCES `bms`.`bms_movie` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_bms_bookings_3`
    FOREIGN KEY (`theaterId`)
    REFERENCES `bms`.`bms_theater` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION);
