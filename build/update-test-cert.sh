#!/bin/bash

if [ "$(basename `pwd`)" == "build" ]
then
  cd ..
fi

curl -o build/output/cert.crt https://core-pb.github.io/testcert/server.crt

curl -o build/output/cert.key https://core-pb.github.io/testcert/server.key
