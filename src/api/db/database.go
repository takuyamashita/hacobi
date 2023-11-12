package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

type database sql.DB

func NewDatabase() *sql.DB {

	db := open(getConfigParams())
	if !isCoccencted(db) {
		log.Fatal("Could not connect to the database")
	}

	return db
}

func getConfigParams() mysql.Config {

	DBName := os.Getenv("DB_NAME")
	User := os.Getenv("DB_USER")
	Passwd := os.Getenv("DB_PASSWORD")
	Addr := fmt.Sprintf("%s:3306", os.Getenv("DB_HOST"))

	return mysql.Config{
		DBName:    DBName,
		User:      User,
		Passwd:    Passwd,
		Addr:      Addr,
		Collation: "utf8mb4_unicode_ci",
		Net:       "tcp",
	}
}

func open(config mysql.Config) *sql.DB {

	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func isCoccencted(db *sql.DB) bool {

	for i := 0; i < 5; i++ {
		if err := db.Ping(); err == nil {
			return true
		}

		time.Sleep(1 * time.Second)
	}

	return false
}
