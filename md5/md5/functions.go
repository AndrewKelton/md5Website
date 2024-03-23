package main

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	_ "github.com/mattn/go-sqlite3"
)

func Table() {
    // Open the HTML file
    file, err := os.Open("frontpage/admin.html")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Create a scanner to read the file line by line
    scanner := bufio.NewScanner(file)
    var lines []string

    // Read the file and remove the table content
    var inTable bool
    for scanner.Scan() {
        line := scanner.Text()

        // Check if we are inside the table
        if strings.Contains(line, "<table") {
            inTable = true
        } else if strings.Contains(line, "</table>") {
            inTable = false
            continue
        }

        // Skip lines inside the table
        if inTable {
            continue
        }

        // Add lines outside the table to the slice
        lines = append(lines, line)
    }

    // Check for errors in scanner
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    // Reopen the file for writing (truncate the existing content)
    file, err = os.OpenFile("frontpage/admin.html", os.O_WRONLY|os.O_TRUNC, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Write the modified lines back to the file
    writer := bufio.NewWriter(file)
    for _, line := range lines {
        _, err := writer.WriteString(line + "\n")
        if err != nil {
            log.Fatal(err)
        }
    }

    // Flush the buffer to ensure all data is written to the file
    if err := writer.Flush(); err != nil {
        log.Fatal(err)
    }

    log.Println("Table deleted from existing HTML file successfully!")
	// Open the SQLite database file
	db, err := sql.Open("sqlite3", "lamp.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Query the database
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Open the existing HTML file in append mode
	file, err = os.OpenFile("frontpage/admin.html", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "<html><head><title>Table</title></head><body><table border='1'><tr><th>Username</th><th>Password</th></tr>")
    if err != nil {
        log.Fatal(err)
    }

	// Write the HTML table rows to the file
	for rows.Next() {
		var username string
		var password string
		if err := rows.Scan(&username, &password); err != nil {
			log.Fatal(err)
		}
		_, err := fmt.Fprintf(file, "<tr><td>%s</td><td>%s</td></tr>", username, password)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Table appended to existing HTML file successfully!")
}

func Login(user string, pass string) bool {
	// Open the SQLite database file
	db, err := sql.Open("sqlite3", "lamp.db")
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer db.Close()

	// Query the database to retrieve a single value
	query := "SELECT pass FROM users WHERE username = ?"

	// get password in db
	var password string
	err = db.QueryRow(query, user).Scan(&password)
	if err != nil {
		if err == sql.ErrNoRows {
			// User not found
			fmt.Println("User not found")
			return false
		} else {
			// Other error occurred
			fmt.Println("Error:", err)
		}
	}

	// get user password convert to md5
	h := md5.New()
	h.Write([]byte(pass))
	sum1 := h.Sum(nil)
	hS := hex.EncodeToString(sum1)

	// check if md5s are same
	if hS == password {
		fmt.Println("Login granted")
		return true
	} else {
		fmt.Println("Incorrect password")
		return false
	}	
	return false
}

func Delete(username string) bool {

	// Open the SQLite database file
	db, err := sql.Open("sqlite3", "lamp.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := "DELETE FROM users WHERE username = ?"
	result, err := db.Exec(query, username)
	if err != nil {
		log.Fatal(err)
		return false
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
		return false
	}
	if rowsAffected > 0 {
		fmt.Println(rowsAffected)
		return true
	} else {
		return false
	}
}

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

func CreateCap() string {
	cmd := exec.Command("python3", "rdm.py")

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
    if err != nil {
        fmt.Printf("Error: %s\n", err)
        return ""
    }

	pythonOutput := strings.TrimSpace(stdout.String())
    result := pythonOutput
    fmt.Printf("Result from Python: %s\n", result)
	return result
}