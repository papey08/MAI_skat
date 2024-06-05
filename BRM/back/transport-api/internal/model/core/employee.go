package core

type Employee struct {
	Id           uint64
	CompanyId    uint64
	FirstName    string
	SecondName   string
	Email        string
	Password     string
	JobTitle     string
	Department   string
	ImageURL     string
	CreationDate int64
	IsDeleted    bool
}

type UpdateEmployee struct {
	FirstName  string
	SecondName string
	JobTitle   string
	Department string
	ImageURL   string
}

type FilterEmployee struct {
	ByJobTitle bool
	JobTitle   string

	ByDepartment bool
	Department   string

	Limit  uint
	Offset uint
}

type EmployeeByName struct {
	Pattern string

	Limit  uint
	Offset uint
}
