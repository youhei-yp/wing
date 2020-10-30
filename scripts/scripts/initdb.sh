#!/usr/bin/env bash

# Copyright (c) 2019-2029 Dunyu All Rights Reserved.
#
# Author      : yangping
# Email       : youhei_yp@163.com
# Version     : 1.0.1
# Description :
#   Create database for xxx server.
#
# Prismy.No | Date       | Modified by. | Description
# -------------------------------------------------------------------
# 00001       2020/05/08   yangping       New version
# 00002       2020/08/16   yangping       Support for windows
# -------------------------------------------------------------------

# if no args specified, show usage
if [ $# -ne 1 ]; then
  echo "Usage: initdb.sh (database user)"
  exit 1
fi

# enter service bin folder
bin=`dirname "$0"`
bin=`cd "$bin"; pwd`
source ${bin}/exports.sh

DATABASE_USER=$1

# init database from .sql file
echo "Start init database ${SERVICE_DATABASE} for mysql user : ${DATABASE_USER}"
mysql -u$DATABASE_USER -p < ${bin}/db_create.sql --default-character-set=utf8mb4

# upgrade database to version-1.0.1
echo "Next to upgrade database to version-1.0.1"
mysql -u$DATABASE_USER -p < ${bin}/db_upgrade_v1.sql --default-character-set=utf8mb4

# finished init database
echo "Finished init and upgrade database"
