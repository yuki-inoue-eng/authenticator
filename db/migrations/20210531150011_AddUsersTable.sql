
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `Users` (
  `id` int AUTO_INCREMENT,
  `login_id` varchar(20) NOT NULL DEFAULT '',
  `password` varchar(70) NOT NULL DEFAULT '',
  `qos` varchar(20) NOT NULL DEFAULT 'default',
  `last_login_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  UNIQUE KEY (`login_id`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- LOCK TABLES `Users` WRITE;
-- INSERT INTO `Users` (`login_id`, `password`, `last_login_at`, `created_at`, `updated_at`, `deleted_at`)
-- VALUES ('test_login_id','test_password','2021-05-27 20:08:46','2021-05-27 20:09:46','2021-05-27 20:10:46', NULL);
-- UNLOCK TABLES;


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `Users`;
