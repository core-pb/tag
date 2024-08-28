#!/bin/bash

if [ "$(basename `pwd`)" == "build" ]
then
  cd ..
fi

export CGO_ENABLED=0
export GOOS=linux

ldflags="-s -w"

go mod download

cd serer

go build -buildvcs=false -ldflags "$ldflags" -o ../build/output/server

chmod a+x ../build/output/*
