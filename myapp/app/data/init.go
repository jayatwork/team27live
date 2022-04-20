package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

func Initdb() {
	// Being set at the systemOS level with export DATABASE_URL="postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	var (
		implementation int64
		account        string
		status         string
	)
	err = conn.QueryRow(context.Background(), "select implementation, account, status from trading where id=$1", 42).Scan(&implementation, &account, &status)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(implementation, account, status)
}

func main() {
	fmt.Println("Program exited.")
}
