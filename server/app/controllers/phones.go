package controllers

import (
	"log"
	"net/http"
	"postapi/app/models"
	"postapi/app/services"
	"postapi/app/utils"
	"strconv"
	"strings"
)

func NewPhone() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := models.PostRequest{}
		err := utils.Parse(r, &req)
		if err != nil {
			log.Printf("Cannot parse post body. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}


		p, err := services.CreatePhone(*req.Country, *req.CountryCode, *req.PhoneNumber)
		if err != nil {
			if strings.Contains(err.Error(), "column") {
				utils.SendResponse(w, r, nil, http.StatusBadRequest)
			} else {
				utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			}
			return
		}

		resp := utils.MapPostToJSON(p)
		utils.SendResponse(w, r, resp, http.StatusOK)

	}

}

func GetPhone() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var validState *bool
		country := r.URL.Query().Get("country")
		state := r.URL.Query().Get("state")
		stateBool, err := strconv.ParseBool(state)
		if err == nil {
			validState = &stateBool
		}

		phones, err := services.CapturePhone(country, validState)
		if err != nil {
			log.Printf("Cannot get phones, err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		var resp = make([]models.JsonPost, len(phones))
		for idx, post := range phones {
			resp[idx] = utils.MapPostToJSON(&post)
		}

		utils.SendResponse(w, r, resp, http.StatusOK)
	}
}

