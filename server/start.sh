#!/bin/sh

./linux-build.sh
cd bin
chmod +x ./daigouadmin
./daigouadmin

# nohup ./daigouadmin >> nohup.log 2>&1 &