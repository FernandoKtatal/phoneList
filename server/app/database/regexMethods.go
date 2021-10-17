package database

import (
	"postapi/app/models"
)

func (d DB) SelectRegex(countryCode int64) (*models.Regex, error) {
	var out models.Regex
	rows := d.db.QueryRow(string(selectRegexByCountryCode), countryCode)

	err := rows.Scan(&out.CountryCode, &out.RegexCode)
	if err != nil {
		return nil, err
	}

	return &out, err
}
