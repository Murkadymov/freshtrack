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

	var goodsID int
	err = tx.QueryRow(`
				INSERT INTO goods(cargo)
				VALUES($1)
				RETURNING goods_id;
`, e.Goods.Cargo).Scan(&goodsID)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
		return fmt.Errorf("%s: %w", op, err)
	}

	var manufacturerID int
	err = tx.QueryRow(`
		INSERT INTO manufacturer(name, origin)
		VALUES($1, $2)
		RETURNING manufacturer_id
`, e.Manufacturer.Name, e.Manufacturer.Origin).Scan(&manufacturerID)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = tx.Exec(`
		INSERT INTO supplies(driver_id, goods_id, manufacturer_id)
		VALUES($1,$2,$3)
		`, driverID, goodsID, manufacturerID)

	if err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
		return fmt.Errorf("%s: %w", op, err)
	}

	fmt.Println("End commit")
	return tx.Commit()
}

func (r *Repository) GetSupplyList() ([]entity.Supply, error) {
	const op = "repository.freshtrack.GetSupplyList"

	r.db.Begin()

	var supplyList []entity.Supply
	rows, err := r.db.Query(`
	SELECT
	    public.drivers.tractor_number,
    	public.drivers.trail_number,
    	public.drivers.driver_number,
    	public.goods.cargo,
    	public.manufacturer.name,
    	public.manufacturer.origin
	FROM public.supplies
	INNER JOIN public.drivers  ON public.drivers  .driver_id = public.supplies .driver_id
	INNER JOIN public.goods ON public.goods .goods_id  = public.supplies.goods_id
	INNER JOIN public.manufacturer ON public.manufacturer.manufacturer_id = public.supplies.manufacturer_id;`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	for rows.Next() {
		var supply entity.Supply
		err = rows.Scan(
			&supply.TractorNumber,
			&supply.TrailNumber,
			&supply.DriverNumber,
			&supply.Cargo,
			&supply.Name,
			&supply.Origin,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		supplyList = append(supplyList, supply)
	}
	defer rows.Close()

	return supplyList, nil
}
