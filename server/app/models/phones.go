package models

const COLLECTION_PHONES = "teste"

type Phones struct {
	Country     string `db:"country"`
	State       bool   `db:"state"`
	CountryCode string `db:"countrycode"`
	PhoneNumber string `db:"phonenumber"`
}

type JsonPost struct {
	Country     string `json:"country"`
	State       bool   `json:"state"`
	CountryCode string `json:"countryCode"`
	PhoneNumber string `json:"phoneNumber"`
}

type PostRequest struct {
	Country     *string `json:"country" validate:"required"`
	State       *bool   `json:"state" validate:"required"`
	CountryCode *string `json:"countryCode" validate:"required"`
	PhoneNumber *string `json:"phoneNumber" validate:"required"`
}
