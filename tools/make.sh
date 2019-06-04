#!/usr/bin/env bash

go build -i -o ./genrsakeys ./src/genrsakey/main.go
go build -i -o ./genaeskeys ./src/genaeskey/main.go
go build -i -o ./genuuids   ./src/genuuid/main.go

sudo chown -R youhei:youhei ./
sudo chmod -R 755 ./
