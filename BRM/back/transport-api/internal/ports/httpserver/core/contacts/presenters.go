package contacts

type addContactRequest struct {
	// EmployeeId это id сотрудника, которого добавляют в контакты
	EmployeeId uint64 `json:"employee_id"`
}

type updateContactRequest struct {
	Notes string `json:"notes"`
}
