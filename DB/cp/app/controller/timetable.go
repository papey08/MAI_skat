package controller

import (
	"cp/app/model"
	"cp/app/server"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// variables means which timetables will be shown on web page

var timetableByGroup = 0
var timetableByProgram = 0
var timetableByTrainer = 0

// SelectTimetable shows web page with filled tables of timetables
func SelectTimetable(w http.ResponseWriter, r *http.Request,
	p httprouter.Params) {
	byGroup, err := server.SelectTimetableByGroup(timetableByGroup)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	byProgram, err := server.SelectTimetableByProgram(timetableByProgram)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	byTrainer, err := server.SelectTimetableByTrainer(timetableByTrainer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data := struct {
		ByGroup   []model.Timetable
		ByProgram []model.Timetable
		ByTrainer []model.Timetable
	}{
		byGroup,
		byProgram,
		byTrainer,
	}
	path := filepath.Join("public", "pages", "timetable.html")
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = tmpl.ExecuteTemplate(w, "data", data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

// SelectTimetableByGroup sets new value of variable timetableByGroup
func SelectTimetableByGroup(w http.ResponseWriter, r *http.Request,
	p httprouter.Params) {
	tempStr := r.FormValue("select_group_id")
	if tempStr != "" {
		timetableByGroup, _ = strconv.Atoi(tempStr)
	}
	printAnswer(w, successRes, successAns)
}

// SelectTimetableByProgram sets new value of variable timetableByProgram
func SelectTimetableByProgram(w http.ResponseWriter, r *http.Request,
	p httprouter.Params) {
	tempStr := r.FormValue("select_program_id")
	if tempStr != "" {
		timetableByProgram, _ = strconv.Atoi(tempStr)
	}
	printAnswer(w, successRes, successAns)
}

// SelectTimetableByTrainer sets new value of variable timetableByTrainer
func SelectTimetableByTrainer(w http.ResponseWriter, r *http.Request,
	p httprouter.Params) {
	tempStr := r.FormValue("select_trainer_id")
	if tempStr != "" {
		timetableByTrainer, _ = strconv.Atoi(tempStr)
	}
	printAnswer(w, successRes, successAns)
}

// InsertTimetable reads timetable info from form and inserts it into database
func InsertTimetable(w http.ResponseWriter, r *http.Request,
	p httprouter.Params) {
	var newTimetable model.Timetable
	newTimetable.GroupID, _ = strconv.Atoi(r.FormValue("insert_group_id"))
	newTimetable.Weekday = r.FormValue("insert_weekday")
	newTimetable.TrainingTime = r.FormValue("insert_training_time")
	if err := server.InsertTimetable(newTimetable); err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	printAnswer(w, successRes, successAns)
}

// DeleteTimetable reads timetable info from form and deletes it from database
func DeleteTimetable(w http.ResponseWriter, r *http.Request,
	p httprouter.Params) {
	var newTimetable model.Timetable
	newTimetable.GroupID, _ = strconv.Atoi(r.FormValue("insert_group_id"))
	newTimetable.Weekday = r.FormValue("insert_weekday")
	newTimetable.TrainingTime = r.FormValue("insert_training_time")
	if err := server.DeleteTimetable(newTimetable); err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	printAnswer(w, successRes, successAns)
}
