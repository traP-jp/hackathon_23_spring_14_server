DROP TABLE IF EXISTS `users`;


CREATE TABLE `users` (
  `uuid` char(36) NOT NULL,
  `id` varchar(32) NOT NULL,
  `point` int DEFAULT NULL,
  `date` date NOT NULL,
  PRIMARY KEY (`uuid`,`date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `items`;

CREATE TABLE `items` (
  `uuid` char(36) NOT NULL,
  `id` varchar(32) NOT NULL,
  `description` varchar(86) DEFAULT NULL,
  `point` int NOT NULL,
  `report` int DEFAULT 0,
  PRIMARY KEY (`uuid`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `time_cards`;


CREATE TABLE `time_cards` (
  `id`char(36) NOT NULL,
  `date` date NOT NULL,
  `item_id` char(36) NOT NULL,
  PRIMARY KEY (`id`,`date`,`item_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `users` WRITE;
INSERT INTO `users` (`uuid`,`id`, `point`, `date`)
VALUES
  ("33a123f4-067b-4ec5-9060-4d03da4c4aca",'trap',2,"2023-06-17");
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;

LOCK TABLES `items` WRITE;
INSERT INTO `items` (`uuid`,`id`,`point`,`report`)
VALUES
  ("33a123f4-067b-4ec5-9060-4d03da4c4aca",'Get',2,5);
/*!40000 ALTER TABLE `items` DISABLE KEYS */;
UNLOCK TABLES;

LOCK TABLES `time_cards` WRITE;
INSERT INTO `time_cards` (`id`, `date`, `item_id`)
VALUES
  ('33a123f4-067b-4ec5-9060-4d03da4c4aca', '2023-06-17','33a123f4-067b-4ec5-9060-4d03da4c4aca');
/*!40000 ALTER TABLE `time_cards` DISABLE KEYS */;
UNLOCK TABLES;





