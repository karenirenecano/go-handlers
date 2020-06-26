package database

import (
	"database/sql"
	"log"
)

//DbConn globally declared
var DbConn *sql.DB

//SetupDatabase for referencing db connection
func SetupDatabase() {
	var err error
	DbConn, err = sql.Open("mysql", "root:@tcp(127.0.0.1)/inventorydb")
	if err != nil {
		log.Fatal(err)
	}
}
