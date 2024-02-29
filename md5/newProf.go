package main

import (
	"crypto/md5"
	"database/sql"
	//"fmt"
	"encoding/hex"
	"log" 
	_ "github.com/mattn/go-sqlite3"
)

func Newpro(username string, password string) bool {

	// open db
	db, err := sql.Open("sqlite3", "lamp.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// create table
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (\nusername TEXT NOT NULL,\npass TEXT NOT NULL\n)")

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Print("Username: ")
	// fmt.Scanf("%s", &username)
	// fmt.Print("Password: ")
	// fmt.Scanf("%s", &password)

	// encode password to md5
	h := md5.New()
	h.Write([]byte(password))
	sum1 := h.Sum(nil)
	hS := hex.EncodeToString(sum1)

	// add to db
	_, err = db.Exec("INSERT INTO users (username, pass) VALUES (?, ?)", username, hS)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}