package leads

import "transport-api/internal/model/leads"

func errorResponse(err error) leadResponse {
	if err == nil {
		return leadResponse{}
	}
	errStr := err.Error()
	return leadResponse{
		Data: nil,
		Err:  &errStr,
	}
}

type leadData struct {
	Id             uint64 `json:"id"`
	AdId           uint64 `json:"ad_id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Price          uint   `json:"price"`
	Status         string `json:"status"`
	Responsible    uint64 `json:"responsible"`
	CompanyId      uint64 `json:"company_id"`
	ClientCompany  uint64 `json:"client_company"`
	ClientEmployee uint64 `json:"client_employee"`
	CreationDate   int64  `json:"creation_date"`
	IsDeleted      bool   `json:"is_deleted"`
}

func leadToLeadData(lead leads.Lead) leadData {
	return leadData{
		Id:             lead.Id,
		AdId:           lead.AdId,
		Title:          lead.Title,
		Description:    lead.Description,
		Price:          lead.Price,
		Status:         lead.Status,
		Responsible:    lead.Responsible,
		CompanyId:      lead.CompanyId,
		ClientCompany:  lead.ClientCompany,
		ClientEmployee: lead.ClientEmployee,
		CreationDate:   lead.CreationDate,
		IsDeleted:      lead.IsDeleted,
	}
}

type leadResponse struct {
	Data *leadData `json:"data"`
	Err  *string   `json:"error"`
}

type leadsListData struct {
	Leads  []leadData `json:"leads"`
	Amount uint       `json:"amount"`
}

type leadsListResponse struct {
	Data *leadsListData `json:"data"`
	Err  *string        `json:"error"`
}

func leadsToLeadsDataList(leadsList []leads.Lead) []leadData {
	data := make([]leadData, len(leadsList))
	for i, lead := range leadsList {
		data[i] = leadToLeadData(lead)
	}
	return data
}

type statusesResponse struct {
	Data map[string]uint64 `json:"data"`
	Err  *string           `json:"error"`
}

type statusResponse struct {
	Data string  `json:"data"`
	Err  *string `json:"error"`
}
