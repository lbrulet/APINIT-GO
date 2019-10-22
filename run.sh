#!/bin/bash

echo "[DB] Run migrations..."
migrate -source "file://src/" -database "mysql://root:root@tcp(mysql:3306)/apinit_go" up
echo "[DB] End of migrations..."
echo "[API] Start..."
./APINIT-GO