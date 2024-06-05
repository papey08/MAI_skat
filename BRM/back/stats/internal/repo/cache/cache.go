package cache

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"stats/internal/model"
	"strconv"
	"time"
)

type cacheImpl struct {
	expirationTime time.Duration
	*redis.Client
}

func (c *cacheImpl) AddCompanyData(ctx context.Context, companyId uint64, data model.MainPageStats) {
	_ = c.Set(ctx, strconv.FormatUint(companyId, 10), dataToJSON(data), c.expirationTime)
}

func (c *cacheImpl) GetCompanyData(ctx context.Context, companyId uint64) (model.MainPageStats, bool) {
	if jsonData, err := c.Get(ctx, strconv.FormatUint(companyId, 10)).Result(); err != nil {
		return model.MainPageStats{}, false
	} else {
		return jsonToData([]byte(jsonData)), true
	}
}

func dataToJSON(data model.MainPageStats) []byte {
	jsonData, _ := json.Marshal(companyData{
		ActiveLeadsAmount:     data.ActiveLeadsAmount,
		ActiveLeadsPrice:      data.ActiveLeadsPrice,
		ClosedLeadsAmount:     data.ClosedLeadsAmount,
		ClosedLeadsPrice:      data.ClosedLeadsPrice,
		ActiveAdsAmount:       data.ActiveAdsAmount,
		CompanyAbsoluteRating: data.CompanyAbsoluteRating,
		CompanyRelativeRating: data.CompanyRelativeRating,
	})
	return jsonData
}

func jsonToData(jsonData []byte) model.MainPageStats {
	var data companyData
	_ = json.Unmarshal(jsonData, &data)
	return model.MainPageStats{
		ActiveLeadsAmount:     data.ActiveLeadsAmount,
		ActiveLeadsPrice:      data.ActiveLeadsPrice,
		ClosedLeadsAmount:     data.ClosedLeadsAmount,
		ClosedLeadsPrice:      data.ClosedLeadsPrice,
		ActiveAdsAmount:       data.ActiveAdsAmount,
		CompanyAbsoluteRating: data.CompanyAbsoluteRating,
		CompanyRelativeRating: data.CompanyRelativeRating,
	}
}

type companyData struct {
	ActiveLeadsAmount     uint    `json:"active_leads_amount"`
	ActiveLeadsPrice      uint    `json:"active_leads_price"`
	ClosedLeadsAmount     uint    `json:"closed_leads_amount"`
	ClosedLeadsPrice      uint    `json:"closed_leads_price"`
	ActiveAdsAmount       uint    `json:"active_ads_amount"`
	CompanyAbsoluteRating float64 `json:"company_absolute_rating"`
	CompanyRelativeRating float64 `json:"company_relative_rating"`
}
