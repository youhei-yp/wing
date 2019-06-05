#!/usr/bin/env bash

go build -i -o ./tools ./src/tools/main.go

sudo chown -R youhei:youhei ./
sudo chmod -R 755 ./
