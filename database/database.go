package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Conectar() {
	connStr := "user=postgres password=1234 dbname=produtos_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("✅ Conectado ao banco de dados com sucesso!")
	DB = db
}
