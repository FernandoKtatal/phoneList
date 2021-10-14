package app

import (
	"fmt"
	"log"
	"net/http"
	"postapi/app/models"
	"postapi/app/utils"
	"strings"
)

func (a *App) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Post API")
	}
}

func (a *App) CreatePostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := models.PostRequest{}
		err := utils.Parse(r, &req)
		if err != nil {
			log.Printf("Cannot parse post body. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		// Create the post
		p := &models.Phones{
			Country:     *req.Country,
			State:       *req.State,
			CountryCode: *req.CountryCode,
			PhoneNumber: *req.PhoneNumber,
		}

		// Save in DB
		err = a.DB.Insert(models.COLLECTION_PHONES, p)
		if err != nil {
			if strings.Contains(err.Error(), "column") {
				utils.SendResponse(w, r, nil, http.StatusBadRequest)
			} else {
				log.Printf("Cannot save post in DB. err=%v \n", err)
				utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			}
			return
		}

		resp := utils.MapPostToJSON(p)
		utils.SendResponse(w, r, resp, http.StatusOK)

	}
}

//func (a *App) GetPostsHandler() http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		phones, err := a.DB.Select()
//		if err != nil {
//			log.Printf("Cannot get phones, err=%v \n", err)
//			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
//			return
//		}
//
//		var resp = make([]models.JsonPost, len(phones))
//		for idx, post := range phones {
//			resp[idx] = utils.MapPostToJSON(post)
//		}
//
//		utils.SendResponse(w, r, resp, http.StatusOK)
//	}
//}
