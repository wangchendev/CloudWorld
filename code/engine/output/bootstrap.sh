#!/bin/bash

CURDIR=$(cd $(dirname $0); pwd)
CONF_FILE=$CURDIR/conf/server_config.yaml

echo "$CONF_FILE"
go run $CURDIR/../main.go -conf="$CONF_FILE"