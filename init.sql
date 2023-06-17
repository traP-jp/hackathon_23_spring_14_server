DROP TABLE IF EXISTS `users`;
DROP TABLE IF EXISTS `timeCards`;
DROP TABLE IF EXISTS `items`;

CREATE TABLE `users` (
  `id` varchar(32) NOT NULL,
  `datePoint` int(11) DEFAULT NULL,
  `date` date NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `timeCards` (
  `date` date NOT NULL,
  `id` varchar(32) NOT NULL,
  `itemID` varchar(86) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `items` (
  `id` varchar(32) NOT NULL,
  `description` varchar(86) DEFAULT NULL,
  `point` int(11) DEFAULT NULL,
  `report` int(11) DEFAULT NULL,
  UNIQUE KEY `id` (`id`)/*,
  UNIQUE KEY `description` (`description`)*/
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `users` WRITE;
INSERT INTO `users` (`id`, `datePoint`, `date`)
VALUES
  ('trap',2,'20230617');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;

LOCK TABLES `timeCards` WRITE;
INSERT INTO `timeCards` (`date`, `id`, `itemID`)
VALUES
  ('20230617','traP','Get Up early');
/*!40000 ALTER TABLE `timeCards` DISABLE KEYS */;
UNLOCK TABLES;

LOCK TABLES `items` WRITE;
INSERT INTO `items` (`id`,`point`,`report`)
VALUES
  ('Get Up early',2,0);
/*!40000 ALTER TABLE `items` DISABLE KEYS */;
UNLOCK TABLES;




