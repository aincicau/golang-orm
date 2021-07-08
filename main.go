package main

import "golangorm/db"

func main() {
	db.InitDatabase("postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable")
}
