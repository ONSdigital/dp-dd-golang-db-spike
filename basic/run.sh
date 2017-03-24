#!/usr/bin/env bash

export DB_USER=
export DB_PWD=
export DB_NAME=

echo
echo "> Compiling..."
go build -o ../lib/spike
echo "> Compilation completed"

echo
echo "> Starting application:"
.././lib/spike --db_user=$DB_USER --db_pwd=$DB_PWD --db_name=$DB_NAME