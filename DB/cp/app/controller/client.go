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

// SelectClients shows web page with filled tables of clients
func SelectClients(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	clients, err := server.SelectClients()
	unsubscribedClients, err := server.SelectUnsubscribedClients()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data := struct {
		Clients             []model.Client
		UnsubscribedClients []model.Client
	}{
		clients,
		unsubscribedClients,
	}
	path := filepath.Join("public", "pages", "client.html")
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

// InsertClient reads client info from form and inserts it into database
func InsertClient(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var newClient model.Client
	newClient.ClientSecondName = r.FormValue("insert_second_name")
	newClient.ClientName = r.FormValue("insert_name")
	newClient.ClientThirdName = r.FormValue("insert_third_name")
	newClient.Sex = r.FormValue("insert_sex")
	newClient.Birthdate = r.FormValue("insert_birthdate")
	newClient.Height, _ = strconv.ParseFloat(r.FormValue("insert_height"), 64)
	newClient.Weight, _ = strconv.ParseFloat(r.FormValue("insert_weight"), 64)
	newClient.SubscriptionBegin = r.FormValue("insert_subscription_begin")
	newClient.SubscriptionEnd = r.FormValue("insert_subscription_end")
	if err := server.InsertNewClient(newClient); err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	printAnswer(w, successRes, successAns)
}

// UpdateClientSubscription reads client info from form and updates info
func UpdateClientSubscription(w http.ResponseWriter, r *http.Request,
	p httprouter.Params) {
	id, err := strconv.Atoi(r.FormValue("update_subscription_id"))
	if err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	date := r.FormValue("update_subscription_subscription_end")
	if err = server.UpdateClientSubscription(id, date); err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	printAnswer(w, successRes, successAns)
}

// UpdateHeightAndWeight reads client info from form and updates info
func UpdateHeightAndWeight(w http.ResponseWriter, r *http.Request,
	p httprouter.Params) {
	id, err := strconv.Atoi(r.FormValue("update_height&weight_id"))
	if err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	newHeight, err :=
		strconv.ParseFloat(r.FormValue("update_height&weight_height"), 64)
	if err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	newWeight, err :=
		strconv.ParseFloat(r.FormValue("update_height&weight_weight"), 64)
	if err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	if err = server.UpdateHeightAndWeight(id, newHeight, newWeight); err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	printAnswer(w, successRes, successAns)
}

// DeleteClient reads client info from form and deletes client from database
func DeleteClient(w http.ResponseWriter, r *http.Request,
	p httprouter.Params) {
	id, err := strconv.Atoi(r.FormValue("delete_id"))
	if err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	if err = server.DeleteClient(id); err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	printAnswer(w, successRes, successAns)
}
