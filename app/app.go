package app

import (
	"starwars/app/config"
	"starwars/app/database"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *database.DB
	Config *config.Config
}

func New() *App {
	//create a router and prepare configurations required in the application
	a := &App{
		Router: mux.NewRouter(),
		Config: config.PrepareConfigurations(),
	}

	//register all the routes in application
	a.routes()
	return a
}

func (a *App) routes() {
	a.Router.HandleFunc("/", a.IndexHandler()).Methods("GET")
	a.Router.HandleFunc("/spaceships", a.GetSpaceshipsHandler()).Methods("GET")
	a.Router.HandleFunc("/spaceship", a.GetSpaceshipHandler()).Methods("GET")
	a.Router.HandleFunc("/spaceship", a.CreateSpaceshipHandler()).Methods("POST")
	a.Router.HandleFunc("/spaceship", a.UpdateSpaceshipHandler()).Methods("PUT")
	a.Router.HandleFunc("/spaceship", a.DeleteSpaceshipHandler()).Methods("DELETE")
}
