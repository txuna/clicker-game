-- Database export via SQLPro (https://www.sqlprostudio.com/)
-- Exported by tuuna at 15-06-2024 15:50.
-- WARNING: This file may contain descructive statements such as DROPs.
-- Please ensure that you are running the script at the proper location.

CREATE DATABASE IF NOT EXISTS `clicker-game` charset=utf8mb3;
USE `clicker-game`;

-- BEGIN TABLE accounts
DROP TABLE IF EXISTS accounts;
CREATE TABLE `accounts` (
  `id` int NOT NULL AUTO_INCREMENT,
  `userid` varchar(45) NOT NULL,
  `userpw` varchar(128) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb3;

-- Inserting 3 rows into accounts
-- Insert batch #1
INSERT INTO accounts (id, userid, userpw) VALUES
(2, 'helloworld', '$2a$04$Pww3FaZ9OfriTRTfYoz/GumRzrCNme.X5O5pcFLgB3VQemZY26lWC'),
(3, 'tuuna', '$2a$04$48R66Q.FffyV.GaHDPrLXeV64FRejEb9w5Kro4SVyGRioEKRjAS2K'),
(4, 'cloud', '$2a$04$4uDnj5iNgGXZcCaNhvcvw.wDUTSGWLE067bKYd2l2Fu.U7lYfuxZC');

-- END TABLE accounts

-- BEGIN TABLE players
DROP TABLE IF EXISTS players;
CREATE TABLE `players` (
  `player_id` int NOT NULL,
  `coin` int NOT NULL,
  `max_coin` int NOT NULL,
  PRIMARY KEY (`player_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

-- Inserting 3 rows into players
-- Insert batch #1
INSERT INTO players (player_id, coin, max_coin) VALUES
(2, 47, 5000),
(3, 352, 5000),
(4, 211, 5000);

-- END TABLE players

