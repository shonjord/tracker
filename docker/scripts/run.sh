#!/bin/bash

sleep 15

./docker/scripts/migrate.sh

CompileDaemon --build='make install' --command=tracker
