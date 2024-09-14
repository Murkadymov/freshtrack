package entity

type Supply struct {
	Driver       `json:"driver,required"`
	Goods        `json:"goods,required"`
	Manufacturer `json:"manufacturer,required"`
}

type Driver struct {
	DriverNumber  string `json:"driverNumber"`
	TractorNumber string `json:"tractorNumber"`
	TrailNumber   string `json:"trailNumber"`
}

type Goods struct {
	Goods string `json:"goods"`
}
type Manufacturer struct {
	Name   string `json:"name"`
	Origin string `json:"origin"`
}
