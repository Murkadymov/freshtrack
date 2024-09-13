package freshtrack

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewFreshTrackRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) AddSupply() {
	const op = "repository.freshtrack.AddSupply"

}
