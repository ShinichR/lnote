#!/bin/sh

mongod --dbpath /home/lei/dbs/master --port 10000 --master &
mongod --dbpath /home/lei/dbs/slave --port 10001 --slave --source localhost:10000 &
go run main.go
