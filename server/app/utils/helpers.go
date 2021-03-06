package utils

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"io"
	"log"
	"net/http"
	"postapi/app/models"
)

func Parse(r *http.Request, data *models.PostRequest) error {
	v := validator.New()
	req, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(req, &data)
	if err != nil {
		return err
	}
	err = v.Struct(data)

	return err
}

func SendResponse(w http.ResponseWriter, _ *http.Request, data interface{}, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("Cannot format json. err=%v\n", err)
	}
}

func MapPostToJSON(p *models.Phones) models.JsonPost {
	return models.JsonPost{
		Country:     p.Country,
		State:       p.State,
		CountryCode: p.CountryCode,
		PhoneNumber: p.PhoneNumber,
	}
}
