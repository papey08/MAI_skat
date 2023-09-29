package model

type Driver struct {
	DriverID         int    `json:"driverID"`
	DriverSecondName string `json:"driverSecondName"`
	DriverName       string `json:"driverName"`
	DriverThirdName  string `json:"driverThirdName"`
	DriverClass      string `json:"driverClass"`
	VehicleSigh      string `json:"vehicleSigh"`
}
