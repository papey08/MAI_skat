package contacts

import "transport-api/internal/model/core"

func errorResponse(err error) contactResponse {
	if err == nil {
		return contactResponse{}
	}
	errStr := err.Error()
	return contactResponse{
		Data: nil,
		Err:  &errStr,
	}
}

type contactData struct {
	Id           uint64       `json:"id"`
	OwnerId      uint64       `json:"owner_id"`
	EmployeeId   uint64       `json:"employee_id"`
	Notes        string       `json:"notes"`
	CreationDate int64        `json:"creation_date"`
	IsDeleted    bool         `json:"is_deleted"`
	Empl         employeeData `json:"employee"`
}

func contactToContactData(contact core.Contact) contactData {
	return contactData{
		Id:           contact.Id,
		OwnerId:      contact.OwnerId,
		EmployeeId:   contact.EmployeeId,
		Notes:        contact.Notes,
		CreationDate: contact.CreationDate,
		IsDeleted:    contact.IsDeleted,
		Empl:         employeeToEmployeeData(contact.Empl),
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

type contactResponse struct {
	Data *contactData `json:"data"`
	Err  *string      `json:"error"`
}

func contactsToContactDataList(contacts []core.Contact) []contactData {
	data := make([]contactData, len(contacts))
	for i, contact := range contacts {
		data[i] = contactToContactData(contact)
	}
	return data
}

type contactListData struct {
	Contacts []contactData `json:"contacts"`
	Amount   uint          `json:"amount"`
}

type сontactListResponse struct {
	Data contactListData `json:"data"`
	Err  *string         `json:"error"`
}
