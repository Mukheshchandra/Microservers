package driver

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // connection to driver
)

func Connection() *sql.DB {
	db, err := sql.Open("mysql", "root:mukheshM1@25-03@/testingDB")
	if err != nil {
		log.Print("failed to connect with database", err)
	}

	return db
}
