package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"postapi/app/models"
)

type PostDB interface {
	Connect() error
	Disconnect() error
	Insert(p *models.Phones) error
	Select(country *string, valid *bool) ([]models.Phones, error)
}

type DB struct {
	db *sql.DB
}

func (d *DB) Connect() error {
	db, err := sql.Open("sqlite3", "../sample.db")
	if err != nil {
		return err
	}
	d.db = db

	files, err := os.ReadDir(MigrationPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		query, err := os.ReadFile(MigrationPath + file.Name())
		if err != nil || query == nil {
			return err
		}

		migration, err := db.Prepare(string(query))
		if err != nil {
			return err
		}
		migration.Exec()
	}


	return nil
}

func (d *DB) Disconnect() error {
	return d.db.Close()
}
