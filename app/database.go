package app

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Postgres struct {
	DB *sql.DB
}

func NewPostgresDatabase(env *Env) Postgres {
	db, err := sql.Open("postgres", fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		env.DBUser,
		env.DBPass,
		env.DBHost,
		env.DBPort,
		env.DBName,
	))
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return Postgres{DB: db}
}

func (db *Postgres) ClosePostgresDBConnection() {
	if db == nil {
		return
	}

	err := db.DB.Close()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to PostgresDB closed")
}
