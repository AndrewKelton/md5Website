package main

import (
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"fmt"
	"log"
)

func Delete() {

	// Open the SQLite database file
	db, err := sql.Open("sqlite3", "lamp.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var username string
	fmt.Print("Username to delete: ")
	fmt.Scanf("%s", &username)

	query := "DELETE FROM users WHERE username = ?"
	result, err := db.Exec(query, username) // Replace 123 with the ID or condition you want to delete
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rowsAffected)
}