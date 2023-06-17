DROP TABLE IF EXISTS `users`;
DROP TABLE IF EXISTS `time_cards`;
DROP TABLE IF EXISTS `items`;

CREATE TABLE `users` (
  `uuid` varchar(32),
  `id` varchar(32),
  `point` int DEFAULT NULL,
  `date` date NOT NULL,
  PRIMARY KEY (`uuid`,`date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `time_cards` (
  `id` varchar(32),
  `date` date,
  `item_id` varchar(86) NOT NULL,
  PRIMARY KEY (`id`,`date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `items` (
  `id` varchar(32) NOT NULL,
  `description` varchar(86) DEFAULT NULL,
  `point` int DEFAULT NULL,
  `report` int DEFAULT NULL,
  UNIQUE KEY `id` (`id`),
  UNIQUE KEY `description` (`description`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `users` WRITE;
INSERT INTO `users` (`uuid`,`id`, `point`, `date`)
VALUES
  (`fwafewa`,'trap',2,"2023-06-17");
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;

LOCK TABLES `time_cards` WRITE;
INSERT INTO `time_cards` (`id`, `date`, `item_id`)
VALUES
  ('traP', '2023-06-17','Get');
/*!40000 ALTER TABLE `time_cards` DISABLE KEYS */;
UNLOCK TABLES;

LOCK TABLES `items` WRITE;
INSERT INTO `items` (`id`,`point`,`report`)
VALUES
  ('Get',3,5);
/*!40000 ALTER TABLE `items` DISABLE KEYS */;
UNLOCK TABLES;




