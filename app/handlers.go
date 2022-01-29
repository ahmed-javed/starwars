package app

import (
	"net/http"
	"starwars/app/action"
)

//each handler will call the singal action and there should be much code in the handler function, keep it clean

func (a *App) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		action.Index(a.DB, a.Config, w, r)
	}
}

func (a *App) GetSpaceshipsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		action.GetSpacecrafts(a.DB, a.Config, w, r)
	}
}

func (a *App) GetSpaceshipHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		action.GetSpacecraft(a.DB, a.Config, w, r)
	}
}

func (a *App) CreateSpaceshipHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		action.CreateSpacecraft(a.DB, a.Config, w, r)
	}
}

func (a *App) UpdateSpaceshipHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		action.UpdateSpacecraft(a.DB, a.Config, w, r)
	}
}

func (a *App) DeleteSpaceshipHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		action.DeleteSpacecraft(a.DB, a.Config, w, r)
	}
}
