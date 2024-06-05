package employees

import "transport-api/internal/model/core"

func errorResponse(err error) employeeResponse {
	if err == nil {
		return employeeResponse{}
	}
	errStr := err.Error()
	return employeeResponse{
		Data: nil,
		Err:  &errStr,
	}
}

func employeeToEmployeeData(employee core.Employee) employeeData {
	return employeeData{
		Id:           employee.Id,
		CompanyId:    employee.CompanyId,
		FirstName:    employee.FirstName,
		SecondName:   employee.SecondName,
		Email:        employee.Email,
		JobTitle:     employee.JobTitle,
		Department:   employee.Department,
		ImageURL:     employee.ImageURL,
		CreationDate: employee.CreationDate,
		IsDeleted:    employee.IsDeleted,
	}
}

func employeesToEmployeeDataList(employees []core.Employee) []employeeData {
	data := make([]employeeData, len(employees))
	for i, employee := range employees {
		data[i] = employeeToEmployeeData(employee)
	}
	return data
}

type employeeData struct {
	Id           uint64 `json:"id"`
	CompanyId    uint64 `json:"company_id"`
	FirstName    string `json:"first_name"`
	SecondName   string `json:"second_name"`
	Email        string `json:"email"`
	JobTitle     string `json:"job_title"`
	Department   string `json:"department"`
	ImageURL     string `json:"image_url"`
	CreationDate int64  `json:"creation_date"`
	IsDeleted    bool   `json:"is_deleted"`
}

type employeeResponse struct {
	Data *employeeData `json:"data"`
	Err  *string       `json:"error"`
}

type employeeListData struct {
	Employees []employeeData `json:"employees"`
	Amount    uint           `json:"amount"`
}

type employeeListResponse struct {
	Data *employeeListData `json:"data"`
	Err  *string           `json:"error"`
}
