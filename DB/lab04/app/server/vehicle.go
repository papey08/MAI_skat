package server

import "lab04/app/model"

func SelectVehicles() ([]model.Vehicle, error) {
	rows, err := db.Query(`SELECT * FROM "Vehicle"`)
	if err != nil {
		return nil, err
	}
	var vehicles []model.Vehicle
	for rows.Next() {
		temp := model.Vehicle{}
		err = rows.Scan(&temp.VehicleSigh, &temp.Model,
			&temp.TypeID, &temp.PriceCoeff)
		if err != nil {
			return nil, err
		}
		vehicles = append(vehicles, temp)
	}
	err = rows.Close()
	if err != nil {
		return nil, err
	}
	return vehicles, nil
}

func InsertVehicle(newVehicle model.Vehicle) error {
	_, err := db.Exec(`
		INSERT INTO "Vehicle"
			(vehicle_sigh, model, type_id, price_coeff)
		VALUES ($1, $2, $3, $4)`,
		newVehicle.VehicleSigh, newVehicle.Model,
		newVehicle.TypeID, newVehicle.PriceCoeff)
	return err
}

func DeleteVehicle(sigh string) error {
	_, err := db.Exec(`
		DELETE FROM "Vehicle"
		WHERE vehicle_sigh = $1`,
		sigh)
	return err
}
