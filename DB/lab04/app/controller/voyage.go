package controller

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"lab04/app/model"
	"lab04/app/server"
	"net/http"
	"path/filepath"
	"strconv"
)

func GetVoyages(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	voyages, err := server.SelectVoyages()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	path := filepath.Join("public", "pages", "voyage.html")
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = tmpl.ExecuteTemplate(w, "data", voyages)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func InsertVoyage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var newVoyage model.Voyage
	newVoyage.DriverID, _ = strconv.Atoi(r.FormValue("driver_id"))
	newVoyage.PointBegin = r.FormValue("point_begin")
	newVoyage.PointEnd = r.FormValue("point_end")
	newVoyage.DateBegin = r.FormValue("date_begin")
	newVoyage.DateEnd = r.FormValue("date_end")
	err := server.InsertVoyage(newVoyage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func DeleteVoyage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, _ := strconv.Atoi(r.FormValue("voyage_id"))
	err := server.DeleteVoyage(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
