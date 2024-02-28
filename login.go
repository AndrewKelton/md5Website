package main

import (
	"crypto/md5"
	"encoding/hex"
	"database/sql"
	"fmt"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

func Login() {
	// Open the SQLite database file
	db, err := sql.Open("sqlite3", "lamp.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var user string
	fmt.Print("Username: ")
	fmt.Scanf("%s", &user)
	// get password
	var pass string
	fmt.Print("Password: ")
	fmt.Scanf("%s", &pass)

	// Query the database to retrieve a single value
	query := "SELECT pass FROM users WHERE username = ?"

	// get password in db
	var password string
	err = db.QueryRow(query, user).Scan(&password)
	if err != nil {
		log.Fatal(err)
	}

	// get user password convert to md5
	h := md5.New()
	h.Write([]byte(pass))
	sum1 := h.Sum(nil)
	hS := hex.EncodeToString(sum1)

	// check if md5s are same
	if hS == password {
		fmt.Println("Login granted")
		return
	}
	fmt.Println("Incorrect password")
	return
}