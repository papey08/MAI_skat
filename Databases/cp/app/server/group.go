package server

import "cp/app/model"

// InsertNewGroup inserts new group in database
func InsertNewGroup(newGroup model.Group) error {
	_, err := db.Exec(`
		INSERT INTO FC_Group 
		    (program_id, notes, trainer_id, clients_amount) 
		VALUES ($1, $2, $3, $4)`,
		newGroup.ProgramID, newGroup.Notes,
		newGroup.TrainerID, newGroup.ClientsAmount)
	return err
}

// SelectGroupList returns slice of clients in group with id number
func SelectGroupList(number int) ([]model.Client, error) {
	rows, err := db.Query(`
		SELECT subscription_id, client_second_name, client_name, 
		       client_third_name, sex, birthdate, height, weight, 
		       subscription_begin, subscription_end 
		FROM Client 
		    INNER JOIN group_client USING(subscription_id) 
		WHERE group_id = $1
		ORDER BY subscription_id DESC`, number)
	if err != nil {
		return nil, err
	}
	var clients []model.Client
	for rows.Next() {
		temp := model.Client{}
		err = rows.Scan(&temp.SubscriptionID, &temp.ClientSecondName,
			&temp.ClientName, &temp.ClientThirdName, &temp.Sex, &temp.Birthdate,
			&temp.Height, &temp.Weight, &temp.SubscriptionBegin,
			&temp.SubscriptionEnd)
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

// SelectPrograms returns slice of all programs
func SelectPrograms() ([]model.Program, error) {
	rows, err := db.Query(`
		SELECT * FROM program
		ORDER BY program_id`)
	if err != nil {
		return nil, err
	}
	var programs []model.Program
	for rows.Next() {
		var temp model.Program
		err = rows.Scan(&temp.ProgramID, &temp.ProgramName)
		if err != nil {
			return nil, err
		}
		programs = append(programs, temp)
	}
	if err = rows.Close(); err != nil {
		return nil, err
	}
	return programs, nil
}

// SelectGroupsList returns slice of all groups
func SelectGroupsList() ([]model.Group, error) {
	rows, err := db.Query(`
		SELECT * FROM FC_Group
		ORDER BY group_id DESC`)
	if err != nil {
		return nil, err
	}
	var groups []model.Group
	for rows.Next() {
		temp := model.Group{}
		err = rows.Scan(&temp.GroupID, &temp.ProgramID, &temp.Notes,
			&temp.TrainerID, &temp.ClientsAmount)
		if err != nil {
			return nil, err
		}
		groups = append(groups, temp)
	}
	if err = rows.Close(); err != nil {
		return nil, err
	}
	return groups, nil
}

// InsertClientIntoGroup inserts new client in group by ids
func InsertClientIntoGroup(clientID int, groupID int) error {
	_, err := db.Exec(`
		INSERT INTO group_client (group_id, subscription_id)
		VALUES ($1, $2)`,
		groupID, clientID)
	return err
}

// DeleteClientFromGroup deletes client from group database
func DeleteClientFromGroup(clientID int, groupID int) error {
	_, err := db.Exec(`
		DELETE FROM group_client
		WHERE group_id = $1 AND subscription_id = $2`,
		groupID, clientID)
	return err
}

// DeleteGroup deletes group by its id from database
func DeleteGroup(id int) error {
	_, err := db.Exec(`
		DELETE FROM FC_Group
		WHERE (group_id = $1)`,
		id)
	return err
}
