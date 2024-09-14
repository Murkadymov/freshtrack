package entity

type Supply struct {
	Driver       `json:"driver"`
	Goods        `json:"goods"`
	Manufacturer `json:"manufacturer"`
}

type Driver struct {
	DriverNumber  string `json:"driverNumber"`
	TractorNumber string `json:"tractorNumber"`
	TrailNumber   string `json:"trailNumber"`
}

type Goods struct {
	Cargo string `json:"cargo"`
}
type Manufacturer struct {
	Name   string `json:"name"`
	Origin string `json:"origin"`
}
