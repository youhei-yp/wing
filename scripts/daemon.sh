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

show_usage() {
  echo "Usage: daemon.sh [start|stop|status]"
  echo "options as :"
  echo "  start          - start server"
  echo "  stop [-f]      - stop server, use -f to force kill server"
  echo "       [--force] - as use -f to kill server"
  echo "  status         - show process id and running status"
  echo "  "
}

# if no args specified, show usage
if [ $# -le 0 ]; then
  echo show_usage
  exit 1
fi

# enter root path and set to ROOT
bin=`dirname "$0"`
bin=`cd "$bin"; pwd`
ROOT=`cd "${bin}/.."; pwd`

# ------------------------------------------------------- #
# TODO :                                                  #
#                                                         #
# YOU MUST CHANGE XX_SERVER_HOME TO YOUR SERVER HOME NAME #
# e.g: XX_SERVER_HOME -> PASTER_SERVER_HOME               #
# ------------------------------------------------------- #
# export server home path
[[ ! -d ${XX_SERVER_HOME} ]] && XX_SERVER_HOME="${ROOT}"
export XX_SERVER_HOME=${XX_SERVER_HOME}

# ------------------------------------------------------- #
# TODO :                                                  #
#                                                         #
# YOU MUST CHANGE server_name_sample TO YOUR SERVER NAME  #
# e.g: DAEMON_SERVER_NAME=paster                          #
# ------------------------------------------------------- #
DAEMON_SERVER_NAME=server_name_sample

# get start|stop|status arguments
commands=$1
shift
umask 0000

# get force kill argument
forced=$1
FORCED_KILL=false
shopt -s extglob
if [ ! -z ${forced} ]; then
  case ${forced} in
    *(-)f )
      FORCED_KILL=true
      ;;
    *(-)force )
      FORCED_KILL=true
      ;;
  esac
fi

# create log directory and process file
LogDir="${XX_SERVER_HOME}/logs"
mkdir -p "$LogDir"
pid=${LogDir}/${DAEMON_SERVER_NAME}.pid

# dispath start|stop|status commands
case $commands in

  (start)
    if [ -f $pid ]; then
      if ps -p `cat $pid` > /dev/null 2>&1; then
        echo server running as process `cat $pid`.
        exit 1
      fi
    fi

    cd ${XX_SERVER_HOME}
    ./${DAEMON_SERVER_NAME} >&1 </dev/null &
    echo $! > ${pid}
    ;;

  (stop)
    if [ -f $pid ]; then
      if ps -p `cat $pid` > /dev/null 2>&1; then
        if ! kill -0 `cat $pid` > /dev/null 2>&1; then
          echo cannot stop server on pid `cat $pid` - permission denied
        elif ${FORCED_KILL}; then
          kill -9 `cat $pid` > /dev/null 2>&1;
          sleep 1;
          echo server killed
        else
          echo -n stopping server
          kill `cat $pid` > /dev/null 2>&1
          while ps -p `cat $pid` > /dev/null 2>&1; do
            echo -n "."
            sleep 1;
          done
          echo
          echo stoped server
        fi
        if ! ps -p `cat $pid` > /dev/null 2>&1; then
          rm $pid
        fi
      else
        echo server not running
      fi
    else
      echo server not running
    fi
    ;;

  (status)
    if [ -f $pid ]; then
      if ps -p `cat $pid` > /dev/null 2>&1; then
        echo server running on process `cat $pid`.
      else
        echo server not running.
      fi
    else
      echo server not running.
    fi
    ;;

  (*)
    echo $usage
    exit 1
    ;;

esac
