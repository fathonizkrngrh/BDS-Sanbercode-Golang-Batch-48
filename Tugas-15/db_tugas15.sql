-- MySQL dump 10.13  Distrib 8.0.34, for Win64 (x86_64)
--
-- Host: localhost    Database: db_tugas15
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
-- Table structure for table `mahasiswa`
--

DROP TABLE IF EXISTS `mahasiswa`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `mahasiswa` (
  `id` int NOT NULL AUTO_INCREMENT,
  `nama` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=20005 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `mahasiswa`
--

LOCK TABLES `mahasiswa` WRITE;
/*!40000 ALTER TABLE `mahasiswa` DISABLE KEYS */;
INSERT INTO `mahasiswa` VALUES (20001,'Fathoni','2023-08-19 12:53:06','2023-08-19 12:53:06'),(20002,'Zikri','2023-08-19 21:07:43','2023-08-19 21:07:43'),(20003,'Nugroho','2023-08-19 21:08:07','2023-08-19 21:08:07'),(20004,'Patini','2023-08-19 21:39:53','2023-08-19 21:39:53');
/*!40000 ALTER TABLE `mahasiswa` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `mata_kuliah`
--

DROP TABLE IF EXISTS `mata_kuliah`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `mata_kuliah` (
  `id` int NOT NULL AUTO_INCREMENT,
  `nama` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=105 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `mata_kuliah`
--

LOCK TABLES `mata_kuliah` WRITE;
/*!40000 ALTER TABLE `mata_kuliah` DISABLE KEYS */;
INSERT INTO `mata_kuliah` VALUES (101,'Analisis Kebutuhan Perangkat Lunak','2023-08-19 13:00:07','2023-08-19 13:00:07'),(102,'Keamanan Perangkat Lunak','2023-08-19 13:00:19','2023-08-19 13:01:06'),(103,'Pengenalan Rekayasa Perangkat Lunak','2023-08-19 21:08:26','2023-08-19 21:08:26'),(104,'Pengenalan Database','2023-08-19 21:40:43','2023-08-19 21:40:43');
/*!40000 ALTER TABLE `mata_kuliah` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `nilai`
--

DROP TABLE IF EXISTS `nilai`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `nilai` (
  `id` int NOT NULL AUTO_INCREMENT,
  `indeks` varchar(255) NOT NULL,
  `nilai` int NOT NULL,
  `mahasiswa_id` int NOT NULL,
  `mata_kuliah_id` int NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `mahasiswa_id` (`mahasiswa_id`),
  KEY `mata_kuliah_id` (`mata_kuliah_id`),
  CONSTRAINT `nilai_ibfk_1` FOREIGN KEY (`mahasiswa_id`) REFERENCES `mahasiswa` (`id`),
  CONSTRAINT `nilai_ibfk_2` FOREIGN KEY (`mata_kuliah_id`) REFERENCES `mata_kuliah` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `nilai`
--

LOCK TABLES `nilai` WRITE;
/*!40000 ALTER TABLE `nilai` DISABLE KEYS */;
INSERT INTO `nilai` VALUES (1,'A',90,20001,101,'2023-08-19 21:49:33','2023-08-19 22:25:36'),(2,'A',80,20001,102,'2023-08-19 22:20:47','2023-08-19 22:20:47');
/*!40000 ALTER TABLE `nilai` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-08-19 22:28:29
