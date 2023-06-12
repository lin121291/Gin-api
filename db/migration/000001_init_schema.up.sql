CREATE TABLE `users` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `username` VARCHAR(255) NOT NULL,
  `password` VARCHAR(255) NOT NULL
);

CREATE TABLE `news` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `title` VARCHAR(255) NOT NULL,
  `content` TEXT,
  `date_published` DATE
);

CREATE TABLE `reading_list` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `user_id` INT,
  `news_id` INT
);

ALTER TABLE `reading_list` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `reading_list` ADD FOREIGN KEY (`news_id`) REFERENCES `news` (`id`);
  