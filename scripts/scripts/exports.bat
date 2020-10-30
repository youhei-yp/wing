@echo off

:: Copyright (c) 2019-2029 Dunyu All Rights Reserved.
::
:: Author      : yangping
:: Email       : youhei_yp@163.com
:: Version     : 1.0.1
:: Description :
::   Export the temp params for bat script.
::
:: Prismy.No | Date       | Modified by. | Description
:: -------------------------------------------------------------------
:: 00001       2020/05/08   yangping       New version
:: 00002       2020/08/16   yangping       Support for windows
:: -------------------------------------------------------------------

set SERVICE_APP_NAME=xxx
set SERVICE_DATABASE=xxxdb

:: use for auto start when system boot completed,
:: and start browser to load the follow web host.
set SERVICE_HOST_PORT=192.168.0.1:6100

:: use for output the deploy binarys.
set SERVICE_DEPLOY_DIR=xxxx

:: FOR WINDOWS : use for copy the web pages to deploy folder.
:: FOR LINUX   : use for delopy the service to target path.
set SERVICE_DEPLOY_WEB=E:\01_Webpace\xxx\build
