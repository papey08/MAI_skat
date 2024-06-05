package httpserver

import "registration/internal/model"

func errorResponse(err error) companyAndOwnerResponse {
	if err == nil {
		return companyAndOwnerResponse{}
	}
	errStr := err.Error()
	return companyAndOwnerResponse{
		Data: nil,
		Err:  &errStr,
	}
}

func companyToCompanyData(company model.Company) companyData {
	return companyData{
		Id:           company.Id,
		Name:         company.Name,
		Description:  company.Description,
		Industry:     company.Industry,
		OwnerId:      company.OwnerId,
		Rating:       company.Rating,
		CreationDate: company.CreationDate,
		IsDeleted:    company.IsDeleted,
	}
}

func ownerToOwnerData(owner model.Employee) ownerData {
	return ownerData{
		Id:           owner.Id,
		CompanyId:    owner.CompanyId,
		FirstName:    owner.FirstName,
		SecondName:   owner.SecondName,
		Email:        owner.Email,
		JobTitle:     owner.JobTitle,
		Department:   owner.Department,
		CreationDate: owner.CreationDate,
		IsDeleted:    owner.IsDeleted,
	}
}

type companyAndOwnerResponse struct {
	Data *companyAndOwnerData `json:"data"`
	Err  *string              `json:"err"`
}

type companyAndOwnerData struct {
	Company companyData `json:"company"`
	Owner   ownerData   `json:"owner"`
}

type ownerData struct {
	Id           uint64 `json:"id"`
	CompanyId    uint64 `json:"company_id"`
	FirstName    string `json:"first_name"`
	SecondName   string `json:"second_name"`
	Email        string `json:"email"`
	JobTitle     string `json:"job_title"`
	Department   string `json:"department"`
	CreationDate int64  `json:"creation_date"`
	IsDeleted    bool   `json:"is_deleted"`
}

type companyData struct {
	Id           uint64  `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Industry     string  `json:"industry"`
	OwnerId      uint64  `json:"owner_id"`
	Rating       float64 `json:"rating"`
	CreationDate int64   `json:"creation_date"`
	IsDeleted    bool    `json:"is_deleted"`
}

type industriesResponse struct {
	Data map[string]uint64 `json:"data"`
	Err  *string           `json:"error"`
}
