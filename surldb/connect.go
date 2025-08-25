package surldb

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq" // for sql.Open("postgres", psqlInfo)
)

type DbConnection struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

func NewConnection(host string, port int, user string, pw string, dbname string) DbConnection {
	return DbConnection{Host: host, Port: port, User: user, Password: pw, DbName: dbname}
}

// Connect opens and return the db object. Caller must close db when done with it.
func (conn DbConnection) Connect() (db *sql.DB) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conn.Host, conn.Port, conn.User, conn.Password, conn.DbName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return db
}

// Connect opens and return the db object. Caller must close db when done with it.
func (conn DbConnection) TryConnect() (db *sql.DB, err error) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conn.Host, conn.Port, conn.User, conn.Password, conn.DbName)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// GetConnectionFromEnv returns a DbConnection based on environment variables:
//
//	DBHOST - the hostname or IP where db is running.
//	         db may not reside in application docker, therefore must not be localhost for docker.
//	DBPORT - usually 5432 for postgres0
//	DBNAME - name of database to connect
//	DBUSER - username
//	DBPASS - password
func GetConnectionFromEnv() (dbconn DbConnection, err error) {
	// NOTE: db may not reside in application docker, therefore must not be localhost for docker.
	//dbconn = NewConnection("localhost", 5432, "pduser", "pduser", "pddb")
	dbconn = DbConnection{}

	// https://stackoverflow.com/questions/31249112/allow-docker-container-to-connect-to-a-local-host-postgres-database
	// From env vars (e.g. docker --env --env-file)

	// Update postgres config files:
	// https://www.postgresql.org/docs/current/auth-pg-hba-conf.html
	/*
	   $ sudo nano /etc/postgresql/14/main/postgresql.conf
	   #listen_addresses = 'localhost'         # what IP address(es) to listen on;
	                                           # comma-separated list of addresses;
	                                           # defaults to 'localhost'; use '*' for all
	                                           # (change requires restart)
	   listen_addresses = '*'

	   $ sudo nano /etc/postgresql/14/main/pg_hba.conf
	   # IPv4 local connections:
	   host    all             all             127.0.0.1/32            scram-sha-256
	   host    all             all             172.17.0.0/16           password

	   $ sudo service postgresql restart
	*/
	dbhost := os.Getenv("DBHOST") // the hostname or IP where db is running
	dbport := os.Getenv("DBPORT") // 5432
	dbname := os.Getenv("DBNAME")
	dbuser := os.Getenv("DBUSER")
	dbpass := os.Getenv("DBPASS")

	if dbhost != "" && dbport != "" && dbname != "" && dbuser != "" && dbpass != "" {
		port, err := strconv.Atoi(dbport)
		if err == nil {
			dbconn = NewConnection(dbhost, port, dbuser, dbpass, dbname)
		}
	}
	if (DbConnection{}) == dbconn {
		return dbconn, errors.New("Error getting DB connection from environment variables.")
	}
	_, err = dbconn.TryConnect()
	if err != nil {
		return dbconn, err
	}
	return dbconn, nil
}
