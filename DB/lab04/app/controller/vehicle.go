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

func GetVehicles(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vehicles, err := server.SelectVehicles()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	path := filepath.Join("public", "pages", "vehicle.html")
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = tmpl.ExecuteTemplate(w, "data", vehicles)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func InsertVehicle(w http.ResponseWriter, r *http.Request,
	p httprouter.Params) {
	var newVehicle model.Vehicle
	newVehicle.VehicleSigh = r.FormValue("vehicle_sigh")
	newVehicle.Model = r.FormValue("model")
	newVehicle.TypeID, _ = strconv.Atoi(r.FormValue("type_id"))
	var err error
	newVehicle.PriceCoeff, err =
		strconv.ParseFloat(r.FormValue("price_coeff"), 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = server.InsertVehicle(newVehicle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func DeleteVehicle(w http.ResponseWriter, r *http.Request,
	p httprouter.Params) {
	sigh := r.FormValue("vehicle_sigh_delete")
	err := server.DeleteVehicle(sigh)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
