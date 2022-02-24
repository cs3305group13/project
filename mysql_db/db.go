package mysql_db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/cs3305group13/project/mysql_db/utils"
	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	HOSTNAME string
	DBNAME string
	PORT string
	USERNAME string
	PASSWORD string
}

// Creates a struct with hostname, dbname, port, username, password string fields.
func NewDB(envs map[string]string) *DB {

	hostname := envs["HOSTNAME"]
	dbname := envs["DBNAME"]
	port := envs["PORT"]
	username := envs["USERNAME"]
	password := envs["PASSWORD"]

	return &DB{hostname, dbname, port, username, password}
}

func dsn(db *DB) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", db.USERNAME, db.PASSWORD, db.HOSTNAME, db.PORT, db.DBNAME)
}

// Function establishes a connection with the database using the DB struct provided.
func EstablishConnection(db *DB) *sql.DB {
	// establish connection with database.
	sqlDB, err := sql.Open("mysql", dsn(db))
	if err != nil {
		log.Fatal(err)
	}

	return sqlDB
}


func NewTransaction(db *sql.DB) *sql.Tx {

	// Get a Tx for making transaction requests.
	ctx := context.Background()
    tx, err := db.BeginTx(ctx, nil)
	
    utils.CheckError(err)

	return tx
}