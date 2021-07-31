CREATE SCHEMA `koperasi` ;

CREATE TABLE `koperasi`.`member` (
  `id` BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID())),
  `member_name` VARCHAR(45) NOT NULL,
  `kta_number` VARCHAR(45) NOT NULL,
  `nik_number` VARCHAR(45) NOT NULL,
  `date_of_birth` DATE NOT NULL,
  `place_of_birth` VARCHAR(45) NULL,
  `address` VARCHAR(200) NOT NULL,
  `business_sector` VARCHAR(45) NULL,
  `phone_number` VARCHAR(45) NOT NULL,
  `status` VARCHAR(2) NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NULL,
  `deleted_at` DATETIME NULL,
  PRIMARY KEY (`id`));


INSERT INTO `koperasi`.`member` (`member_name`, `kta_number`, `nik_number`,  `date_of_birth`, `place_of_birth`, `address`, `business_sector`, `phone_number`, `status`) 
VALUES ('Sinar Dyas 1', '999999', '999999999', '1996-05-31', 'Karanganyar', 'Kab. Tangerang, Banten', 'Technology', '082227777619', '1');
INSERT INTO `koperasi`.`member` (`member_name`, `kta_number`, `nik_number`,  `date_of_birth`, `place_of_birth`, `address`, `business_sector`, `phone_number`, `status`) 
VALUES ('Sinar Dyas 2', '999998', '888888888', '1996-05-31', 'Karanganyar', 'Kab. Tangerang, Banten', 'Technology', '082227777619', '1');
INSERT INTO `koperasi`.`member` (`member_name`, `kta_number`, `nik_number`,  `date_of_birth`, `place_of_birth`, `address`, `business_sector`, `phone_number`, `status`) 
VALUES ('Sinar Dyas 3', '999997', '777777777', '1996-05-31', 'Karanganyar', 'Kab. Tangerang, Banten', 'Technology', '082227777619', '1');
INSERT INTO `koperasi`.`member` (`member_name`, `kta_number`, `nik_number`,  `date_of_birth`, `place_of_birth`, `address`, `business_sector`, `phone_number`, `status`) 
VALUES ('Sinar Dyas 4', '999996', '666666666', '1996-05-31', 'Karanganyar', 'Kab. Tangerang, Banten', 'Technology', '082227777619', '1');
INSERT INTO `koperasi`.`member` (`member_name`, `kta_number`, `nik_number`,  `date_of_birth`, `place_of_birth`, `address`, `business_sector`, `phone_number`, `status`) 
VALUES ('Sinar Dyas 5', '999995', '555555555', '1996-05-31', 'Karanganyar', 'Kab. Tangerang, Banten', 'Technology', '082227777619', '1');
INSERT INTO `koperasi`.`member` (`member_name`, `kta_number`, `nik_number`,  `date_of_birth`, `place_of_birth`, `address`, `business_sector`, `phone_number`, `status`) 
VALUES ('Sinar Dyas 6', '999994', '444444444', '1996-05-31', 'Karanganyar', 'Kab. Tangerang, Banten', 'Technology', '082227777619', '1');
INSERT INTO `koperasi`.`member` (`member_name`, `kta_number`, `nik_number`,  `date_of_birth`, `place_of_birth`, `address`, `business_sector`, `phone_number`, `status`) 
VALUES ('Sinar Dyas 7', '999993', '333333333', '1996-05-31', 'Karanganyar', 'Kab. Tangerang, Banten', 'Technology', '082227777619', '1');
INSERT INTO `koperasi`.`member` (`member_name`, `kta_number`, `nik_number`,  `date_of_birth`, `place_of_birth`, `address`, `business_sector`, `phone_number`, `status`) 
VALUES ('Sinar Dyas 8', '999992', '222222222', '1996-05-31', 'Karanganyar', 'Kab. Tangerang, Banten', 'Technology', '082227777619', '1');
INSERT INTO `koperasi`.`member` (`member_name`, `kta_number`, `nik_number`,  `date_of_birth`, `place_of_birth`, `address`, `business_sector`, `phone_number`, `status`) 
VALUES ('Sinar Dyas 9', '999991', '111111111', '1996-05-31', 'Karanganyar', 'Kab. Tangerang, Banten', 'Technology', '082227777619', '1');
INSERT INTO `koperasi`.`member` (`member_name`, `kta_number`, `nik_number`,  `date_of_birth`, `place_of_birth`, `address`, `business_sector`, `phone_number`, `status`) 
VALUES ('Sinar Dyas 10', '999990', '000000000', '1996-05-31', 'Karanganyar', 'Kab. Tangerang, Banten', 'Technology', '082227777619', '1');