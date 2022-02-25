#-- CREATE DATABASE  IF NOT EXISTS `mcp_v2` /*!40100 DEFAULT CHARACTER SET latin1 */;
#-- USE `mcp_v2`;
-- MySQL dump 10.13  Distrib 8.0.16, for macos10.14 (x86_64)
--
-- Host: 192.168.88.188    Database: mcp_v2
-- ------------------------------------------------------
-- Server version	5.7.32

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
 SET NAMES utf8 ;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `block_scan_record`
--

DROP TABLE IF EXISTS `block_scan_record`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `block_scan_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `network_id` bigint(20) DEFAULT NULL,
  `last_current_block_number` bigint(20) NOT NULL,
  `update_at` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `number_UNIQUE` (`last_current_block_number`),
  KEY `block_scan_record_network_id_fk` (`network_id`),
  CONSTRAINT `block_scan_record_network_id_fk` FOREIGN KEY (`network_id`) REFERENCES `network` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `coin`
--

DROP TABLE IF EXISTS `coin`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `coin` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `short_name` varchar(255) COLLATE utf8_bin NOT NULL,
  `full_name` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `cn_name` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `coin_address` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `uuid` varchar(128) COLLATE utf8_bin DEFAULT NULL,
  `network_id` bigint(20) DEFAULT NULL,
  `gas_price` int(11) DEFAULT '0',
  `gas_limit` int(11) DEFAULT '0',
  `description` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `coin_uuid_uindex` (`uuid`),
  KEY `coin_network_id_fk` (`network_id`),
  CONSTRAINT `coin_network_id_fk` FOREIGN KEY (`network_id`) REFERENCES `network` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `dao_fetched_deal`
--

DROP TABLE IF EXISTS `dao_fetched_deal`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `dao_fetched_deal` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `deal_id` bigint(20) NOT NULL,
  `create_at` bigint(20) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=762 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `dao_info`
--

DROP TABLE IF EXISTS `dao_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `dao_info` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `dao_name` varchar(255) COLLATE utf8_bin NOT NULL,
  `dao_address` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `order_index` int(11) NOT NULL,
  `description` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `create_at` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `deal_file`
--

DROP TABLE IF EXISTS `deal_file`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `deal_file` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `deal_cid` varchar(255) COLLATE utf8_bin DEFAULT '',
  `car_file_name` varchar(255) CHARACTER SET utf8 DEFAULT NULL,
  `payload_cid` varchar(255) CHARACTER SET utf8 DEFAULT NULL,
  `deal_id` bigint(20) DEFAULT NULL,
  `piece_cid` varchar(255) CHARACTER SET utf8 DEFAULT NULL,
  `car_file_size` bigint(20) DEFAULT NULL,
  `miner_fid` varchar(128) COLLATE utf8_bin DEFAULT NULL,
  `source_file_path` varchar(255) CHARACTER SET utf8 DEFAULT NULL,
  `car_file_path` varchar(255) CHARACTER SET utf8 DEFAULT NULL,
  `car_md5` varchar(45) CHARACTER SET utf8 DEFAULT NULL,
  `deal_status` varchar(64) CHARACTER SET utf8 DEFAULT NULL,
  `pin_status` varchar(64) COLLATE utf8_bin DEFAULT NULL,
  `duration` int(11) DEFAULT NULL,
  `task_uuid` varchar(128) COLLATE utf8_bin DEFAULT '',
  `cost` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `client_wallet_address` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `lock_payment_tx` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `lock_payment_status` varchar(32) COLLATE utf8_bin DEFAULT 'Pending',
  `lock_payment_network` bigint(20) DEFAULT NULL,
  `dao_sign_status` varchar(32) COLLATE utf8_bin DEFAULT NULL,
  `send_deal_status` varchar(32) COLLATE utf8_bin DEFAULT '',
  `verified` tinyint(1) DEFAULT '0',
  `create_at` bigint(20) DEFAULT NULL,
  `delete_at` bigint(20) DEFAULT NULL,
  `update_at` bigint(20) DEFAULT NULL,
  `is_deleted` tinyint(1) DEFAULT '0',
  `max_price` double DEFAULT NULL,
  `refund_status_after_unlock` varchar(45) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `deal_file_network_id_fk` (`lock_payment_network`)
) ENGINE=InnoDB AUTO_INCREMENT=1024 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `event_dao_signature`
--

DROP TABLE IF EXISTS `event_dao_signature`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `event_dao_signature` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `tx_hash` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `recipient` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `payload_cid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `order_id` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `deal_id` bigint(20) DEFAULT NULL,
  `dao_pass_time` varchar(64) COLLATE utf8_bin DEFAULT NULL,
  `block_no` bigint(20) DEFAULT NULL,
  `coin_id` bigint(20) DEFAULT NULL,
  `network_id` bigint(20) DEFAULT NULL,
  `dao_address` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `block_time` varchar(64) COLLATE utf8_bin DEFAULT NULL,
  `signature_unlock_status` varchar(8) COLLATE utf8_bin DEFAULT '0',
  `status` tinyint(1) DEFAULT NULL,
  `tx_hash_unlock` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `dao_event_log_coin_id_fk` (`coin_id`),
  KEY `dao_event_log_network_id_fk` (`network_id`),
  KEY `dao_event_log_order_id_index` (`order_id`),
  KEY `dao_event_log_payload_cid_index` (`payload_cid`),
  KEY `dao_event_log_tx_hash_index` (`tx_hash`),
  CONSTRAINT `dao_event_log_coin_id_fk` FOREIGN KEY (`coin_id`) REFERENCES `coin` (`id`),
  CONSTRAINT `dao_event_log_network_id_fk` FOREIGN KEY (`network_id`) REFERENCES `network` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=920 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `event_expire_payment`
--

DROP TABLE IF EXISTS `event_expire_payment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `event_expire_payment` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `tx_hash` varchar(255) DEFAULT NULL,
  `payload_cid` varchar(255) DEFAULT NULL,
  `block_no` varchar(128) DEFAULT NULL,
  `token_address` varchar(255) DEFAULT NULL,
  `contract_address` varchar(255) DEFAULT NULL,
  `user_address` varchar(255) DEFAULT NULL,
  `expire_user_amount` varchar(255) DEFAULT NULL,
  `block_time` varchar(64) DEFAULT NULL,
  `create_at` varchar(64) DEFAULT NULL,
  `network_id` bigint(20) DEFAULT NULL,
  `coin_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `event_expire_payment_coin_info_id_fk` (`coin_id`),
  KEY `event_expire_payment_network_info_id_fk` (`network_id`),
  CONSTRAINT `event_expire_payment_coin_info_id_fk` FOREIGN KEY (`coin_id`) REFERENCES `coin` (`id`),
  CONSTRAINT `event_expire_payment_network_info_id_fk` FOREIGN KEY (`network_id`) REFERENCES `network` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `event_lock_payment`
--

DROP TABLE IF EXISTS `event_lock_payment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `event_lock_payment` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `network_id` bigint(20) DEFAULT NULL,
  `tx_hash` varchar(255) COLLATE utf8_bin NOT NULL,
  `payload_cid` varchar(256) COLLATE utf8_bin DEFAULT NULL,
  `token_address` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `contract_address` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `locked_fee` double DEFAULT NULL,
  `min_payment` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `deadline` varchar(64) COLLATE utf8_bin DEFAULT NULL,
  `block_no` bigint(20) DEFAULT NULL,
  `address_from` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `address_to` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `create_at` bigint(20) DEFAULT NULL,
  `lock_payment_time` bigint(20) DEFAULT NULL,
  `coin_id` bigint(20) DEFAULT NULL,
  `vrf_rand` varchar(100) COLLATE utf8_bin NOT NULL DEFAULT '',
  `source_file_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `un_event_lock_payment` (`payload_cid`,`address_from`),
  KEY `event_lock_payment_coin_info_id_fk` (`coin_id`),
  KEY `event_polygon_network_info_id_fk` (`network_id`),
  KEY `fk_event_lock_payment_source_file_id` (`source_file_id`),
  CONSTRAINT `event_lock_payment_coin_info_id_fk` FOREIGN KEY (`coin_id`) REFERENCES `coin` (`id`),
  CONSTRAINT `event_lock_payment_ibfk_1` FOREIGN KEY (`source_file_id`) REFERENCES `source_file` (`id`),
  CONSTRAINT `event_polygon_network_info_id_fk` FOREIGN KEY (`network_id`) REFERENCES `network` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=89586 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `event_unlock_payment`
--

DROP TABLE IF EXISTS `event_unlock_payment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `event_unlock_payment` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `tx_hash` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `payload_cid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `block_no` varchar(128) COLLATE utf8_bin DEFAULT NULL,
  `token_address` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `unlock_from_address` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `unlock_to_user_address` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `unlock_to_user_amount` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `unlock_to_admin_address` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `unlock_to_admin_amount` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `unlock_time` bigint(20) DEFAULT NULL,
  `create_at` bigint(20) DEFAULT NULL,
  `network_id` bigint(20) DEFAULT NULL,
  `coin_id` bigint(20) DEFAULT NULL,
  `unlock_status` varchar(32) COLLATE utf8_bin DEFAULT NULL,
  `deal_id` bigint(20) DEFAULT NULL,
  `source_file_id` bigint(20) DEFAULT NULL,
  `locked_fee_before_unlock` decimal(20,0) DEFAULT NULL,
  `locked_fee_after_unlock` decimal(20,0) DEFAULT NULL,
  `update_at` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `event_unlock_payment_coin_info_id_fk` (`coin_id`),
  KEY `event_unlock_payment_network_info_id_fk` (`network_id`),
  KEY `fk_event_unlock_payment_source_file_id` (`source_file_id`),
  CONSTRAINT `event_unlock_payment_coin_info_id_fk` FOREIGN KEY (`coin_id`) REFERENCES `coin` (`id`),
  CONSTRAINT `event_unlock_payment_ibfk_1` FOREIGN KEY (`source_file_id`) REFERENCES `source_file` (`id`),
  CONSTRAINT `event_unlock_payment_network_info_id_fk` FOREIGN KEY (`network_id`) REFERENCES `network` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1751 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `mint_info`
--

DROP TABLE IF EXISTS `mint_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `mint_info` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `nft_tx_hash` varchar(255) COLLATE utf8_bin NOT NULL,
  `token_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `network`
--

DROP TABLE IF EXISTS `network`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `network` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `network_name` varchar(255) COLLATE utf8_bin NOT NULL,
  `uuid` varchar(128) COLLATE utf8_bin DEFAULT NULL,
  `rpc_url` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `native_coin` varchar(128) COLLATE utf8_bin DEFAULT NULL,
  `description` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `network_info_uuid_uindex` (`uuid`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `offline_deal`
--

DROP TABLE IF EXISTS `offline_deal`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `offline_deal` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `deal_file_id` bigint(20) NOT NULL,
  `deal_cid` varchar(100) NOT NULL,
  `miner_fid` varchar(45) NOT NULL,
  `start_epoch` int(11) NOT NULL,
  `sender_wallet` varchar(200) NOT NULL,
  `status` varchar(45) NOT NULL,
  `deal_id` bigint(20) NOT NULL,
  `create_at` bigint(20) NOT NULL,
  `update_at` bigint(20) NOT NULL,
  `unlock_status` varchar(45) NOT NULL,
  `note` text,
  `unlock_at` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `ind_offline_deal_deal_file_id` (`deal_file_id`),
  CONSTRAINT `fk_ofline_deal_deal_file_id` FOREIGN KEY (`deal_file_id`) REFERENCES `deal_file` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=139 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `source_file`
--

DROP TABLE IF EXISTS `source_file`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `source_file` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `resource_uri` varchar(255) CHARACTER SET utf8 DEFAULT NULL,
  `md5` varchar(45) CHARACTER SET utf8 DEFAULT NULL,
  `uuid` varchar(45) CHARACTER SET utf8 DEFAULT NULL,
  `status` varchar(45) CHARACTER SET utf8 DEFAULT NULL,
  `dataset` varchar(45) COLLATE utf8_bin DEFAULT NULL,
  `create_at` bigint(20) DEFAULT NULL,
  `file_size` bigint(20) DEFAULT NULL,
  `delete_at` varchar(128) COLLATE utf8_bin DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  `ipfs_url` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `pin_status` varchar(32) COLLATE utf8_bin DEFAULT NULL,
  `payload_cid` varchar(100) COLLATE utf8_bin NOT NULL DEFAULT '',
  `nft_tx_hash` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `token_id` varchar(45) COLLATE utf8_bin DEFAULT NULL,
  `mint_address` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `file_type` int(11) DEFAULT NULL,
  `update_at` bigint(20) DEFAULT NULL,
  `refund_status` varchar(60) COLLATE utf8_bin DEFAULT NULL,
  `refund_amount` decimal(20,0) DEFAULT NULL,
  `refund_at` bigint(20) DEFAULT NULL,
  `refund_tx_hash` varchar(100) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=708 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `source_file_deal_file_map`
--

DROP TABLE IF EXISTS `source_file_deal_file_map`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `source_file_deal_file_map` (
  `source_file_id` bigint(20) NOT NULL,
  `file_index` int(11) NOT NULL,
  `deal_file_id` bigint(20) NOT NULL,
  `create_at` bigint(20) DEFAULT NULL,
  `update_at` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`source_file_id`,`deal_file_id`),
  KEY `fk_source_file_has_deal_file_deal_file1_idx` (`deal_file_id`),
  KEY `fk_source_file_has_deal_file_source_file1_idx` (`source_file_id`),
  CONSTRAINT `fk_source_file_has_deal_file_deal_file1` FOREIGN KEY (`deal_file_id`) REFERENCES `deal_file` (`id`),
  CONSTRAINT `fk_source_file_has_deal_file_source_file1` FOREIGN KEY (`source_file_id`) REFERENCES `source_file` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `source_file_upload_history`
--

DROP TABLE IF EXISTS `source_file_upload_history`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `source_file_upload_history` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `source_file_id` bigint(20) NOT NULL,
  `file_name` varchar(200) NOT NULL,
  `wallet_address` varchar(200) NOT NULL,
  `status` varchar(45) NOT NULL,
  `create_at` bigint(20) NOT NULL,
  `update_at` bigint(20) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_source_file_upload_history_source_file_id` (`source_file_id`),
  CONSTRAINT `fk_source_file_upload_history_source_file_id` FOREIGN KEY (`source_file_id`) REFERENCES `source_file` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=274 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `system_config_param`
--

DROP TABLE IF EXISTS `system_config_param`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `system_config_param` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `param_key` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `param_value` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `module` varchar(128) COLLATE utf8_bin DEFAULT NULL,
  `desc` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `table_name_param_key_uindex` (`param_key`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping events for database 'mcp_v2'
--

--
-- Dumping routines for database 'mcp_v2'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-02-25 13:54:19
