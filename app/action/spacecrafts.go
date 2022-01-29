package action

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"starwars/app/config"
	"starwars/app/database"
	"starwars/app/helper"
	"starwars/app/models"
)

type ValidationError struct {
	Err error
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%v", e.Err)
}

//get all spacecrafts
func GetSpacecrafts(db *database.DB, c *config.Config, w http.ResponseWriter, r *http.Request) {
	resp, err := GetAllSpacecrafts(db, c)
	if helper.CheckErr(w, r, helper.GetSpacecraftErrorCode, http.StatusInternalServerError, err) {
		return
	}
	helper.SendResponse(w, r, resp, http.StatusOK, nil, "")
}

func GetAllSpacecrafts(db *database.DB, c *config.Config) ([]models.JsonSpacecraft, error) {
	sc, err := db.GetSpacecrafts(c)
	if err != nil {
		return []models.JsonSpacecraft{}, err
	}

	var resp = make([]models.JsonSpacecraft, len(sc))
	for idx, s := range sc {
		resp[idx] = models.MapSpaceshipToJSON(s)
	}
	return resp, err
}

//get single spacecraft by id
func GetSpacecraft(db *database.DB, c *config.Config, w http.ResponseWriter, r *http.Request) {
	req := models.SpacecraftRequest{}
	//parse payload and check for errors
	err := helper.Parse(w, r, &req)
	if helper.CheckErr(w, r, helper.InvalidJsonCode, http.StatusBadRequest, err) {
		return
	}

	sID := req.ID
	resp, err := db.GetSpacecraftByID(sID, c)
	if helper.CheckErr(w, r, helper.GetSpacecraftErrorCode, http.StatusInternalServerError, err) {
		return
	}
	helper.SendResponse(w, r, models.MapSpaceshipToJSON(resp), http.StatusOK, nil, "")
}

func CreateSpacecraft(db *database.DB, c *config.Config, w http.ResponseWriter, r *http.Request) {
	req := models.SpacecraftRequest{}
	//parse payload and check for errors
	err := helper.Parse(w, r, &req)
	if helper.CheckErr(w, r, helper.InvalidJsonCode, http.StatusBadRequest, err) {
		return
	}

	if helper.CheckErr(w, r, helper.InvalidJsonCode, http.StatusBadRequest, validateCreateRequest(req)) {
		return
	}

	// Create Spacecraft
	s := &models.Spacecraft{
		Name:     sql.NullString{String: req.Name},
		Class:    sql.NullString{String: req.Class},
		Crew:     sql.NullInt64{Int64: req.Crew},
		Image:    sql.NullString{String: req.Image},
		Value:    sql.NullFloat64{Float64: req.Value},
		Status:   sql.NullString{String: req.Status},
		Armament: sql.NullString{String: req.Armament},
	}

	// Save in DB
	_, err = db.CreateSpacecraft(s)
	if helper.CheckErr(w, r, helper.CreateSpacecraftErrorCode, http.StatusInternalServerError, err) {
		return
	}

	helper.SendResponse(w, r, nil, http.StatusOK, nil, "Spaceship created successfully.")
}

func UpdateSpacecraft(db *database.DB, c *config.Config, w http.ResponseWriter, r *http.Request) {
	req := models.SpacecraftRequest{}
	//parse payload and check for errors
	err := helper.Parse(w, r, &req)
	if helper.CheckErr(w, r, helper.InvalidJsonCode, http.StatusBadRequest, err) {
		return
	}

	if helper.CheckErr(w, r, helper.InvalidJsonCode, http.StatusBadRequest, validateUpdateRequest(req)) {
		return
	}

	// Prepare Spacecraft
	s := &models.Spacecraft{
		ID:       sql.NullInt64{Int64: req.ID},
		Name:     sql.NullString{String: req.Name},
		Class:    sql.NullString{String: req.Class},
		Crew:     sql.NullInt64{Int64: req.Crew},
		Image:    sql.NullString{String: req.Image},
		Value:    sql.NullFloat64{Float64: req.Value},
		Status:   sql.NullString{String: req.Status},
		Armament: sql.NullString{String: req.Armament},
	}

	// Update in DB
	err = db.UpdateSpacecraft(s, c)
	if helper.CheckErr(w, r, helper.UpdateSpacecraftErrorCode, http.StatusInternalServerError, err) {
		return
	}

	helper.SendResponse(w, r, nil, http.StatusOK, nil, "Spaceship updated successfully.")
}

func DeleteSpacecraft(db *database.DB, c *config.Config, w http.ResponseWriter, r *http.Request) {
	req := models.SpacecraftRequest{}
	//parse payload and check for errors
	err := helper.Parse(w, r, &req)
	if helper.CheckErr(w, r, helper.InvalidJsonCode, http.StatusBadRequest, err) {
		return
	}

	if helper.CheckErr(w, r, helper.InvalidJsonCode, http.StatusBadRequest, validateDeleteRequest(req)) {
		return
	}

	// Prepare Spacecraft
	s := &models.Spacecraft{
		ID: sql.NullInt64{Int64: req.ID},
	}

	// Delete in DB
	deletedRec, err := db.DeleteSpacecraft(s, c)
	if err != nil {
		helper.SendResponse(w, r, helper.DeleteSpacecraftError, http.StatusInternalServerError, err, "")
		return
	}

	if deletedRec == 0 {
		helper.SendResponse(w, r, helper.DeleteSpacecraftError, http.StatusInternalServerError, &ValidationError{Err: errors.New("record not found")}, "")
		return
	}

	helper.SendResponse(w, r, nil, http.StatusOK, nil, "Spacecraft deleted.")
}

func validateDeleteRequest(s models.SpacecraftRequest) error {
	if s.ID == 0 {
		return &ValidationError{
			Err: errors.New("please enter id to delete the spaceship"),
		}
	}
	return nil
}

func validateUpdateRequest(s models.SpacecraftRequest) error {
	if s.Name == "" && s.Class == "" && s.Crew == 0 && s.Image == "" && s.Value == 0 && s.Status == "" {
		return &ValidationError{
			Err: errors.New("nothing to update"),
		}
	}
	return nil
}

func validateCreateRequest(s models.SpacecraftRequest) error {
	if s.Name == "" {
		return &ValidationError{
			Err: errors.New("please enter name"),
		}
	}

	if s.Class == "" {
		return &ValidationError{
			Err: errors.New("please enter class"),
		}
	}

	if s.Image == "" {
		return &ValidationError{
			Err: errors.New("please provide an image"),
		}
	}

	if s.Crew == 0 {
		return &ValidationError{
			Err: errors.New("please provide crew size"),
		}
	}

	if s.Status == "" {
		return &ValidationError{
			Err: errors.New("please enter status"),
		}
	}

	if s.Value == 0 {
		return &ValidationError{
			Err: errors.New("please provide value"),
		}
	}
	return nil
}
