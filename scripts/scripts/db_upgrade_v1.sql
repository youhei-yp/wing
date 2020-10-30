CREATE DATABASE IF NOT EXISTS xxxdb;

USE xxxdb;


-- use procedure batch to
-- add describtion and activate columns into medias table
DELIMITER // -- user defined terminator, try to avoid\\
DROP PROCEDURE IF EXISTS medias_AddColumn //
CREATE PROCEDURE medias_AddColumn() BEGIN

IF EXISTS(SELECT NULL FROM information_schema.columns WHERE table_name='medias' AND COLUMN_NAME='titlefont') THEN
ALTER TABLE medias CHANGE COLUMN titlefont tfsizecn int DEFAULT 24;
END IF;

IF NOT EXISTS(SELECT NULL FROM information_schema.columns WHERE table_name='medias' AND COLUMN_NAME='title') THEN
ALTER TABLE medias ADD COLUMN title varchar(512) CHARACTER SET utf8mb4 DEFAULT '';
END IF;

IF NOT EXISTS(SELECT NULL FROM information_schema.columns WHERE table_name='medias' AND COLUMN_NAME='activate') THEN
ALTER TABLE medias ADD COLUMN activate bool DEFAULT FALSE;
END IF;

IF EXISTS(SELECT NULL FROM information_schema.columns WHERE table_name='products' AND COLUMN_NAME='prodnum') THEN
ALTER TABLE medias MODIFY COLUMN prodnum varchar(64) CHARACTER SET utf8mb4 DEFAULT '';
END IF;

END //
CALL medias_AddColumn;
DROP PROCEDURE medias_AddColumn;
//


-- add trigger to flag unused state when delete photo flow node
-- 7 : IT_Course - ImgType - types.go
DELIMITER $
DROP TRIGGER IF EXISTS trigger_delete_course;
CREATE TRIGGER trigger_delete_course
AFTER DELETE ON course
FOR EACH ROW
BEGIN
    UPDATE medias SET isused=false WHERE owner=OLD.id AND category=7;
END
$

