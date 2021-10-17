package main

import (
	"log"
	"net/http"
	"os"
	"postapi/app"
	"postapi/app/database"
)

func main() {
	app := app.New()
	err := database.Connect()
	check(err)

	defer database.Disconnect()

	http.HandleFunc("/", app.Router.ServeHTTP)

	log.Println("App running..")
	err = http.ListenAndServe(":9000", nil)
	check(err)
}

func check(e error) {
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}
}
