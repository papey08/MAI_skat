package httpserver

type addCompanyAndOwnerRequest struct {
	Company addCompanyData `json:"company"`
	Owner   addOwnerData   `json:"owner"`
}

type addCompanyData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Industry    string `json:"industry"`
}

type addOwnerData struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	JobTitle   string `json:"job_title"`
	Department string `json:"department"`
}
