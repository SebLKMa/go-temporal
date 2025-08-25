#!/bin/bash

echo "Usage: <username> <password> <dbname> <sql script file>"
echo "Examples:"
echo "./managetables.sh pdtester pdtester pddbtest tables_<...>.sql"
echo "./managetables.sh pduser pduser pddb tables_<...>.sql"
echo "./managetables.sh pduser pduser pddb tables_surl.sql"
echo "Login to check your tables..."
#echo psql \"host=localhost port=5432 dbname=pddb user=pduser password=pduser\"
#echo psql \"host=localhost port=5432 dbname=pddbtest user=pdtester password=pdtester\"

dbpass=$1
dbuser=$2
dbname=$3
sqlscript=$4

echo "dbpass:$dbpass dbuser:$dbuser dbname:$dbname sqlscript:$sqlscript"

# Create tables
#PGPASSWORD=pduser psql -h 'localhost' -U 'pduser' -d 'pddb' -f $sqlscript
PGPASSWORD=$dbpass psql -h 'localhost' -U $dbuser -d $dbname -f $sqlscript