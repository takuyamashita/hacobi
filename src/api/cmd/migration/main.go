package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/takuyamashita/hacobi/src/api/db"
)

const (
	MigrationFileDir = "migration"
	VersionDigit     = 8
)

func main() {
	var (
		upFlag = flag.Bool("up", false, "up")
	)
	flag.Parse()

	switch {
	case *upFlag:
		up()
	default:
		log.Fatal("invalid flag")
	}
}

func up() {

	db := db.NewDatabase(context.Background())
	defer db.Close()

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal(err)
	}
	wd, err := os.Getwd()
	file := fmt.Sprintf("file://%s/%s", wd, MigrationFileDir)
	fmt.Println(file)
	m, err := migrate.NewWithDatabaseInstance("file://migration", "mysql", driver)
	if err != nil {
		log.Fatal(err, file)
	}

	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}
