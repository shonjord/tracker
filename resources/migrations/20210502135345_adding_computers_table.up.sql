CREATE TABLE computers (
     id MEDIUMINT NOT NULL AUTO_INCREMENT,
     uuid CHAR(36) NOT NULL,
     employee_uuid CHAR(36) NULL,
     mac_address VARCHAR(255) NOT NULL,
     ip_address VARCHAR(255) NOT NULL,
     name VARCHAR(255) NOT NULL,
     description VARCHAR(255) NULL,
     created_at DATETIME NOT NULL,
     updated_at DATETIME NOT NULL,
     UNIQUE KEY `UNIQ_COMPUTER_UUID` (`uuid`),
     UNIQUE KEY `UNIQ_MAC_ADDRESS` (`mac_address`),
     INDEX IDX_UUID (uuid),
     PRIMARY KEY (id),
     CONSTRAINT FK_EMPLOYEE FOREIGN KEY (`employee_uuid`) REFERENCES employees (uuid)
) DEFAULT CHARACTER SET UTF8 COLLATE `UTF8_unicode_ci` ENGINE = InnoDB;
