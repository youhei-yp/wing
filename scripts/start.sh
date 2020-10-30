#!/usr/bin/env bash

# Copyright (c) 2019-2029 Dunyu All Rights Reserved.
#
# Author      : yangping
# Email       : ping.yang@wengold.net
# Version     : 1.0.1
# Description :
#   Start scenegreeting server.
#
# Prismy.No | Date       | Modified by. | Description
# -------------------------------------------------------------------
# 00001       2020/05/08   yangping       New version
# 00002       2020/08/16   yangping       Support for windows
# -------------------------------------------------------------------

bin=`dirname "$0"`
bin=`cd "$bin"; pwd`

source ${bin}/scripts/exports.sh
${bin}/scripts/daemon.sh start ${SERVICE_APP_NAME}

