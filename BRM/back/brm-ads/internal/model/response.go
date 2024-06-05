package model

import "time"

type Response struct {
	Id           uint64
	CompanyId    uint64
	EmployeeId   uint64
	AdId         uint64
	CreationDate time.Time
}
