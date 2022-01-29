package models

import "database/sql"

type Armament struct {
	ID    sql.NullInt64
	Title sql.NullString
	Qty   sql.NullString
}

type JsonArmament struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Qty   string `json:"qty"`
}

type ArmamentRequest struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Qty   string `json:"qty"`
}

func MapArmamentToJSON(t *Armament) JsonArmament {
	return JsonArmament{
		ID:    t.ID.Int64,
		Title: t.Title.String,
		Qty:   t.Qty.String,
	}
}
