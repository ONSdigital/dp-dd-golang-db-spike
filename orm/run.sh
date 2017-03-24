#!/usr/bin/env bash

export DB_USER=
export DB_PWD=
export DB_NAME=

echo
echo "> Compiling 'orm'"
go build -o ../lib/orm
echo "> Compilation completed"

echo
echo "> Starting application:"
.././lib/orm --db_user=$DB_USER --db_pwd=$DB_PWD --db_name=$DB_NAME