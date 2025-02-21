package database

import (
    "database/sql"
    "fmt"
    "log"
    _ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
    connStr := "host=localhost user=postgres password=Admin123 dbname=task_manager sslmode=disable"
    var err error
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }

    if err := DB.Ping(); err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected to PostgreSQL!")
}
