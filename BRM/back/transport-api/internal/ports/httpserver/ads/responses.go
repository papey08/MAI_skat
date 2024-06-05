package ads

import "transport-api/internal/model/ads"

func errorResponse(err error) adResponse {
	if err == nil {
		return adResponse{}
	}
	errStr := err.Error()
	return adResponse{
		Data: nil,
		Err:  &errStr,
	}
}

type adData struct {
	Id           uint64 `json:"id"`
	CompanyId    uint64 `json:"company_id"`
	Title        string `json:"title"`
	Text         string `json:"text"`
	Industry     string `json:"industry"`
	Price        uint   `json:"price"`
	ImageURL     string `json:"image_url"`
	CreationDate int64  `json:"creation_date"`
	CreatedBy    uint64 `json:"created_by"`
	Responsible  uint64 `json:"responsible"`
	IsDeleted    bool   `json:"is_deleted"`
}

func adToAdData(ad ads.Ad) adData {
	return adData{
		Id:           ad.Id,
		CompanyId:    ad.CompanyId,
		Title:        ad.Title,
		Text:         ad.Text,
		Industry:     ad.Industry,
		Price:        ad.Price,
		ImageURL:     ad.ImageURL,
		CreationDate: ad.CreationDate,
		CreatedBy:    ad.CreatedBy,
		Responsible:  ad.Responsible,
		IsDeleted:    ad.IsDeleted,
	}
}

type responseData struct {
	Id           uint64 `json:"id"`
	CompanyId    uint64 `json:"company_id"`
	EmployeeId   uint64 `json:"employee_id"`
	AdId         uint64 `json:"ad_id"`
	CreationDate int64  `json:"creation_date"`
}

func responseToResponseData(resp ads.Response) responseData {
	return responseData{
		Id:           resp.Id,
		CompanyId:    resp.CompanyId,
		EmployeeId:   resp.EmployeeId,
		AdId:         resp.AdId,
		CreationDate: resp.CreationDate,
	}
}

type adResponse struct {
	Data *adData `json:"data"`
	Err  *string `json:"error"`
}

type adListData struct {
	Ads    []adData `json:"ads"`
	Amount uint     `json:"amount"`
}

type adListResponse struct {
	Data *adListData `json:"data"`
	Err  *string     `json:"error"`
}

func adsToAdDataList(adList []ads.Ad) []adData {
	data := make([]adData, len(adList))
	for i, ad := range adList {
		data[i] = adToAdData(ad)
	}
	return data
}

type responseResponse struct {
	Data *responseData `json:"data"`
	Err  *string       `json:"error"`
}

type responseListData struct {
	Responses []responseData `json:"responses"`
	Amount    uint           `json:"amount"`
}

type responseListResponse struct {
	Data *responseListData `json:"data"`
	Err  *string           `json:"error"`
}

func responsesToResponseDataList(resps []ads.Response) []responseData {
	data := make([]responseData, len(resps))
	for i, resp := range resps {
		data[i] = responseToResponseData(resp)
	}
	return data
}

type industriesResponse struct {
	Data map[string]uint64 `json:"data"`
	Err  *string           `json:"error"`
}
