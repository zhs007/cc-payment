# ************************************************************
# Sequel Pro SQL dump
# Version 4096
#
# http://www.sequelpro.com/
# http://code.google.com/p/sequel-pro/
#
# Host: 47.90.46.159 (MySQL 5.7.23)
# Database: ccpayment
# Generation Time: 2019-02-17 06:53:15 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table users
# ------------------------------------------------------------

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `userid` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'User ID',
  `username` varchar(32) NOT NULL DEFAULT '' COMMENT 'User name',
  `status` int(11) NOT NULL DEFAULT '0' COMMENT 'Status',
  `registertime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'User registered timestamp',
  PRIMARY KEY (`userid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table usercurrencies
# ------------------------------------------------------------

DROP TABLE IF EXISTS `usercurrencies`;

CREATE TABLE `usercurrencies` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'Unique identifier',
  `userid` bigint(20) NOT NULL COMMENT 'User ID',
  `currency` varchar(16) NOT NULL DEFAULT '' COMMENT 'Currency',
  `balance` bigint(20) NOT NULL COMMENT 'Balance',
  PRIMARY KEY (`id`),
  UNIQUE KEY `usercurrency` (`userid`,`currency`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table userpayments
# ------------------------------------------------------------

DROP TABLE IF EXISTS `userpayments`;

CREATE TABLE `userpayments` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'Unique identifier of payment',
  `fromuserid` bigint(20) NOT NULL COMMENT 'User ID of initiate payment',
  `touserid` bigint(20) NOT NULL COMMENT 'User ID of the payment received',
  `currency` varchar(16) NOT NULL DEFAULT '' COMMENT 'Currency',
  `starttime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Timestamp of initiating payment',
  `donetime` timestamp NULL DEFAULT NULL COMMENT 'Timestamp of payment completion',
  `paymentstatus` int(11) NOT NULL COMMENT 'Payment status',
  `startbalance0` bigint(11) NOT NULL COMMENT 'The pre-payment balance of the user who initiated the payment',
  `endbalance0` bigint(20) NOT NULL COMMENT 'The end-payment balance of the user who initiated the payment',
  `startbalance1` bigint(20) NOT NULL COMMENT 'The pre-payment balance of the user who payment received',
  `endbalance1` bigint(20) NOT NULL COMMENT 'The end-payment balance of the user who payment received',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
