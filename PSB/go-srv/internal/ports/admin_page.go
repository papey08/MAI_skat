package ports

import (
	"context"
	"go-srv/internal/entities"
	"go-srv/internal/services"
	"net/http"
	"path/filepath"
	"strconv"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

type AdminPage struct {
	adminService *services.AdminService
	router       *httprouter.Router
}

func NewAdminPage(adminService *services.AdminService) *AdminPage {
	adminPage := &AdminPage{
		adminService: adminService,
		router:       httprouter.New(),
	}
	adminPage.setRoutes()
	return adminPage
}

func (a *AdminPage) Run(addr string) {
	_ = http.ListenAndServe(addr, a.router)
}

func (a *AdminPage) setRoutes() {
	a.router.ServeFiles("/public/*filepath", http.Dir("public"))
	a.router.GET("/", a.GetResponses)
	a.router.PUT("/response", a.UpdateResponse)
}

func (a *AdminPage) GetResponses(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	responses, err := a.adminService.GetResponses(context.Background(), 0, 10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	statistics, err := a.adminService.GetStatistics(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data := struct {
		Responses  []entities.Response
		Statistics []entities.Statistics
	}{
		Responses:  responses,
		Statistics: []entities.Statistics{statistics},
	}

	path := filepath.Join("public", "pages", "index.html")
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

func (a *AdminPage) UpdateResponse(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(r.FormValue("update_response_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	category := r.FormValue("update_response_category")
	if category != entities.Gratitude && category != entities.Suggestion && category != entities.Claim {
		http.Error(w, "wrong category", http.StatusBadRequest)
		return
	}

	if err := a.adminService.UpdateResponse(context.Background(), id, category); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
