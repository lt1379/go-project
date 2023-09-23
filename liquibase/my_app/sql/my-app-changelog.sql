--liquibase formatted sql

--changeset lamboktulus1379:1 labels:my-app-label context:my-app-context
--comment: my-app creating table persons
CREATE TABLE `persons` (
  `name` VARCHAR(50) NOT NULL,
  `country` VARCHAR(60) NOT NULL,
  INDEX `idx_name` (`name` ASC) VISIBLE);
--rollback DROP TABLE persons;

--changeset lamboktulus1379:2 labels:my-app-label context:my-app-context
--comment: my-app insert data into table persons
INSERT INTO `persons` (`name`, `country`) values ('Adam', 'Kuala Lumpur');
--rollback DELETE `persons` WHERE `name` = 'Adam';

--changeset lamboktulus1379:3 labels:my-app-label context:my-app-context
--comment: my-app insert data into table persons
INSERT INTO `persons` (`name`, `country`) values ('John', 'Singapore');
--rollback DELETE `persons` WHERE `name` = 'John';

--changeset lamboktulus1379:4 labels:my-app-label context:my-app-context
--comment: my-app insert data into table persons
INSERT INTO `persons` (`name`, `country`) values ('Henry', 'Singapore');
--rollback DELETE `persons` WHERE `name` = 'Henry';

--changeset lamboktulus1379:5 labels:my-app-label context:my-app-context
--comment: my-app insert data into table persons
INSERT INTO `persons` (`name`, `country`) values ('Dominic', 'Thailand');
--rollback DELETE `persons` WHERE `name` = 'Dominic';