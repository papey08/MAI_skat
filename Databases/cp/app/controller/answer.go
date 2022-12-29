package controller

import (
	"html/template"
	"net/http"
	"path/filepath"
)

const (
	successRes = "Успешно"
	errorRes   = "Ошибка"
	successAns = "Операция успешно проведена. По возвращении обновите страницу, чтобы увидеть изменения."
)

// printAnswer shows page with status of operation
func printAnswer(w http.ResponseWriter, res string, ans string) {
	path := filepath.Join("public", "pages", "answer.html")
	tmpl, err := template.ParseFiles(path)
	data := struct {
		Res string
		Ans string
	}{
		res,
		ans,
	}
	if err = tmpl.ExecuteTemplate(w, "data", data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
