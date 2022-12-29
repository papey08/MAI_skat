package server

import "cp/app/model"

// SelectTimetableByGroup returns timetable of group with id groupID
func SelectTimetableByGroup(groupID int) ([]model.Timetable, error) {
	rows, err := db.Query(`
	SELECT group_id, weekday, training_time 
	FROM timetable 
	INNER JOIN times 
	USING(time_id) 
	WHERE group_id = $1`,
		groupID)
	if err != nil {
		return nil, err
	}
	var timetables []model.Timetable
	for rows.Next() {
		var temp model.Timetable
		err = rows.Scan(&temp.GroupID, &temp.Weekday, &temp.TrainingTime)
		if err != nil {
			return nil, err
		}
		temp.TrainingTime = temp.TrainingTime[11:16]
		timetables = append(timetables, temp)
	}
	if err = rows.Close(); err != nil {
		return nil, err
	}
	return timetables, nil
}

// SelectTimetableByProgram returns timetable of groups
// with program programID
func SelectTimetableByProgram(programID int) ([]model.Timetable, error) {
	rows, err := db.Query(`
	SELECT group_id, weekday, training_time 
	FROM timetable 
	INNER JOIN times 
	USING(time_id) 
	WHERE program_id = $1`,
		programID)
	if err != nil {
		return nil, err
	}
	var timetables []model.Timetable
	for rows.Next() {
		var temp model.Timetable
		err = rows.Scan(&temp.GroupID, &temp.Weekday, &temp.TrainingTime)
		if err != nil {
			return nil, err
		}
		temp.TrainingTime = temp.TrainingTime[11:16]
		timetables = append(timetables, temp)
	}
	if err = rows.Close(); err != nil {
		return nil, err
	}
	return timetables, nil
}

// SelectTimetableByTrainer returns timetable of groups
// // with trainer trainerID
func SelectTimetableByTrainer(trainerID int) ([]model.Timetable, error) {
	rows, err := db.Query(`
	SELECT group_id, weekday, training_time 
	FROM timetable 
		INNER JOIN times USING(time_id) 
		INNER JOIN fc_group USING(group_id)
	WHERE trainer_id = $1`,
		trainerID)
	if err != nil {
		return nil, err
	}
	var timetables []model.Timetable
	for rows.Next() {
		var temp model.Timetable
		err = rows.Scan(&temp.GroupID, &temp.Weekday, &temp.TrainingTime)
		if err != nil {
			return nil, err
		}
		temp.TrainingTime = temp.TrainingTime[11:16]
		timetables = append(timetables, temp)
	}
	if err = rows.Close(); err != nil {
		return nil, err
	}
	return timetables, nil
}

// InsertTimetable inserts new timetable in database
func InsertTimetable(newTimetable model.Timetable) error {
	_, err := db.Exec(`
	INSERT INTO timetable (time_id, group_id)
	VALUES
		((SELECT time_id
		 FROM times
		 WHERE weekday = $2 AND training_time = $3 AND program_id = (
		     SELECT program_id
		     FROM fc_group
		     WHERE fc_group.group_id = $1
		 )), $1)`,
		newTimetable.GroupID, newTimetable.Weekday, newTimetable.TrainingTime)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTimetable deletes timetable from database
func DeleteTimetable(newTimetable model.Timetable) error {
	_, err := db.Exec(`
	DELETE FROM timetable
	WHERE time_id = (
		SELECT time_id
		FROM times
		WHERE weekday = $2 AND training_time = $3 AND program_id = (
		    SELECT program_id
		    FROM fc_group
		    WHERE fc_group.group_id = $1
		 )) AND group_id = $1`,
		newTimetable.GroupID, newTimetable.Weekday, newTimetable.TrainingTime)
	if err != nil {
		return err
	}
	return nil
}
