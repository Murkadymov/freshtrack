package freshtrackrepo

import (
	"database/sql"
	"fmt"
	"freshtrack/internal/entity"
)

type Repository struct {
	db *sql.DB
}

func NewFreshTrackRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) AddSupply(e *entity.Supply) error {
	const op = "repository.freshtrackrepo.AddSupply"

	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	var driverID int
	err = tx.QueryRow(`
				INSERT INTO drivers (driver_number, tractor_number, trail_number)
				VALUES ($1,$2,$3)
				RETURNING driver_id;`,
		e.DriverNumber, e.TractorNumber, e.TrailNumber).Scan(&driverID)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

		return fmt.Errorf("%s: %w", op, err)
	}

	fmt.Println("REPO REPO", driverID, e.DriverNumber)
	return tx.Commit()
}
