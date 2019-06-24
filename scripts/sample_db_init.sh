#!/usr/bin/env bash

# Copyright (c) 2018-2019 WING All Rights Reserved.
#
# Author : yangping
# Email  : youhei_yp@163.com
#
# Prismy.No | Date       | Modified by. | Description
# -------------------------------------------------------------------
# 00001       2019/05/22   yangping       New version
# -------------------------------------------------------------------

usage="Usage: initdb.sh (database user)"

# if no args specified, show usage
if [ $# -ne 1 ]; then
  echo $usage
  exit 1
fi

LOGIN_USER=$1

# ------------------------------------------------------- #
# TODO :                                                  #
#                                                         #
# YOU MUST CHANGE database_sample TO YOUR DATABASE NAME   #
# e.g: APP_DATABASE_NAME="xappdb"                         #
# ------------------------------------------------------- #
APP_DATABASE_NAME="database_sample"

# execute init script
mysql -u$LOGIN_USER -p -e "
CREATE DATABASE IF NOT EXISTS "${APP_DATABASE_NAME}";

USE "${APP_DATABASE_NAME}";

CREATE TABLE IF NOT EXISTS account (
	id				int				NOT NULL AUTO_INCREMENT,
	uuid			varchar	(64)	CHARACTER SET utf8 NOT NULL,
	nickname		varchar	(64)	CHARACTER SET utf8 DEFAULT '',
	createtime		bigint			NOT NULL,
	drafts			int				DEFAULT 0,
	activities		int				DEFAULT 0,
	offshelves		int				DEFAULT 0,
	PRIMARY KEY (id),
	UNIQUE(uuid)
) DEFAULT CHARSET=utf8 COMMENT='Account table';

CREATE TABLE IF NOT EXISTS goods (
	id				int				NOT NULL AUTO_INCREMENT,
	aid				int				NOT NULL,
	box				int				DEFAULT 0,
	PRIMARY KEY (id),
) DEFAULT CHARSET=utf8 COMMENT='Goods table';

describe account;
describe goods;

DELIMITER $
DROP TRIGGER IF EXISTS trigger_insert_goods;
CREATE TRIGGER trigger_insert_goods
AFTER INSERT ON goods
FOR EACH ROW
BEGIN
    UPDATE account SET drafts=drafts+1 WHERE id=NEW.aid;
END
$

DELIMITER $
DROP TRIGGER IF EXISTS trigger_update_goods;
CREATE TRIGGER trigger_update_goods
AFTER UPDATE ON goods
FOR EACH ROW
BEGIN
    IF OLD.box = 0 THEN
        IF NEW.box = 1 THEN
            UPDATE account SET drafts=drafts-1, activities=activities+1 WHERE id=OLD.aid;
        ELSEIF NEW.box = 2 THEN
            UPDATE account SET drafts=drafts-1, offshelves=offshelves+1 WHERE id=OLD.aid;
        END IF;
    ELSEIF OLD.box = 1 THEN
        IF NEW.box = 0 THEN
            UPDATE account SET activities=activities-1, drafts=drafts+1 WHERE id=OLD.aid;
        ELSEIF NEW.box = 2 THEN
            UPDATE account SET activities=activities-1, offshelves=offshelves+1 WHERE id=OLD.aid;
        END IF;
	ELSEIF OLD.box = 2 THEN
        IF NEW.box = 0 THEN
            UPDATE account SET offshelves=offshelves-1, drafts=drafts+1 WHERE id=OLD.aid;
        ELSEIF NEW.box = 1 THEN
            UPDATE account SET offshelves=offshelves-1, activities=activities+1 WHERE id=OLD.aid;
        END IF;
    END IF;
END
$

DELIMITER $
DROP TRIGGER IF EXISTS trigger_delete_goods;
CREATE TRIGGER trigger_delete_goods
AFTER DELETE ON goods
FOR EACH ROW
BEGIN
    IF OLD.box = 0 THEN
        UPDATE account SET drafts=drafts-1 WHERE id=OLD.aid;
    ELSEIF OLD.box = 1 THEN
        UPDATE account SET activities=activities-1 WHERE id=OLD.aid;
    ELSEIF OLD.box = 2 THEN
        UPDATE account SET offshelves=offshelves-1 WHERE id=OLD.aid;
    END IF;
END
$

"