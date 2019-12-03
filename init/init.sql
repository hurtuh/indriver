CREATE DATABASE IF NOT EXISTS `interviews` CHARACTER SET utf8 COLLATE utf8_general_ci;
CREATE USER IF NOT EXISTS 'inDriver'@'%' IDENTIFIED BY 'IsCool';
GRANT ALL PRIVILEGES ON *.* TO 'inDriver'@'%';
USE `interviews`;
CREATE TABLE IF NOT EXISTS `interviews`.`candidate` (`id` int auto_increment, PRIMARY KEY(id), `created` datetime default NOW(), `name` varchar(200), `lastname` varchar(200), `interview` datetime, `description` varchar(200));
INSERT INTO `interviews`.candidate (created, name, lastname, interview, description) VALUE (NOW(), 'evgeny', 'lilekov', '2019-12-06', 'Good candidate, you have to take');
