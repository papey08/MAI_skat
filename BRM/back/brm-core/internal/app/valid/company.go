package valid

import "brm-core/internal/model"

func CreateCompany(company model.Company) bool {
	return len([]rune(company.Name)) <= 100 && len([]rune(company.Description)) <= 1000
}

func UpdateCompany(upd model.UpdateCompany) bool {
	return len([]rune(upd.Name)) <= 100 && len([]rune(upd.Description)) <= 1000
}
