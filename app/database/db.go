package database

import (
	"database/sql"
	"starwars/app/config"
	"starwars/app/models"

	_ "github.com/go-sql-driver/mysql"
)

type OperateDB interface {
	Open() error
	Close() error
	GetSpacecrafts() ([]*models.Spacecraft, error)
	CreateSpacecraft(*models.Spacecraft) error
	UpdateSpacecraft() ([]*models.Spacecraft, error)
	DeleteSpacecraft() ([]*models.Spacecraft, error)
}

type DB struct {
	db *sql.DB
}

func prepareConnectionString(c *config.Config) string {
	return c.DbConfig.DbUsername + ":" + c.DbConfig.DbPassword + "@tcp(" + c.DbConfig.DbHost + ":" + c.DbConfig.DbPort + ")/" + c.DbConfig.DbSchema
}

func (d *DB) Open(c *config.Config) error {
	db, err := sql.Open(c.DbConfig.DbDriver, prepareConnectionString(c))
	if err != nil {
		return err
	}

	db.SetMaxIdleConns(500)
	db.SetMaxOpenConns(500)

	err = db.Ping()
	if err != nil {
		return err
	}

	d.db = db
	return err
}

func (d *DB) Close() error {
	return d.db.Close()
}

func executeQuery(db *sql.DB, q string, args ...interface{}) (int64, int64, error) {
	err := db.Ping()
	if err != nil {
		return 0, 0, err
	}

	tx, err := db.Begin()
	if err != nil {
		return 0, 0, err
	}

	u, err := tx.Prepare(q)
	if err != nil {
		return 0, 0, err
	}

	r, err := u.Exec(args...)
	if err != nil {
		return 0, 0, err
	}

	u.Close()
	err = tx.Commit()
	if err != nil {
		return 0, 0, err
	}

	n, err := r.RowsAffected()
	if err != nil {
		return 0, 0, err
	}

	lastID, err := r.LastInsertId()
	if err != nil {
		return 0, n, err
	}

	return lastID, n, nil
}
