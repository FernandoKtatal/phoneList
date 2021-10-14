package app

import (
	"github.com/gorilla/mux"
	"postapi/app/database"
)

type App struct {
	Router *mux.Router
	DB     database.PostDB
}

func New() *App {
	a := &App{
		Router: mux.NewRouter(),
	}

	a.initRoutes()
	return a
}

func (a *App) initRoutes() {
	a.Router.HandleFunc("/", a.IndexHandler()).Methods("GET")
	a.Router.HandleFunc("/api/phone", a.CreatePostHandler()).Methods("POST")
	//a.Router.HandleFunc("/api/phone", a.GetPostsHandler()).Methods("GET")
}
