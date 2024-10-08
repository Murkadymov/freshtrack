package entity

type Supply struct {
	Driver       `json:"driver"`
	Goods        `json:"goods"`
	Manufacturer `json:"manufacturer"`
}

type Driver struct {
	TractorNumber string `json:"tractorNumber" validate:"required"`
	TrailNumber   string `json:"trailNumber" validate:"required"`
	DriverNumber  string `json:"driverNumber" validate:"required"`
}

type Goods struct {
	Cargo string `json:"cargo" validate:"required"`
}
type Manufacturer struct {
	Name   string `json:"name" validate:"required"`
	Origin string `json:"origin" validate:"required"`
}
