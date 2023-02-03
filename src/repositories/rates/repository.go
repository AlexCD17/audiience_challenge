package rates

import (
	"database/sql"
	"fmt"
	"log"
)

type IRepository interface {
	GetRates(state string, estimationType string) (float32, error)
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// GetRates retrieves rate from sqlite file
func (r *Repository) GetRates(state, estimationType string) (rate float32, err error) {

	row := r.db.QueryRow(fmt.Sprintf("SELECT %s FROM rates WHERE state = '%s'", estimationType, state))

	if err = row.Scan(&rate); err != nil {
		log.Fatal(err)
	}

	err = nil
	return

}
