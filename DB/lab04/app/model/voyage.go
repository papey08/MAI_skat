package model

type Voyage struct {
	VoyageID   int    `json:"voyageID"`
	DriverID   int    `json:"driverID"`
	PointBegin string `json:"pointBegin"`
	PointEnd   string `json:"pointEnd"`
	DateBegin  string `json:"dateBegin"`
	DateEnd    string `json:"dateEnd"`
}
