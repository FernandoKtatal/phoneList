package test

import (
	"github.com/go-playground/assert/v2"
	"postapi/app/database"
	"postapi/app/models"
	"postapi/app/services"
	"testing"
)

var expected = models.Phones{
	Country:     "Mozambique",
	State:       true,
	CountryCode: 258,
	PhoneNumber: 21123456,
}

 var _ = database.Connect()

func TestGetPhones(t *testing.T) {

	out, _ := services.CapturePhone("Mozambique", nil)
	assert.Equal(t, out, expected)
}

func TestPostPhones(t *testing.T) {
	out, _ := services.CreatePhone(expected.Country, expected.CountryCode, expected.PhoneNumber)
	assert.Equal(t, out, expected)

}
