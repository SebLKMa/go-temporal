#!/bin/bash

echo "Usage: <username> <password> <dbname> <sql script file>"

dbpass=$1
dbuser=$2
dbname=$3
sqlscript=$4

echo "dbpass:$dbpass dbuser:$dbuser dbname:$dbname sqlscript:$sqlscript"

# Create tables
PGPASSWORD=$dbpass psql -h 'localhost' -U $dbuser -d $dbname -f $sqlscript