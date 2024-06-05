package model

type Employee struct {
	Id           uint64
	CompanyId    uint64
	FirstName    string
	SecondName   string
	Email        string
	Password     string
	JobTitle     string
	Department   string
	CreationDate int64
	IsDeleted    bool
}
