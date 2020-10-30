CREATE DATABASE IF NOT EXISTS scenegreetingdb;

USE scenegreetingdb;

CREATE TABLE IF NOT EXISTS account (
    uuid             varchar (64)       NOT NULL,
    pwd              varchar (24)       NOT NULL,
    pubkey           varchar (512)      NOT NULL,
    prikey           varchar (1024)     NOT NULL,
    PRIMARY KEY (uuid)
) DEFAULT CHARSET=utf8mb4 COMMENT='Account table';

-- comment out describe table structure
-- describe account;

-- add trigger to flag unused state when delete greeting node
-- 0 : IT_Greeting - ImgType - types.go
DELIMITER $
DROP TRIGGER IF EXISTS trigger_delete_greetings;
CREATE TRIGGER trigger_delete_greetings
AFTER DELETE ON greetings
FOR EACH ROW
BEGIN
    UPDATE medias SET isused=false WHERE owner=OLD.id AND category=0;
END
$

-- comment out these trigger and re-create at the version 1
--
-- DELIMITER $
-- DROP TRIGGER IF EXISTS trigger_delete_products;
-- CREATE TRIGGER trigger_delete_products
-- AFTER DELETE ON products
-- FOR EACH ROW
-- BEGIN
--     UPDATE medias SET isused=false WHERE owner=OLD.id AND (category=1 OR category=2 OR category=3 OR category=9);
-- END
-- $
--
-- DELIMITER $
-- DROP TRIGGER IF EXISTS trigger_delete_declaration;
-- CREATE TRIGGER trigger_delete_declaration
-- AFTER DELETE ON declaration
-- FOR EACH ROW
-- BEGIN
--     UPDATE medias SET isused=false WHERE owner=OLD.id AND category=7;
-- END
-- $
