#!/usr/bin/env bash

# Copyright (c) 2018-2019 Dunyu All Rights Reserved.
#
# Author : yangping
# Email  : youhei_yp@163.com
#
# Prismy.No | Date       | Modified by. | Description
# -------------------------------------------------------------------
# 00001       2019/05/22   yangping       New version
# -------------------------------------------------------------------

usage="Usage: clear.sh (database user)"

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
# e.g: APP_DATABASE_NAME="paster"                         #
# ------------------------------------------------------- #
APP_DATABASE_NAME="database_sample"

# execute clear script
mysql -u$LOGIN_USER -p -e "

USE "${APP_DATABASE_NAME}";

DELETE FROM goods;
DELETE FROM account;

"

# SELECT * FROM table LIMIT [[offset,] rows | rows OFFSET offset]
#
# e.g :
# select * from account where drafts>10 limit 10;
# -> return 10 rows result start 0 recode position
#
# select * from account where drafts>10 limit 1000, 10;
# select * from account where drafts>10 limit 10 offset 1000;
# -> return 10 rows result start 1000 recode position

# ALTER TABLE table ADD column datatype;
# e.g :
# alter table account add code varchar(10) not null;

# ALTER TABLE table MODIFY fromecolumn AFTER tocolumn
# e.g :
# alter table account modify uuid AFTER nickname
# -> move uuid column after nickname

# ALTER TABLE table MODIFY COLUMN column datatype
# e.g :
# alter table account modify column uuid varchar(255) character set utf8 not null default '--'
# -> change uuid column data length and set default string

# ALTER TABLE table CHANGE fromecolumn tocolumn datatype
# e.g :
# alter table account change uuid uid varchar(255) character set utf8 not null default '--'
# -> change uuid column name, data length and set default string
