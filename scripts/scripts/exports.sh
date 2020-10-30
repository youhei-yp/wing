#!/usr/bin/env bash

# Copyright (c) 2019-2029 Dunyu All Rights Reserved.
#
# Author      : yangping
# Email       : youhei_yp@163.com
# Version     : 1.0.1
# Description :
#   Export the temp params for sh script.
#
# Prismy.No | Date       | Modified by. | Description
# -------------------------------------------------------------------
# 00001       2020/05/08   yangping       New version
# 00002       2020/08/16   yangping       Support for windows
# -------------------------------------------------------------------

export SERVICE_APP_NAME=xxx
export SERVICE_DATABASE=xxxdb

# use for auto start when system boot completed,
# and start browser to load the follow web host.
export SERVICE_HOST_PORT=192.168.0.1:6100

# use for output the deploy binarys.
export SERVICE_DEPLOY_DIR=xxxxx

# FOR WINDOWS : use for copy the web pages to deploy folder.
# FOR LINUX   : use for delopy the service to target path.
export SERVICE_DEPLOY_WEB=~/ws/env_xxx/