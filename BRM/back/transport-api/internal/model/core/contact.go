package core

type Contact struct {
	Id           uint64
	OwnerId      uint64
	EmployeeId   uint64
	Notes        string
	CreationDate int64
	IsDeleted    bool
	Empl         Employee
}

type UpdateContact struct {
	Notes string
}

type GetContacts struct {
	Limit  uint
	Offset uint
}
