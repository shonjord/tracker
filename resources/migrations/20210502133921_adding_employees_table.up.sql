CREATE TABLE employees (
    id MEDIUMINT NOT NULL AUTO_INCREMENT,
    uuid CHAR(36) NOT NULL ,
    name VARCHAR(255) NOT NULL ,
    abbreviation CHAR(3) NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    UNIQUE KEY `UNIQ_EMPLOYEE_UUID` (`uuid`),
    INDEX IDX_UUID (uuid),
    PRIMARY KEY (id)
) DEFAULT CHARACTER SET UTF8 COLLATE `UTF8_unicode_ci` ENGINE = InnoDB;

INSERT INTO employees (uuid, name, abbreviation, created_at, updated_at)
VALUES
    ('20587b2c-3969-49b6-add1-27fe09006ef9', 'Albert', NULL, NOW(), NOW()),
    ('1c5f944a-b9fa-41e1-a83c-8e6c5dea3a82', 'Paola', 'pal', NOW(), NOW())
