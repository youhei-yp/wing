@echo off

:: Copyright (c) 2019-2029 Dunyu All Rights Reserved.
::
:: Author      : yangping
:: Email       : ping.yang@wengold.net
:: Version     : 1.0.1
:: Description :
::   Start scenegreeting server.
::
:: Prismy.No | Date       | Modified by. | Description
:: -------------------------------------------------------------------
:: 00001       2020/05/08   yangping       New version
:: 00002       2020/08/16   yangping       Support for windows
:: -------------------------------------------------------------------

set BINPATH=%~dp0
call %BINPATH%\scripts\exports.bat

cd /d %BINPATH%\..
call %SERVICE_APP_NAME%.exe
echo started server...

