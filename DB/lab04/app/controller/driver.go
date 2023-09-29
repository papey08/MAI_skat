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

func GetDrivers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	drivers, err := server.SelectDrivers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	path := filepath.Join("public", "pages", "driver.html")
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = tmpl.ExecuteTemplate(w, "data", drivers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func InsertDriver(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var newDriver model.Driver
	newDriver.DriverSecondName = r.FormValue("driver_second_name")
	newDriver.DriverName = r.FormValue("driver_name")
	newDriver.DriverThirdName = r.FormValue("driver_third_name")
	newDriver.DriverClass = r.FormValue("driver_class")
	newDriver.VehicleSigh = r.FormValue("vehicle_sigh")
	err := server.InsertDriver(newDriver)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func DeleteDriver(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, _ := strconv.Atoi(r.FormValue("driver_id"))
	err := server.DeleteDriver(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
