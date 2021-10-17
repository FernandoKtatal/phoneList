package services

import (
	"postapi/app/database"
	"postapi/app/models"
)

func GetRegex(countryCode int64) (*models.Regex, error) {

	objRegex, err := database.Mgr.SelectRegex(countryCode)
	if err != nil {
		return nil, err
	}

	return objRegex, nil
}

