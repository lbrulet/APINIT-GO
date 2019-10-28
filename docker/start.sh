#!/usr/bin/env bash

FILE=/etc/init.d/mysql
if test -f "$FILE"; then
    echo "shutdown mysql..."
    /etc/init.d/mysql stop
fi
docker-compose up -d
