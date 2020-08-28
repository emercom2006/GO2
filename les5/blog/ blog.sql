-- MySQL dump 10.13  Distrib 5.7.17, for macos10.12 (x86_64)
--
-- Host: localhost    Database: blog_bd
-- ------------------------------------------------------
-- Server version	5.7.26

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
-- Table structure for table `tablename`
--

DROP TABLE IF EXISTS `tablename`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tablename` (
  `Id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `Title` varchar(255) DEFAULT NULL,
  `Short` mediumtext,
  `Long` mediumtext,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tablename`
--

LOCK TABLES `tablename` WRITE;
/*!40000 ALTER TABLE `tablename` DISABLE KEYS */;
INSERT INTO `tablename` VALUES (1,'Go — простой язык для серьезных задач','Ответ на вопрос, какой язык выучить в новом году и полное руководство к действию.','Если вы не знаете, какой язык программирования выучить в 2016 году — попробуйте освоить Go. Основное преимущество языка в эффективности разработки, так что его можно применять в серьезных проектах. При этом он прост в изучении: если вы уже программируете на Python или C++, то научитесь писать годный код на нем буквально за несколько дней.\n\n«Ровно год назад в пятницу я получил offer от Lazada (без знания Go), и в понедельник, после собеседования, HR’ы подписали мне его на должность Senior’a» — Иван Кищенко, Lazada Tech Hub.\n\nДля человека, знающего ООП и С++, изучение языка Go не проблема.'),(2,'Хочу всё знать. Язык Go','Один из лидеров уходящего года.','Надеюсь, вы уже ознакомились со свежими рейтингами, поэтому с сегодняшним героем нашей традиционной рубрики более-менее знакомы. В прошлом году Go появлялся в новостях почти каждую неделю, рейтинги популярности ползли вверх, книжные развалы заполонили профессиональные книги. Сегодня ажиотаж немного угас, гики трезво оценили качество языка. Никто уже не говорит о тотальном доминировании, но стало ясно, что у Go яркое будущее. Идеальное время, чтобы познакомиться с ним и вам.');
/*!40000 ALTER TABLE `tablename` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-08-28 20:02:48
