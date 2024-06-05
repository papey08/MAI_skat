package app

import (
	"brm-core/internal/adapters/grpcauth"
	"brm-core/internal/app/valid"
	"brm-core/internal/model"
	"brm-core/internal/repo"
	"brm-core/pkg/logger"
	"context"
	"errors"
	"time"
)

type appImpl struct {
	coreRepo repo.CoreRepo
	auth     grpcauth.AuthClient

	logs logger.Logger
}

func (a *appImpl) GetCompany(ctx context.Context, id uint64) (company model.Company, err error) {
	defer func() {
		a.writeLog(logger.Fields{
			"company_id": id,
			"Method":     "GetCompany",
		}, err)
	}()
	return a.coreRepo.GetCompany(ctx, id)
}

func (a *appImpl) CreateCompanyAndOwner(ctx context.Context, company model.Company, owner model.Employee) (comp model.Company, empl model.Employee, err error) {
	defer func() {
		a.writeLog(logger.Fields{
			"company_name": company.Name,
			"owner_email":  owner.Email,
			"Method":       "CreateCompanyAndOwner",
		}, err)
	}()

	if err != nil {
		return model.Company{}, model.Employee{}, errors.Join(model.ErrDatabaseError, err)
	}

	// setting company fields
	company.Id = 0
	company.OwnerId = 0
	company.Rating = 4.
	company.CreationDate = time.Now().UTC()
	company.IsDeleted = false

	// setting owner fields
	owner.Id = 0
	owner.CompanyId = 0
	owner.CreationDate = time.Now().UTC()
	owner.IsDeleted = false

	if !(valid.CreateCompany(company) && valid.CreateEmployee(owner)) {
		err = model.ErrValidationError
		return model.Company{}, model.Employee{}, err
	}

	newCompany, newOwner, err := a.coreRepo.CreateCompanyAndOwner(ctx, company, owner)
	if err != nil {
		return model.Company{}, model.Employee{}, err
	}

	if err = a.auth.RegisterEmployee(ctx, model.EmployeeCredentials{
		Email:      newOwner.Email,
		Password:   newOwner.Password,
		EmployeeId: newOwner.Id,
		CompanyId:  newOwner.CompanyId,
	}); err != nil {
		return model.Company{}, model.Employee{}, err
	}

	return newCompany, newOwner, nil
}

func (a *appImpl) UpdateCompany(ctx context.Context, companyId uint64, ownerId uint64, upd model.UpdateCompany) (comp model.Company, err error) {
	defer func() {
		a.writeLog(logger.Fields{
			"company_id": companyId,
			"owner_id":   ownerId,
			"Method":     "UpdateCompany",
		}, err)
	}()

	comp, err = a.coreRepo.GetCompany(ctx, companyId)
	if err != nil {
		return model.Company{}, err
	} else if comp.OwnerId != ownerId {
		return model.Company{}, model.ErrAuthorization
	}

	if comp.OwnerId != upd.OwnerId {
		var newOwner model.Employee
		newOwner, err = a.coreRepo.GetEmployeeById(ctx, upd.OwnerId)
		if err != nil {
			return model.Company{}, err
		}

		if newOwner.CompanyId != companyId {
			return model.Company{}, model.ErrEmployeeNotExists
		}
	}

	if !valid.UpdateCompany(upd) {
		return model.Company{}, model.ErrValidationError
	}

	return a.coreRepo.UpdateCompany(ctx, companyId, upd)
}

func (a *appImpl) GetIndustries(ctx context.Context) (map[string]uint64, error) {
	return a.coreRepo.GetIndustries(ctx)
}

func (a *appImpl) CreateEmployee(ctx context.Context, companyId uint64, ownerId uint64, employee model.Employee) (empl model.Employee, err error) {
	defer func() {
		a.writeLog(logger.Fields{
			"company_id":         companyId,
			"owner_id":           ownerId,
			"new_employee_email": employee.Email,
			"method":             "CreateEmployee",
		}, err)
	}()

	if companyId != employee.CompanyId {
		return model.Employee{}, model.ErrAuthorization
	}

	company, err := a.coreRepo.GetCompany(ctx, companyId)
	if err != nil {
		return model.Employee{}, err
	} else if company.OwnerId != ownerId {
		return model.Employee{}, model.ErrAuthorization
	}

	// setting up employee fields
	employee.Id = 0
	employee.CreationDate = time.Now().UTC()
	employee.IsDeleted = false

	if !valid.CreateEmployee(employee) {
		return model.Employee{}, model.ErrValidationError
	}

	newEmployee, err := a.coreRepo.CreateEmployee(ctx, employee)
	if err != nil {
		return model.Employee{}, err
	}

	err = a.auth.RegisterEmployee(ctx, model.EmployeeCredentials{
		Email:      newEmployee.Email,
		Password:   newEmployee.Password,
		EmployeeId: newEmployee.Id,
		CompanyId:  newEmployee.CompanyId,
	})
	if err != nil {
		return model.Employee{}, err
	}

	return newEmployee, nil
}

func (a *appImpl) UpdateEmployee(ctx context.Context, companyId uint64, ownerId uint64, employeeId uint64, upd model.UpdateEmployee) (empl model.Employee, err error) {
	defer func() {
		a.writeLog(logger.Fields{
			"company_id": companyId,
			"owner_id":   ownerId,
			"method":     "UpdateEmployee",
		}, err)
	}()

	employee, err := a.coreRepo.GetEmployeeById(ctx, employeeId)
	if err != nil {
		return model.Employee{}, err
	} else if companyId != employee.CompanyId {
		return model.Employee{}, model.ErrAuthorization
	}

	company, err := a.coreRepo.GetCompany(ctx, companyId)
	if err != nil {
		return model.Employee{}, err
	} else if company.OwnerId != ownerId {
		return model.Employee{}, model.ErrAuthorization
	}

	if !valid.UpdateEmployee(upd) {
		return model.Employee{}, model.ErrValidationError
	}

	return a.coreRepo.UpdateEmployee(ctx, employeeId, upd)
}

func (a *appImpl) DeleteEmployee(ctx context.Context, companyId uint64, ownerId uint64, employeeId uint64) (err error) {
	defer func() {
		a.writeLog(logger.Fields{
			"company_id": companyId,
			"owner_id":   ownerId,
			"Method":     "DeleteEmployee",
		}, err)
	}()

	if ownerId == employeeId {
		return model.ErrOwnerDeletion
	}

	employee, err := a.coreRepo.GetEmployeeById(ctx, employeeId)
	if err != nil {
		return err
	} else if companyId != employee.CompanyId {
		return model.ErrAuthorization
	}

	company, err := a.coreRepo.GetCompany(ctx, companyId)
	if err != nil {
		return err
	} else if company.OwnerId != ownerId {
		return model.ErrAuthorization
	}

	err = a.coreRepo.DeleteEmployee(ctx, employeeId)
	if err != nil {
		return err
	}
	return a.auth.DeleteEmployee(ctx, employee.Email)
}

func (a *appImpl) GetCompanyEmployees(ctx context.Context, companyId uint64, employeeId uint64, filter model.FilterEmployee) (empl []model.Employee, amount uint, err error) {
	defer func() {
		a.writeLog(logger.Fields{
			"company_id":  companyId,
			"employee_id": employeeId,
			"Method":      "GetCompanyEmployees",
		}, err)
	}()

	_, err = a.coreRepo.GetCompany(ctx, companyId)
	if err != nil {
		return []model.Employee{}, 0, err
	}
	employee, err := a.coreRepo.GetEmployeeById(ctx, employeeId)
	if err != nil {
		return []model.Employee{}, 0, err
	} else if companyId != employee.CompanyId {
		return []model.Employee{}, 0, model.ErrAuthorization
	}

	return a.coreRepo.GetCompanyEmployees(ctx, companyId, filter)
}

func (a *appImpl) GetEmployeeByName(ctx context.Context, companyId uint64, employeeId uint64, ebn model.EmployeeByName) (empl []model.Employee, amount uint, err error) {
	defer func() {
		a.writeLog(logger.Fields{
			"company_id":  companyId,
			"employee_id": employeeId,
			"Method":      "GetEmployeeByName",
		}, err)
	}()

	_, err = a.coreRepo.GetCompany(ctx, companyId)
	if err != nil {
		return []model.Employee{}, 0, err
	}
	employee, err := a.coreRepo.GetEmployeeById(ctx, employeeId)
	if err != nil {
		return []model.Employee{}, 0, err
	} else if companyId != employee.CompanyId {
		return []model.Employee{}, 0, model.ErrAuthorization
	}

	return a.coreRepo.GetEmployeeByName(ctx, companyId, ebn)
}

func (a *appImpl) GetEmployeeById(ctx context.Context, companyId uint64, employeeId uint64, employeeIdToFind uint64) (empl model.Employee, err error) {
	defer func() {
		a.writeLog(logger.Fields{
			"company_id":  companyId,
			"employee_id": employeeId,
			"Method":      "GetEmployeeById",
		}, err)
	}()

	_, err = a.coreRepo.GetCompany(ctx, companyId)
	if err != nil {
		return model.Employee{}, err
	}
	employee, err := a.coreRepo.GetEmployeeById(ctx, employeeIdToFind)
	if err != nil {
		return model.Employee{}, err
	} else if companyId != employee.CompanyId {
		return model.Employee{
			Id:           employee.Id,
			CompanyId:    employee.CompanyId,
			FirstName:    employee.FirstName,
			SecondName:   employee.SecondName,
			Email:        "",
			Password:     "",
			JobTitle:     "",
			Department:   "",
			ImageURL:     employee.ImageURL,
			CreationDate: employee.CreationDate,
			IsDeleted:    employee.IsDeleted,
		}, nil
	}
	return employee, err
}

func (a *appImpl) CreateContact(ctx context.Context, ownerId uint64, employeeId uint64) (cnt model.Contact, err error) {
	defer func() {
		a.writeLog(logger.Fields{
			"contact_owner_id":         ownerId,
			"employee_id_from_contact": employeeId,
			"Method":                   "CreateContact",
		}, err)
	}()

	_, err = a.coreRepo.GetEmployeeById(ctx, ownerId)
	if err != nil {
		return model.Contact{}, err
	}

	_, err = a.coreRepo.GetEmployeeById(ctx, employeeId)
	if err != nil {
		return model.Contact{}, err
	}

	if ownerId == employeeId {
		return model.Contact{}, model.ErrSelfContact
	}

	return a.coreRepo.CreateContact(ctx, model.Contact{
		Id:           0,
		OwnerId:      ownerId,
		EmployeeId:   employeeId,
		Notes:        "",
		CreationDate: time.Now().UTC(),
		IsDeleted:    false,
		Empl:         model.Employee{},
	})
}

func (a *appImpl) UpdateContact(ctx context.Context, ownerId uint64, contactId uint64, upd model.UpdateContact) (cnt model.Contact, err error) {
	defer func() {
		a.writeLog(logger.Fields{
			"contact_owner_id": ownerId,
			"contact_id":       contactId,
			"Method":           "UpdateContact",
		}, err)
	}()

	_, err = a.coreRepo.GetEmployeeById(ctx, ownerId)
	if err != nil {
		return model.Contact{}, err
	}

	contact, err := a.coreRepo.GetContactById(ctx, ownerId, contactId)
	if err != nil {
		return model.Contact{}, err
	} else if contact.OwnerId != ownerId {
		return model.Contact{}, err
	}

	if !valid.UpdateContact(upd) {
		return model.Contact{}, model.ErrValidationError
	}

	return a.coreRepo.UpdateContact(ctx, ownerId, contactId, upd)
}

func (a *appImpl) DeleteContact(ctx context.Context, ownerId uint64, contactId uint64) (err error) {
	defer func() {
		a.writeLog(logger.Fields{
			"contact_owner_id": ownerId,
			"contact_id":       contactId,
			"Method":           "DeleteContact",
		}, err)
	}()

	_, err = a.coreRepo.GetEmployeeById(ctx, ownerId)
	if err != nil {
		return err
	}

	contact, err := a.coreRepo.GetContactById(ctx, ownerId, contactId)
	if err != nil {
		return err
	} else if contact.OwnerId != ownerId {
		return model.ErrAuthorization
	}

	return a.coreRepo.DeleteContact(ctx, ownerId, contactId)
}

func (a *appImpl) GetContacts(ctx context.Context, ownerId uint64, pagination model.GetContacts) (cnt []model.Contact, amount uint, err error) {
	defer func() {
		a.writeLog(logger.Fields{
			"contact_owner_id": ownerId,
			"Method":           "GetContacts",
		}, err)
	}()

	_, err = a.coreRepo.GetEmployeeById(ctx, ownerId)
	if err != nil {
		return []model.Contact{}, 0, err
	}

	amount, err = a.coreRepo.GetContactsAmount(ctx, ownerId)
	if err != nil || amount == 0 {
		return []model.Contact{}, 0, err
	}

	cnt, err = a.coreRepo.GetContacts(ctx, ownerId, pagination)
	return cnt, amount, err
}

func (a *appImpl) GetContactById(ctx context.Context, ownerId uint64, contactId uint64) (cnt model.Contact, err error) {
	defer func() {
		a.writeLog(logger.Fields{
			"contact_owner_id": ownerId,
			"contact_id":       contactId,
			"Method":           "GetContactById",
		}, err)
	}()

	_, err = a.coreRepo.GetEmployeeById(ctx, ownerId)
	if err != nil {
		return model.Contact{}, err
	}

	cnt, err = a.coreRepo.GetContactById(ctx, ownerId, contactId)
	if err != nil {
		return model.Contact{}, err
	} else if cnt.OwnerId != ownerId {
		return model.Contact{}, err
	}

	return a.coreRepo.GetContactById(ctx, ownerId, contactId)
}

func (a *appImpl) writeLog(fields logger.Fields, err error) {
	if errors.Is(err, model.ErrDatabaseError) || errors.Is(err, model.ErrAuthServiceError) {
		a.logs.Error(fields, err.Error())
	} else if err != nil {
		a.logs.Info(fields, err.Error())
	} else {
		a.logs.Info(fields, "ok")
	}
}
