package main

import (
	"database/sql"
	"fmt"

	_ "pq-master"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "DBa"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	fmt.Printf("psqlInfo", psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	//defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	defer db.Close()
}
