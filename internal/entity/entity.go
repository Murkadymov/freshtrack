package entity

type Car struct {
	TractorNumber string `json:"tractorNumber"`
	TrailNumber   string `json:"trailNumber"`
	DriverNumber  string `json:"driverNumber"`
}

type Cargo struct {
	Goods        string `json:"goods"`
	Country      string `json:"country"`
	Manufacturer string `json:"manufacturer"`
}
