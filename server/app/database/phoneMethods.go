package database

import (
	"database/sql"
	"errors"
	"postapi/app/models"
)

func (d *DB) InsertPhone(p *models.Phones) error {
	query, err := d.db.Prepare(insertPostPhones)
	if err != nil {
		return err
	}
	_, err = query.Exec(p.Country, p.State, p.CountryCode, p.PhoneNumber)
	if err != nil {
		return err
	}

	return nil
}

func (d *DB) SelectPhone(country *string, state *bool) (out []models.Phones, err error) {
	var item models.Phones
	var rows *sql.Rows
	if country != nil && state != nil {
		rows, _ = d.db.Query(string(selectPhonesWith2Filters), country, state)
	} else if country != nil {
		rows, _ = d.db.Query(string(selectPhonesCountryFilter), country)
	} else if state != nil {
		rows, _ = d.db.Query(string(selectPhonesStateFilter), state)
	} else {
		rows, _ = d.db.Query(string(selectPhonesWithNoFilters))
	}

	if rows == nil {
		return nil, errors.New("Nenhum item encontrado")
	}

	for rows.Next() {
		rows.Scan(&item.Country, &item.State, &item.CountryCode, &item.PhoneNumber)
		out = append(out, item)
	}

	defer rows.Close()

	return out, err
}
