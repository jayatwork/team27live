package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "personnel"
)

func main() {
	password := os.Getenv("DB_PASS_ENV")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `
INSERT INTO personnel (id, implementation, account, model, status)
VALUES ($1, $2, $3, $4, $5)
RETURNING id`
	id := 0
	err = db.QueryRow(sqlStatement, 0, 1010, "alpaca", "ensemble", "200").Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)
}
