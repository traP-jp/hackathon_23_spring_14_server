DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` varchar(32) NOT NULL,
  `password` varchar(86) DEFAULT NULL,
  `email` varchar(191) DEFAULT NULL,
  `student_number` varchar(8) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`),
  UNIQUE KEY `student_number` (`student_number`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;

INSERT INTO `users` (`id`, `password`, `email`, `student_number`)
VALUES
  ('trap','-dP58137PnIQ7tjYBUiiU9VYMu9vZnfYNfT2G3zFStDR3VWap08funuhYbLHENkPv8xKkk05JdbPa-iCx7Wtww','trap@example.com','15B00001');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;




