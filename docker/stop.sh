#!/usr/bin/env bash

docker-compose -f docker-compose.yml stop -t 1
FILE=/etc/init.d/mysql
if test -f "$FILE"; then
    echo "run mysql..."
    /etc/init.d/mysql start
fi
