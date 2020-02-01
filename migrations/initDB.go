package main

import h "alisa-millionaire-lite/server/helpers"

/*
CREATE SCHEMA `db_aml` DEFAULT CHARACTER SET utf8;
*/

/*
CREATE TABLE `db_aml`.`users (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `login` tinytext,
  `password` tinytext,
  `reg_date` tinytext,
  `rate` int unsigned DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8
*/

/*
CREATE TABLE `db_aml`.`tasks` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` tinytext,
  `valid_variant` tinytext,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8
*/

/*
CREATE TABLE `db_aml`.`answers` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `game_id` tinytext NOT NULL,
  `task_id` int unsigned NOT NULL,
  `variant_id` int unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8
*/

/*
CREATE TABLE `variants` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `task_id` int DEFAULT NULL,
  `content` tinytext,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  KEY `id_idx` (`task_id`)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8
*/

func main() {
	db := h.ConnMySQL()
	db.Close()
}
