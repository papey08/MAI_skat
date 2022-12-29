package server

import "lab04/app/model"

func SelectDrivers() ([]model.Driver, error) {
	rows, err := db.Query(`SELECT * FROM "Driver"`)
	if err != nil {
		return nil, err
	}
	var drivers []model.Driver
	for rows.Next() {
		var temp model.Driver
		err = rows.Scan(&temp.DriverID, &temp.DriverSecondName,
			&temp.DriverName, &temp.DriverThirdName,
			&temp.DriverClass, &temp.VehicleSigh)
		if err != nil {
			return nil, err
		}
		drivers = append(drivers, temp)
	}
	err = rows.Close()
	if err != nil {
		return nil, err
	}
	return drivers, nil
}

func InsertDriver(newDriver model.Driver) error {
	_, err := db.Exec(`
		INSERT INTO "Driver"
			(driver_second_name, driver_name, driver_third_name, 
			 driver_class, vehicle_sigh) 
		VALUES ($1, $2, $3, $4, $5)`,
		newDriver.DriverSecondName, newDriver.DriverName,
		newDriver.DriverThirdName, newDriver.DriverClass, newDriver.VehicleSigh)
	return err
}

func DeleteDriver(id int) error {
	_, err := db.Exec(`
		DELETE FROM "Driver"
		WHERE driver_id = $1`,
		id)
	return err
}
