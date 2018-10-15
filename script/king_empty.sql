-- MySQL dump 10.13  Distrib 5.7.22, for Linux (x86_64)
--
-- Host: localhost    Database: king
-- ------------------------------------------------------
-- Server version	5.7.22-0ubuntu18.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `SignInfo`
--

DROP TABLE IF EXISTS `SignInfo`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `SignInfo` (
  `player_id` int(11) unsigned NOT NULL,
  `last_sign_time` int(11) unsigned NOT NULL,
  `week_sign_days` int(11) unsigned NOT NULL DEFAULT '0',
  `month_sign_days` int(11) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`player_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `SignInfo`
--

LOCK TABLES `SignInfo` WRITE;
/*!40000 ALTER TABLE `SignInfo` DISABLE KEYS */;
/*!40000 ALTER TABLE `SignInfo` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `army`
--

DROP TABLE IF EXISTS `army`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `army` (
  `player_id` int(11) NOT NULL,
  `army_id` int(11) NOT NULL,
  `hero_id1` int(11) NOT NULL DEFAULT '0',
  `hero_id2` int(11) NOT NULL DEFAULT '0',
  `hero_id3` int(11) NOT NULL DEFAULT '0',
  `hero_id4` int(11) NOT NULL DEFAULT '0',
  `hero_id5` int(11) NOT NULL DEFAULT '0',
  UNIQUE KEY `index` (`player_id`,`army_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `army`
--

LOCK TABLES `army` WRITE;
/*!40000 ALTER TABLE `army` DISABLE KEYS */;
/*!40000 ALTER TABLE `army` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `building`
--

DROP TABLE IF EXISTS `building`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `building` (
  `player_id` int(11) NOT NULL,
  `farm1_lv` int(11) unsigned NOT NULL DEFAULT '0',
  `farm2_lv` int(11) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`player_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `building`
--

LOCK TABLES `building` WRITE;
/*!40000 ALTER TABLE `building` DISABLE KEYS */;
/*!40000 ALTER TABLE `building` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `city`
--

DROP TABLE IF EXISTS `city`;

--
-- Table structure for table `hero`
--

DROP TABLE IF EXISTS `hero`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `hero` (
  `player_id` int(11) NOT NULL,
  `id` int(11) NOT NULL,
  `exp` int(11) NOT NULL DEFAULT '0',
  `level` int(11) NOT NULL DEFAULT '0',
  `star` int(11) NOT NULL DEFAULT '0',
  `army_id` int(11) NOT NULL DEFAULT '0',
  `army_lv` int(11) NOT NULL DEFAULT '0',
  `army_skill1_lv` int(11) NOT NULL DEFAULT '0',
  `army_skill2_lv` int(11) NOT NULL DEFAULT '0',
  `army_skill3_lv` int(11) NOT NULL DEFAULT '0',
  `chip` int(11) NOT NULL DEFAULT '0',
  UNIQUE KEY `hindex` (`player_id`,`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `hero`
--

LOCK TABLES `hero` WRITE;
/*!40000 ALTER TABLE `hero` DISABLE KEYS */;
/*!40000 ALTER TABLE `hero` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `map`
--

DROP TABLE IF EXISTS `map`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `map` (
  `player_id` int(11) NOT NULL,
  `map_id` int(11) NOT NULL,
  `city_id1` int(11) NOT NULL DEFAULT '0',
  `city_id2` int(11) NOT NULL DEFAULT '0',
  `city_id3` int(11) NOT NULL DEFAULT '0',
  `city_id4` int(11) NOT NULL DEFAULT '0',
  `city_id5` int(11) NOT NULL DEFAULT '0',
  UNIQUE KEY `mindex` (`player_id`,`map_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `map`
--

LOCK TABLES `map` WRITE;
/*!40000 ALTER TABLE `map` DISABLE KEYS */;
/*!40000 ALTER TABLE `map` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `map_city`
--

DROP TABLE IF EXISTS `map_city`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `map_city` (
  `player_id` int(11) NOT NULL,
  `map_city_idx` int(11) unsigned NOT NULL DEFAULT '0',
  `map_id` int(11) unsigned NOT NULL DEFAULT '0',
  `city_id` int(11) unsigned NOT NULL DEFAULT '0',
  `id_type` int(11) unsigned NOT NULL DEFAULT '0',
  `id` int(11) unsigned NOT NULL DEFAULT '0',
  `x` int(11) unsigned NOT NULL DEFAULT '0',
  `y` int(11) unsigned NOT NULL DEFAULT '0',
  `status` int(11) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`player_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `map_city`
--

LOCK TABLES `map_city` WRITE;
/*!40000 ALTER TABLE `map_city` DISABLE KEYS */;
/*!40000 ALTER TABLE `map_city` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `player`
--

DROP TABLE IF EXISTS `player`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player` (
  `player_id` int(11) NOT NULL,
  `name` varchar(64) NOT NULL,
  `url` varchar(256) NOT NULL,
  `open_id` varchar(64) NOT NULL,
  `coin` int(11) NOT NULL DEFAULT '0',
  `diamond` int(11) NOT NULL DEFAULT '0',
  `mineral` int(11) NOT NULL DEFAULT '0',
  `grain` int(11) NOT NULL DEFAULT '0',
  `wood` int(11) NOT NULL DEFAULT '0',
  `level` int(11) NOT NULL DEFAULT '0',
  `five_star_hero_recruit_times` int(11) NOT NULL DEFAULT '0',
  `ten_times_hero_recurit_times` int(11) NOT NULL DEFAULT '0',
  `free_hero_recruit_timestamp` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`player_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `player`
--

LOCK TABLES `player` WRITE;
/*!40000 ALTER TABLE `player` DISABLE KEYS */;
/*!40000 ALTER TABLE `player` ENABLE KEYS */;
UNLOCK TABLES;

DROP TABLE IF EXISTS `pve_level`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `pve_level` (
  `player_id` int(11) NOT NULL,
  `level_id` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`player_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `open_id` varchar(64) NOT NULL,
  `player_id` int(11) NOT NULL,
  `name` varchar(64) NOT NULL,
  `url` varchar(256) NOT NULL,
  PRIMARY KEY (`open_id`),
  UNIQUE KEY `player_id` (`player_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-07-28 14:40:49
