package models

import "database/sql"

type Spacecraft struct {
	ID       sql.NullInt64
	Name     sql.NullString
	Class    sql.NullString
	Armament sql.NullString
	Crew     sql.NullInt64
	Image    sql.NullString
	Value    sql.NullFloat64
	Status   sql.NullString
}

type JsonSpacecraft struct {
	ID       int64   `json:"id"`
	Name     string  `json:"name"`
	Class    string  `json:"class"`
	Armament string  `json:"armament_id"`
	Crew     int64   `json:"crew"`
	Image    string  `json:"image"`
	Value    float64 `json:"value"`
	Status   string  `json:"status"`
}

type SpacecraftRequest struct {
	ID       int64   `json:"id"`
	Name     string  `json:"name"`
	Class    string  `json:"class"`
	Armament string  `json:"armament_id"`
	Crew     int64   `json:"crew"`
	Image    string  `json:"image"`
	Value    float64 `json:"value"`
	Status   string  `json:"status"`
}

func MapSpaceshipToJSON(s *Spacecraft) JsonSpacecraft {
	return JsonSpacecraft{
		ID:       s.ID.Int64,
		Name:     s.Name.String,
		Class:    s.Class.String,
		Armament: s.Armament.String,
		Crew:     s.Crew.Int64,
		Image:    s.Image.String,
		Value:    s.Value.Float64,
		Status:   s.Status.String,
	}
}
