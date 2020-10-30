@echo off

:: Copyright (c) 2019-2029 Dunyu All Rights Reserved.
::
:: Author      : yangping
:: Email       : youhei_yp@163.com
:: Version     : 1.0.1
:: Description :
::   Create database for xxx server.
::
:: Prismy.No | Date       | Modified by. | Description
:: -------------------------------------------------------------------
:: 00001       2020/05/08   yangping       New version
:: 00002       2020/08/16   yangping       Support for windows
:: -------------------------------------------------------------------

set BINPATH=%~dp0
set DATABASE_USER=%1
call %BINPATH%\exports.bat

:: check database user and password input data
if "%DATABASE_USER%"=="" goto ERROR

:: init database from .sql file
mysql -u%DATABASE_USER% -p < %BINPATH%\db_create.sql --default-character-set=utf8mb4
echo Inited database %SERVICE_DATABASE% for mysql user : %DATABASE_USER%

:: upgrade database to version-1
echo Next to upgrade database to version-1
mysql -u%DATABASE_USER% -p < %BINPATH%\db_upgrade_v1.sql --default-character-set=utf8mb4
goto END

:: print out error or end messages
:ERROR
echo Usage: initdb.bat (database user)
pause
exit 0

:: exit script
:END
exit 0