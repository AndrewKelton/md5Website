package main

import (
	"database/sql"
    "log"
    "os"
	"bufio"
	"strings"
	"fmt"
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