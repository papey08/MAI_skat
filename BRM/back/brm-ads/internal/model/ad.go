package model

import "time"

type Ad struct {
	Id           uint64
	CompanyId    uint64
	Title        string
	Text         string
	Industry     string
	Price        uint
	ImageURL     string
	CreationDate time.Time
	CreatedBy    uint64
	Responsible  uint64
	IsDeleted    bool
}

type UpdateAd struct {
	Title       string
	Text        string
	Industry    string
	Price       uint
	ImageURL    string
	Responsible uint64
}

type AdsListParams struct {
	Search *AdSearcher
	Filter *AdFilter
	Sort   *AdSorter
	Limit  uint
	Offset uint
}

type AdSearcher struct {
	Pattern string
}

type AdFilter struct {
	ByCompany bool
	CompanyId uint64

	ByIndustry bool
	Industry   string
}

type AdSorter struct {
	ByPriceAsc  bool
	ByPriceDesc bool

	ByDateAsc  bool
	ByDateDesc bool
}
