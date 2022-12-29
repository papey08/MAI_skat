package model

type Vehicle struct {
	VehicleSigh string  `json:"vehicleSigh"`
	Model       string  `json:"model"`
	TypeID      int     `json:"typeID"`
	PriceCoeff  float64 `json:"priceCoeff"`
}
