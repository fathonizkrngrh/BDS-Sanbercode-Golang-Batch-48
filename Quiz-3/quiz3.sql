-- MySQL dump 10.13  Distrib 8.0.34, for Win64 (x86_64)
--
-- Host: localhost    Database: quiz3
-- ------------------------------------------------------
-- Server version	8.0.34

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `book`
--

DROP TABLE IF EXISTS `book`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `book` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `image_url` varchar(255) NOT NULL,
  `release_year` int NOT NULL,
  `price` varchar(255) NOT NULL,
  `total_page` int NOT NULL,
  `thickness` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `category_id` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `category_id` (`category_id`),
  CONSTRAINT `book_ibfk_1` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `book`
--

LOCK TABLES `book` WRITE;
/*!40000 ALTER TABLE `book` DISABLE KEYS */;
INSERT INTO `book` VALUES (1,'The Book of Fire','A tale of magic and adventure in a world of fire.','https://example.com/fire-book.jpg',2020,'15.99',320,'tebal','2023-08-21 11:19:33','2023-08-21 11:19:33',3),(2,'Echoes of Eternity','An epic fantasy journey through time and space.','https://example.com/echoes-book.jpg',2020,'12.99',280,'tebal','2023-08-21 11:20:10','2023-08-21 11:20:10',3),(3,'The Last of the Lost','Surviving in a post-apocalyptic wasteland.','https://example.com/lost-book.jpg',2021,'11.99',240,'tebal','2023-08-21 11:20:23','2023-08-21 11:20:23',3),(4,'Realm of Shadows','A mystery set in a world of darkness and intrigue.','https://example.com/shadows-book.jpg',2019,'14.99',300,'tebal','2023-08-21 11:21:34','2023-08-21 11:21:34',3),(5,'The Enchanted Forest','Magical creatures and adventures in an enchanted realm.','https://example.com/enchanted-book.jpg',2019,'9.99',200,'sedang','2023-08-21 11:21:45','2023-08-21 11:21:45',3),(6,'Whispers of Wind','A story of discovery and self-discovery in a world of air and sky.','https://example.com/wind-book.jpg',2020,'10.99',240,'tebal','2023-08-21 11:22:03','2023-08-21 11:22:03',4),(7,'The Wandering Knight','A knight\'s journey across vast landscapes and untold dangers.','https://example.com/knight-book.jpg',2018,'12.99',280,'tebal','2023-08-21 11:22:13','2023-08-21 11:22:13',4),(8,'Love in Bloom','Romantic tales of love and passion in various settings.','https://example.com/love-book.jpg',2020,'9.99',220,'tebal','2023-08-21 11:22:46','2023-08-21 11:22:46',5),(9,'A Heart Unbound','Exploring the complexities of relationships and emotions.','https://example.com/heart-book.jpg',2021,'10.99',260,'tebal','2023-08-21 11:22:56','2023-08-21 11:22:56',5),(10,'Destined Together','Love that transcends time and space.','https://example.com/destined-book.jpg',2019,'11.99',280,'tebal','2023-08-21 11:23:12','2023-08-21 11:23:12',5),(11,'Beyond the Horizon','Exploring the mysteries of distant lands and ancient secrets.','https://example.com/beyond-book.jpg',1998,'8.99',280,'tebal','2023-08-21 11:24:23','2023-08-21 11:24:23',5),(12,'The Forgotten Chronicles','Rediscovering forgotten tales and legends of old.','https://example.com/forgotten-book.jpg',1995,'9.99',240,'tebal','2023-08-21 11:24:34','2023-08-21 11:24:34',5),(13,'Hearts Entwined','Entwined destinies and hearts that beat as one.','https://example.com/entwined-book.jpg',1994,'7.99',260,'tebal','2023-08-21 11:25:04','2023-08-21 11:25:04',6),(14,'Eternal Love','Love that defies time and death itself.','https://example.com/eternal-love-book.jpg',1988,'6.99',280,'tebal','2023-08-21 11:25:15','2023-08-21 11:25:15',6),(15,'The Power of Habits','Exploring the science behind habits and how they shape our lives.','https://example.com/habits-book.jpg',2012,'14.99',320,'tebal','2023-08-21 11:26:17','2023-08-21 11:26:17',8),(16,'Influence: The Psychology of Persuasion','Understanding the psychological principles behind persuasion and influence.','https://example.com/influence-book.jpg',1984,'12.99',320,'tebal','2023-08-21 11:26:27','2023-08-21 11:26:27',8),(17,'Thinking, Fast and Slow','Exploring the two systems of thinking that drive our decisions.','https://example.com/thinking-slow-book.jpg',2011,'15.99',480,'tebal','2023-08-21 11:26:37','2023-08-21 11:26:37',8),(18,'Emotional Intelligence','The importance of emotional intelligence in personal and professional success.','https://example.com/emotional-intelligence-book.jpg',1995,'11.99',320,'tebal','2023-08-21 11:26:53','2023-08-21 11:26:53',8),(19,'Drive: The Surprising Truth About What Motivates Us','Examining the science of motivation and how it impacts performance.','https://example.com/drive-book.jpg',2009,'13.99',288,'tebal','2023-08-21 11:27:04','2023-08-21 11:27:04',7),(20,'Good to Great','Discovering the factors that differentiate great companies from good ones.','https://example.com/good-to-great-book.jpg',2001,'14.99',320,'tebal','2023-08-21 11:27:12','2023-08-21 11:27:12',7),(21,'How to Win Friends and Influence People','Timeless advice on building meaningful relationships and influencing others.','https://example.com/win-friends-book.jpg',1997,'9.99',288,'tebal','2023-08-21 11:27:34','2023-08-21 11:27:34',7),(22,'Adventures of Spacecat','Join Spacecat on an intergalactic journey filled with humor and excitement.','https://example.com/spacecat-book.jpg',1990,'8.99',150,'sedang','2023-08-21 11:33:47','2023-08-21 11:33:47',1),(23,'Mysteries of the Mind','A collection of mind-bending stories that challenge reality.','https://example.com/mind-book.jpg',1985,'7.99',180,'sedang','2023-08-21 11:34:01','2023-08-21 11:34:01',1),(24,'Dreamscapes','Enter a world of dreams and explore the surreal landscapes within.','https://example.com/dreamscapes-book.jpg',1988,'9.99',170,'sedang','2023-08-21 11:34:45','2023-08-21 11:34:45',1),(25,'The Laughing Shadows','Dark humor and twisted tales that will leave you chuckling.','https://example.com/shadows-book.jpg',1995,'8.99',190,'sedang','2023-08-21 11:34:54','2023-08-21 11:34:54',1),(26,'Tales of Timeless Titans','Uncover the stories of larger-than-life characters from ages past.','https://example.com/titans-book.jpg',1986,'8.99',190,'sedang','2023-08-21 12:01:29','2023-08-21 12:01:29',1);
/*!40000 ALTER TABLE `book` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `category`
--

DROP TABLE IF EXISTS `category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `category` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `category`
--

LOCK TABLES `category` WRITE;
/*!40000 ALTER TABLE `category` DISABLE KEYS */;
INSERT INTO `category` VALUES (1,'Comic','2023-08-20 18:37:20','2023-08-20 18:37:20'),(3,'Fiction','2023-08-21 11:00:08','2023-08-21 11:00:08'),(4,'Dystopian','2023-08-21 11:00:15','2023-08-21 11:00:15'),(5,'Fantasy','2023-08-21 11:00:20','2023-08-21 11:00:20'),(6,'Romance','2023-08-21 11:00:27','2023-08-21 11:00:27'),(7,'Business','2023-08-21 11:00:33','2023-08-21 11:00:33'),(8,'Psychology','2023-08-21 11:00:39','2023-08-21 11:00:39');
/*!40000 ALTER TABLE `category` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-08-21 12:05:40
