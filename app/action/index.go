package action

import (
	"net/http"
	"starwars/app/config"
	"starwars/app/database"
	"starwars/app/helper"
)

func Index(db *database.DB, c *config.Config, w http.ResponseWriter, r *http.Request) {
	helper.SendResponse(w, r, nil, http.StatusOK, nil, "Starwars API")
}
