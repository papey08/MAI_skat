package companies

import (
	"transport-api/internal/model/core"
	"transport-api/internal/model/stats"
)

func errorResponse(err error) companyResponse {
	if err == nil {
		return companyResponse{}
	}
	errStr := err.Error()
	return companyResponse{
		Data: nil,
		Err:  &errStr,
	}
}

func companyToCompanyData(company core.Company) companyData {
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

type companyResponse struct {
	Data *companyData `json:"data"`
	Err  *string      `json:"error"`
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

type mainPageResponse struct {
	Data *mainPageData `json:"data"`
	Err  *string       `json:"error"`
}

type mainPageData struct {
	ActiveLeadsAmount     uint    `json:"active_leads_amount"`
	ActiveLeadsPrice      uint    `json:"active_leads_price"`
	ClosedLeadsAmount     uint    `json:"closed_leads_amount"`
	ClosedLeadsPrice      uint    `json:"closed_leads_price"`
	ActiveAdsAmount       uint    `json:"active_ads_amount"`
	CompanyAbsoluteRating float64 `json:"company_absolute_rating"`
	CompanyRelativeRating float64 `json:"company_relative_rating"`
}

func mainPageToData(page stats.MainPage) mainPageData {
	return mainPageData{
		ActiveLeadsAmount:     page.ActiveLeadsAmount,
		ActiveLeadsPrice:      page.ActiveLeadsPrice,
		ClosedLeadsAmount:     page.ClosedLeadsAmount,
		ClosedLeadsPrice:      page.ClosedLeadsPrice,
		ActiveAdsAmount:       page.ActiveAdsAmount,
		CompanyAbsoluteRating: page.CompanyAbsoluteRating,
		CompanyRelativeRating: page.CompanyRelativeRating,
	}
}
