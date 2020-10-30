#!/usr/bin/env bash

# Copyright (c) 2019-2029 Dunyu All Rights Reserved.
#
# Author      : yangping
# Email       : youhei_yp@163.com
# Version     : 1.0.2
# Description :
#   Deploy xxx server.
#
# Prismy.No | Date       | Modified by. | Description
# -------------------------------------------------------------------
# 00001       2020/05/08   yangping       New version
# 00002       2020/08/16   yangping       Support for windows
# -------------------------------------------------------------------

bin=`dirname "$0"`
bin=`cd "$bin"; pwd`
source ${bin}/exports.sh

cd ${bin}/../../
echo ">>> Building Project..."
go get
go build

# mdir target folders
echo ">>> Deploying Files..."
if [ -d "./${SERVICE_DEPLOY_DIR}/" ];then
  rm -r ${SERVICE_DEPLOY_DIR}
fi
mkdir ${SERVICE_DEPLOY_DIR}
cd ${SERVICE_DEPLOY_DIR}
mkdir -p logs public media
cd ..

# copy binary files to target folders
cp -rf ${SERVICE_APP_NAME} ./${SERVICE_DEPLOY_DIR}/
cp -rf bin     ./${SERVICE_DEPLOY_DIR}/bin/
cp -rf conf    ./${SERVICE_DEPLOY_DIR}/conf/
cp -rf swagger ./${SERVICE_DEPLOY_DIR}/swagger/

# mv
echo ">>> Move Service To Target Deploy Path"
if [ -d "${SERVICE_DEPLOY_WEB}/${SERVICE_DEPLOY_DIR}" ];then
  rm -r ${SERVICE_DEPLOY_WEB}/${SERVICE_DEPLOY_DIR}
fi
mv ./${SERVICE_DEPLOY_DIR} ${SERVICE_DEPLOY_WEB}/

# change file execute permission
cd ${SERVICE_DEPLOY_WEB}/${SERVICE_DEPLOY_DIR}
chmod -R 755 .

# finished deploy
echo ">>> -------------------"
echo ">>> Finised Deploy ${SERVICE_APP_NAME} Service To:"
echo ">>> ${SERVICE_DEPLOY_WEB}"