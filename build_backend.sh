#!/bin/bash
apt-get update && apt-get install -y libglib2.0-dev curl make
make build_init
find server/plugin/plg_* -type f -name '*.a' -exec mv {} /usr/local/lib/ \;
go get -v -t ./server/...
make build_backend
timeout 1 ./dist/filestash || true
