package services

import (
	"postapi/app/database"
	"postapi/app/models"
	"regexp"
)


func CreatePhone(country string, countryCode int64, phoneNumer int64) (*models.Phones, error) {

	regexCode, err := GetRegex(countryCode)
	if err != nil {
		return nil, err
	}

	phoneByte := []byte(string(phoneNumer))
	match, err := regexp.Match(regexCode.RegexCode, phoneByte)
	if err != nil {
		return nil, err
	}

	objPhone := models.Phones{
		Country:     country,
		State:       match,
		CountryCode: countryCode,
		PhoneNumber: phoneNumer,
	}

	err = database.Mgr.InsertPhone(&objPhone)
	if err != nil {
		return nil, err
	}
	return &objPhone, nil
}


func CapturePhone(country *string, state *bool) ([]models.Phones, error) {

	out, err := database.Mgr.SelectPhone(country, state)
	if err != nil {
		return nil, err
	}
	return out, err
}
