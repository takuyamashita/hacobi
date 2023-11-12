package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

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
		upFlag         = flag.Bool("up", false, "up")
		createFileName = flag.String("create", "", "create")
	)
	flag.Parse()

	switch {
	case *upFlag:
		up()
	case *createFileName != "":
		create(*createFileName)
	default:
		log.Fatal("invalid flag")
	}
}

func up() {

	db := db.NewDatabase()
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

func create(fileTitle string) {

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	files, err := ioutil.ReadDir(fmt.Sprintf("%s/%s", wd, MigrationFileDir))
	if err != nil {
		log.Fatal(err)
	}

	if len(files) == 0 {
		createMigrationFile(1, fileTitle)
		return
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	lastFile := files[len(files)-1].Name()
	lastVersion, err := strconv.Atoi(strings.Split(lastFile, "_")[0])
	if err != nil {
		log.Fatal(err)
	}

	createMigrationFile(uint(lastVersion+1), fileTitle)
	return
}

func createMigrationFile(version uint, title string) {

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	formatText := fmt.Sprintf("%%0%dd_%%s.%%s.%%s", VersionDigit)
	nextVersionFileNameUp := fmt.Sprintf(formatText, version, title, "up", "sql")
	nextVersionFileNameDown := fmt.Sprintf(formatText, version, title, "down", "sql")
	os.OpenFile(fmt.Sprintf("%s/%s/%s", wd, MigrationFileDir, nextVersionFileNameUp), os.O_CREATE, 0666)
	os.OpenFile(fmt.Sprintf("%s/%s/%s", wd, MigrationFileDir, nextVersionFileNameDown), os.O_CREATE, 0666)

	fmt.Println(nextVersionFileNameUp)
	fmt.Println(nextVersionFileNameDown)
}
