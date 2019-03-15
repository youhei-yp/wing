#!/usr/bin/env bash

# Copyright (c) 2018-2019 Xunmo All Rights Reserved.
#
# Author : yangping
# Email  : youhei_yp@163.com
#
# Prismy.No | Date       | Modified by. | Description
# -------------------------------------------------------------------
# 00001       2018/12/01   youhei         New version
# -------------------------------------------------------------------

bin=`dirname "$0"`
bin=`cd "$bin"; pwd`

${bin}/daemon.sh stop $*

