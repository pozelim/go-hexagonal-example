package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/pozelim/go-hexagonal-example/internal/app"
	"log"
)

type Config struct {
	ConnStr string `mapstructure:"connStr"`
}
type EstateStorage struct {
	db *sql.DB
}

func NewEstateStorage(config Config) *EstateStorage {
	db, err := sql.Open("postgres", config.ConnStr)
	if err != nil {
		log.Fatal(err)
	}
	return &EstateStorage{
		db: db,
	}
}

func (e *EstateStorage) Save(estate app.Estate) error {
	_, err := e.db.Query(`
		INSERT INTO public.estate (estate_id, name, lat, lon, area)
			VALUES($1, $2, $3, $4, $5)`, estate.ID, estate.Name, estate.Point.Lat, estate.Point.Lon, estate.Area)
	return errors.Wrap(err, "Error while saving estate into database")
}

func (e *EstateStorage) Delete(id app.EstateID) error {
	_, err := e.db.Query(`DELETE FROM public.estate WHERE estate_id=$1`, id)
	return errors.Wrap(err, "Error while deleting estate from database")
}
