package server

import "lab04/app/model"

func SelectVoyages() ([]model.Voyage, error) {
	rows, err := db.Query(`SELECT * FROM "Voyage"`)
	if err != nil {
		return nil, err
	}
	var voyages []model.Voyage
	for rows.Next() {
		temp := model.Voyage{}
		err = rows.Scan(&temp.VoyageID, &temp.DriverID,
			&temp.PointBegin, &temp.PointEnd,
			&temp.DateBegin, &temp.DateEnd)
		if err != nil {
			return nil, err
		}
		temp.DateBegin = temp.DateBegin[:10]
		temp.DateEnd = temp.DateEnd[:10]
		voyages = append(voyages, temp)
	}
	err = rows.Close()
	if err != nil {
		return nil, err
	}
	return voyages, nil
}

func InsertVoyage(newVoyage model.Voyage) error {
	_, err := db.Exec(`
		INSERT INTO "Voyage"
			(driver_id, point_begin, point_end, date_begin, date_end)
		VALUES ($1, $2, $3, $4, $5)`,
		newVoyage.DriverID, newVoyage.PointBegin, newVoyage.PointEnd,
		newVoyage.DateBegin, newVoyage.DateEnd)
	return err
}

func DeleteVoyage(id int) error {
	_, err := db.Exec(`
		DELETE FROM "Voyage"
		WHERE voyage_id = $1`,
		id)
	return err
}
