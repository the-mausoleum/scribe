package main

import (
	"database/sql"
	//"errors"
	//"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func Connect() {
	db, err := sql.Open("sqlite3", "./scribe.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var statement string = `
        CREATE TABLE IF NOT EXISTS records (
            id INTEGER PRIMARY KEY,
            project TEXT,
            hours INTEGER
            timestamp TIMESTAMP
        );
    `
	_, err = db.Exec(statement)
	if err != nil {
		log.Printf("%q: %s\n", err, statement)
		return
	}
}

func CreateCommit(comm Commit) {

}
