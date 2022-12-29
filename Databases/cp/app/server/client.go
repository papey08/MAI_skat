package server

import "cp/app/model"

// InsertNewClient inserts new client in database
func InsertNewClient(newClient model.Client) error {
	_, err := db.Exec(`
		INSERT INTO Client 
		    (client_second_name, client_name, client_third_name, sex, 
		     birthdate, height, weight, subscription_begin, subscription_end) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		newClient.ClientSecondName, newClient.ClientName,
		newClient.ClientThirdName, newClient.Sex, newClient.Birthdate,
		newClient.Height, newClient.Weight,
		newClient.SubscriptionBegin, newClient.SubscriptionEnd)
	return err
}

// SelectClients returns slice of all clients
func SelectClients() ([]model.Client, error) {
	rows, err := db.Query(`
		SELECT * FROM Client
		ORDER BY subscription_id DESC`)
	if err != nil {
		return nil, err
	}
	var clients []model.Client
	for rows.Next() {
		var temp model.Client
		err = rows.Scan(&temp.SubscriptionID, &temp.ClientSecondName,
			&temp.ClientName, &temp.ClientThirdName, &temp.Sex,
			&temp.Birthdate, &temp.Height, &temp.Weight,
			&temp.SubscriptionBegin, &temp.SubscriptionEnd)
		if err != nil {
			return nil, err
		}
		temp.Birthdate = temp.Birthdate[:10]
		temp.SubscriptionBegin = temp.SubscriptionBegin[:10]
		temp.SubscriptionEnd = temp.SubscriptionEnd[:10]
		clients = append(clients, temp)
	}
	if err = rows.Close(); err != nil {
		return nil, err
	}
	return clients, nil
}

// SelectUnsubscribedClients returns slice of clients with
// expired subscription
func SelectUnsubscribedClients() ([]model.Client, error) {
	rows, err := db.Query(`
		SELECT * 
		FROM Client 
		WHERE subscription_end < CURRENT_DATE
		ORDER BY subscription_id DESC`)
	if err != nil {
		return nil, err
	}
	var clients []model.Client
	for rows.Next() {
		temp := model.Client{}
		err = rows.Scan(&temp.SubscriptionID, &temp.ClientSecondName,
			&temp.ClientName, &temp.ClientThirdName, &temp.Sex,
			&temp.Birthdate, &temp.Height, &temp.Weight,
			&temp.SubscriptionBegin, &temp.SubscriptionEnd)
		if err != nil {
			return nil, err
		}
		temp.Birthdate = temp.Birthdate[:10]
		temp.SubscriptionBegin = temp.SubscriptionBegin[:10]
		temp.SubscriptionEnd = temp.SubscriptionEnd[:10]
		clients = append(clients, temp)
	}
	if err = rows.Close(); err != nil {
		return nil, err
	}
	return clients, nil
}

// UpdateClientSubscription updates client's subscription by his id
func UpdateClientSubscription(id int, date string) error {
	_, err := db.Exec(`
		UPDATE client
		SET subscription_end = $1
		WHERE subscription_id = $2`,
		date, id)
	return err
}

// UpdateHeightAndWeight updates client's height and weight by his id
func UpdateHeightAndWeight(id int, height float64, weight float64) error {
	var err error
	if height == 0 && weight != 0 {
		_, err = db.Exec(`
		UPDATE Client
		SET weight = $1 WHERE
		(subscription_id = $2)`,
			weight, id)
	} else if weight == 0 && height != 0 {
		_, err = db.Exec(`
		UPDATE Client
		SET height = $1 WHERE
		(subscription_id = $2)`,
			height, id)
	} else {
		_, err = db.Exec(`
		UPDATE Client
		SET height = $1,
		weight = $2 WHERE
		(subscription_id = $3)`,
			height, weight, id)
	}
	return err
}

// DeleteClient deletes client from database by his id
func DeleteClient(id int) error {
	_, err := db.Exec(`
		DELETE FROM group_client
		WHERE (subscription_id = $1)`,
		id)
	if err != nil {
		return err
	}
	_, err = db.Exec(`
		DELETE FROM Client
		WHERE (subscription_id = $1)`,
		id)
	return err
}
