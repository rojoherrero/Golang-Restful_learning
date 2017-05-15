package services

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // This is the MariaDB driver
)

func connectToMariaDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@/golang_test")
	if err != nil {
		panic(err.Error)
	}
	log.Println("Connection to MariaDB created")
	return db
}
