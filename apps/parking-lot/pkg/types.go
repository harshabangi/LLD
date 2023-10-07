package pkg

type Floor struct {
	ID        string
	Spots     []Spot
	SpotMap   map[string]int
	UsedSpots map[string]int
}

type Vehicle interface {
}

type VehicleCommon struct {
	LicenseNumber string `json:"license_number"`
}

type Car struct {
	VehicleCommon
}

type MotorBike struct {
	VehicleCommon
}

type SpotType string

const (
	CarSpotType       SpotType = "car"
	MotorBikeSpotType SpotType = "motorbike"
)

type Spot struct {
	ID          string   `json:"id"`
	SpotType    SpotType `json:"spot_type"`
	Price       int      `json:"price"`
	IsAvailable bool     `json:"is_available"`
	Vehicle     Vehicle  `json:"vehicle"`
}
