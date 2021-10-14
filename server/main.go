package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"postapi/app"
	"postapi/app/database"
	"time"
)

func main() {
	ctx,_ := context.WithTimeout(context.Background(), 180*time.Second)
	app := app.New()
	app.DB = &database.DB{}
	err := app.DB.Connect(ctx)
	check(err)

	defer func(DB database.PostDB, ctx context.Context) {
		err := DB.Disconnect(ctx)
		if err != nil {

		}
	}(app.DB, ctx)

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
