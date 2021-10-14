package services

import (
	"context"
	"postapi/app/database"
	"postapi/app/models"
)

//type Creating interface {
//	CreatePhone(ctx context.Context, country string, state bool, countryCode string, phoneNumer string) error
//}

func CreatePhone(ctx context.Context, country string, state bool, countryCode string, phoneNumer string) error {

	objPhone := models.Phones{
		Country:     country,
		State:       state,
		CountryCode: countryCode,
		PhoneNumber: phoneNumer,
	}

	_, err := database.GetCollection(models.COLLECTION_PHONES).InsertOne(ctx, objPhone)
	if err != nil {
		return err
	}
	return nil
}
