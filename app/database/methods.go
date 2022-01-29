package database

import (
	"starwars/app/config"
	"starwars/app/models"
)

const (
	insertSpacecraftsQuery   = "insert into spacecraft (name, class, crew, image, value, status, armament) values (?, ?, ?, ?, ?, ?, ?);"
	updateSpacecraftsQuery   = "update spacecraft set name=?, class=?, crew=?, image=?, value=?, status=?, armament=? where id=?"
	deleteSpacecraftsQuery   = "delete from spacecraft where id=? "
	selectSpacecraftsQuery   = "select id, name, class, crew, image, value, status, armament from spacecraft "
	selectSpacecraftsQueryID = "select id, name, class, crew, image, value, status, armament from spacecraft where id=?"
)

func (d *DB) GetSpacecraftByID(sID int64, c *config.Config) (*models.Spacecraft, error) {
	var s models.Spacecraft
	q := selectSpacecraftsQueryID
	err := d.db.QueryRow(q, sID).Scan(&s.ID, &s.Name, &s.Class, &s.Crew, &s.Image, &s.Value, &s.Status, &s.Armament)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (d *DB) GetSpacecrafts(c *config.Config) ([]*models.Spacecraft, error) {
	var sc []*models.Spacecraft
	q := selectSpacecraftsQuery
	rows, err := d.db.Query(q)
	if err != nil {
		return sc, err
	}
	for rows.Next() {
		var s models.Spacecraft
		err = rows.Scan(&s.ID, &s.Name, &s.Class, &s.Crew, &s.Image, &s.Value, &s.Status, &s.Armament)
		if err != nil {
			return sc, err
		}
		sc = append(sc, &s)
	}
	return sc, nil
}

func (d *DB) CreateSpacecraft(s *models.Spacecraft) (int64, error) {
	lastID, _, err := executeQuery(d.db, insertSpacecraftsQuery, s.Name.String, s.Class.String, s.Crew.Int64, s.Image.String, s.Value.Float64, s.Status.String, s.Armament.String)
	if err != nil {
		return 0, err
	}
	return lastID, nil
}

func (d *DB) UpdateSpacecraft(s *models.Spacecraft, c *config.Config) error {
	sc, err := d.GetSpacecraftByID(s.ID.Int64, c)
	if err != nil {
		return err
	}

	if sc.ID.Valid {
		//check if request has the valid and updated info, use the old values otherwise
		if s.Name.String != sc.Name.String && s.Name.String == "" {
			s.Name.String = sc.Name.String
		}

		if s.Class.String != sc.Class.String && s.Class.String == "" {
			s.Class.String = sc.Class.String
		}

		if s.Crew.Int64 != sc.Crew.Int64 && s.Crew.Int64 == 0 {
			s.Crew.Int64 = sc.Crew.Int64
		}

		if s.Image.String != sc.Image.String && s.Image.String == "" {
			s.Image.String = sc.Image.String
		}

		if s.Value.Float64 != sc.Value.Float64 && s.Value.Float64 == 0 {
			s.Value.Float64 = sc.Value.Float64
		}

		if s.Status.String != sc.Status.String && s.Status.String == "" {
			s.Status.String = sc.Status.String
		}

		if s.Armament.String != sc.Armament.String && s.Armament.String == "" {
			s.Armament.String = sc.Armament.String
		}

		q := updateSpacecraftsQuery
		_, _, err = executeQuery(d.db, q, s.Name.String, s.Class.String, s.Crew.Int64, s.Image.String, s.Value.Float64, s.Status.String, s.Armament.String, s.ID.Int64)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *DB) DeleteSpacecraft(s *models.Spacecraft, c *config.Config) (int64, error) {
	q := deleteSpacecraftsQuery
	_, deletedRows, err := executeQuery(d.db, q, s.ID.Int64)

	if err != nil {
		return deletedRows, err
	}
	return deletedRows, nil
}
