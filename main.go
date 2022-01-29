package main

import (
	"log"
	"net/http"
	"os"
	"starwars/app"
	"starwars/app/database"
)

func main() {
	//start an application
	app := app.New()

	//setup database connection
	app.DB = &database.DB{}
	err := app.DB.Open(app.Config)
	check(err)
	defer app.DB.Close()

	//start router and listen to port
	http.HandleFunc("/", app.Router.ServeHTTP)

	log.Println("App is running...")
	err = http.ListenAndServe(":"+app.Config.AppPort, nil)
	check(err)
}

func check(e error) {
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}
}
