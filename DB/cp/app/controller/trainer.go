package controller

import (
	"cp/app/model"
	"cp/app/server"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
)

// SelectTrainers shows web page with filled tables of trainers
func SelectTrainers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	trainers, err := server.SelectTrainersList()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	path := filepath.Join("public", "pages", "trainer.html")
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = tmpl.ExecuteTemplate(w, "trainers", trainers); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

// InsertTrainer reads trainer info from form and inserts it into database
func InsertTrainer(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var newTrainer model.Trainer
	newTrainer.TrainerSecondName = r.FormValue("insert_second_name")
	newTrainer.TrainerName = r.FormValue("insert_name")
	newTrainer.TrainerThirdName = r.FormValue("insert_third_name")
	newTrainer.TrainerPhone = r.FormValue("insert_phone")
	if err := server.InsertNewTrainer(newTrainer); err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	printAnswer(w, successRes, successAns)
}

// DeleteTrainer reads trainer info from form and deletes client from database
func DeleteTrainer(w http.ResponseWriter, r *http.Request,
	p httprouter.Params) {
	id, err := strconv.Atoi(r.FormValue("delete_id"))
	if err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	if err = server.DeleteTrainer(id); err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	printAnswer(w, successRes, successAns)
}
