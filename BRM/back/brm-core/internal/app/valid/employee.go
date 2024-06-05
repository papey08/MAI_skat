package valid

import (
	"brm-core/internal/model"
	"net/mail"
	"net/url"
)

func CreateEmployee(empl model.Employee) bool {
	validFirstName := len([]rune(empl.FirstName)) <= 100
	validSecondName := len([]rune(empl.SecondName)) <= 100
	validJobTitle := len([]rune(empl.JobTitle)) <= 100
	validDepartment := len([]rune(empl.Department)) <= 100

	_, err := mail.ParseAddress(empl.Email)
	validEmail := err == nil && len(empl.Email) <= 100

	validImageUrl := true
	if empl.ImageURL != "" {
		_, err = url.ParseRequestURI(empl.ImageURL)
		validImageUrl = err == nil && len(empl.ImageURL) <= 200
	}

	return validFirstName &&
		validSecondName &&
		validJobTitle &&
		validDepartment &&
		validEmail &&
		validImageUrl
}

func UpdateEmployee(upd model.UpdateEmployee) bool {
	validFirstName := len([]rune(upd.FirstName)) <= 100
	validSecondName := len([]rune(upd.SecondName)) <= 100
	validJobTitle := len([]rune(upd.JobTitle)) <= 100
	validDepartment := len([]rune(upd.Department)) <= 100

	validImageUrl := true
	if upd.ImageURL != "" {
		_, err := url.ParseRequestURI(upd.ImageURL)
		validImageUrl = err == nil && len(upd.ImageURL) <= 200
	}

	return validFirstName &&
		validSecondName &&
		validJobTitle &&
		validDepartment &&
		validImageUrl
}
