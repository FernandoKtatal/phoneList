package models

const COLLECTION_PHONES = "teste"

type Phones struct {
	Country     string `db:"country"`
	State       bool   `db:"state"`
	CountryCode int64 `db:"countrycode"`
	PhoneNumber int64 `db:"phonenumber"`
}

type JsonPost struct {
	Country     string `json:"country"`
	State       bool   `json:"state"`
	CountryCode int64 `json:"countryCode"`
	PhoneNumber int64 `json:"phoneNumber"`
}

type PostRequest struct {
	Country     *string `json:"country" validate:"required"`
	CountryCode *int64 `json:"countryCode" validate:"required"`
	PhoneNumber *int64 `json:"phoneNumber" validate:"required"`
}
