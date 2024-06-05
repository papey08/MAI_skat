package employees

type addEmployeeRequest struct {
	CompanyId  uint64 `json:"company_id"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	JobTitle   string `json:"job_title"`
	Department string `json:"department"`
	ImageURL   string `json:"image_url"`
}

type updateEmployeeRequest struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	JobTitle   string `json:"job_title"`
	Department string `json:"department"`
	ImageURL   string `json:"image_url"`
}
