package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

type database sql.DB

func NewDatabase(ctx context.Context) *sql.DB {

	db := open(getConfigParams())
	if err := isCoccencted(ctx, db); err != nil {
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
		ParseTime: true,
	}
}

func open(config mysql.Config) *sql.DB {

	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func isCoccencted(ctx context.Context, db *sql.DB) error {

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	db.PingContext(ctx)

	return nil
}
