CREATE DATABASE testdb1;
CREATE TABLE `testdb1`.`table1` (`id` INT NOT NULL AUTO_INCREMENT , `text1` VARCHAR(10) NOT NULL , `number1` INT NOT NULL , `date1` DATE NULL , PRIMARY KEY (`id`)) ENGINE = InnoDB;
INSERT INTO `testdb1`.`table1` (`id`, `text1`, `number1`, `date1`) VALUES (NULL, 'text1', '1', NULL);
INSERT INTO `testdb1`.`table1` (`id`, `text1`, `number1`, `date1`) VALUES (NULL, 'text2', '2', NULL);
INSERT INTO `testdb1`.`table1` (`id`, `text1`, `number1`, `date1`) VALUES (NULL, 'text3', '3', NULL);
CREATE TABLE `testdb1`.`table2` (`id` INT NOT NULL AUTO_INCREMENT , PRIMARY KEY (`id`)) ENGINE = InnoDB;
INSERT INTO `testdb1`.`table2` (`id`) VALUES (NULL);
CREATE DATABASE testdb2;
CREATE TABLE `testdb2`.`table2` (`id` INT NOT NULL AUTO_INCREMENT , PRIMARY KEY (`id`)) ENGINE = InnoDB;
INSERT INTO `testdb2`.`table2` (`id`) VALUES (NULL);
INSERT INTO `testdb2`.`table2` (`id`) VALUES (NULL);
INSERT INTO `testdb2`.`table2` (`id`) VALUES (NULL);