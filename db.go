package main

import (
	"database/sql"
	//"errors"
	//"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
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
            hours INTEGER,
            message TEXT,
            timestamp TIMESTAMP
        );
    `
	_, err = db.Exec(statement)
	if err != nil {
		log.Printf("%q: %s\n", err, statement)
		return
	}
}

func CreateCommit(commit Commit) {
	db, err := sql.Open("sqlite3", "./scribe.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	transaction, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	statement, err := transaction.Prepare(`
        INSERT INTO records (project, hours, message, timestamp)
        VALUES (?, ?, ?, ?);
    `)
	if err != nil {
		log.Fatal(err)
	}

	defer statement.Close()

	_, err = statement.Exec(commit.Project, commit.Hours, commit.Message, time.Now().UTC())
	if err != nil {
		log.Fatal(err)
	}

	transaction.Commit()
}
