package database

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"postapi/app/models"
)

type PostDB interface {
	Connect() error
	Disconnect() error
	InsertPhone(p *models.Phones) error
	SelectPhone(country *string, valid *bool) ([]models.Phones, error)
	SelectRegex(countryCode int64) (models.Regex, error)
}

type DB struct {
	db *sql.DB
}

var Mgr DB


func Connect() error {
	db, err := sql.Open("sqlite3", "../sample.db")
	if err != nil {
		return err
	}
	Mgr = DB{db: db}

	return execMigrations(db)
}

func Disconnect() error {
	return Mgr.db.Close()
}

func GetDB() DB {
	return Mgr
}

func execMigrations(db *sql.DB) error {
	files, err := os.ReadDir(MigrationPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		query, err := os.ReadFile(MigrationPath + file.Name())
		if err != nil || query == nil {
			return errors.New(PathQueriesIndevido)
		}

		migration, err := db.Prepare(string(query))
		if err != nil {
			return err
		}
		migration.Exec()
	}
	return err
}
