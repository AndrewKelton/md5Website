#!/bin/bash

export GO111MODULE=on
go get github.com/mattn/go-sqlite3
go run main.go functions.go