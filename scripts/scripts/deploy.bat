@echo off

:: Copyright (c) 2019-2029 Dunyu All Rights Reserved.
::
:: Author      : yangping
:: Email       : youhei_yp@163.com
:: Version     : 1.0.2
:: Description :
::   Deploy xxx server.
::
:: Prismy.No | Date       | Modified by. | Description
:: -------------------------------------------------------------------
:: 00001       2020/05/08   yangping       New version
:: 00002       2020/08/16   yangping       Support for windows
:: -------------------------------------------------------------------

set BINPATH=%~dp0
call %BINPATH%\exports.bat

echo ">>> Building Project..."
go get
go build

:: mkdir target folders
echo ">>> Deploying Files..."
cd /d %BINPATH%\..\..
IF EXIST %SERVICE_DEPLOY_DIR% (
  RD /S/Q %SERVICE_DEPLOY_DIR%
)
MKDIR %SERVICE_DEPLOY_DIR%
cd %SERVICE_DEPLOY_DIR%
MKDIR logs public media
cd ..

:: copy binary files to target folders
COPY  %SERVICE_APP_NAME%.exe %SERVICE_DEPLOY_DIR%\
XCOPY /E/Y bin     %SERVICE_DEPLOY_DIR%\bin\
XCOPY /E/Y conf    %SERVICE_DEPLOY_DIR%\conf\
XCOPY /E/Y swagger %SERVICE_DEPLOY_DIR%\swagger\

:: copy web static pagers
echo ">>> Copying Web Pages..."
IF "%SERVICE_DEPLOY_WEB%" == "" (
  echo !!! Abort update the web static pages !!!
) ELSE (
  XCOPY /E/Y %SERVICE_DEPLOY_WEB% %SERVICE_DEPLOY_DIR%\public\
)

:: finished deploy
echo ">>> -------------------"
echo ">>> Finised Deploy %SERVICE_APP_NAME% Service To:"
echo ">>> %BINPATH%..\..\%SERVICE_DEPLOY_DIR%"