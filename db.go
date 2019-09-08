package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

const dbPath = "./test.db"
const cGrpTbl = `CREATE TABLE GROUP (
id integer not null primary key,
name text,
created_at datetime
updated_at datetime
`

func createDatebase() {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal("While creating db:", err)
	}
	defer db.Close()

	if _, err := db.Exec(cGrpTbl); err != nil {
		log.Println("Created group table")
	}
}

func BootstrapDB() {
	if _, err := os.Stat(dbPath); err != nil {
		fmt.Println("Database doesn't exist!")
		createDatebase()
	} else {
		fmt.Println("Database already exists!")
	}
}
