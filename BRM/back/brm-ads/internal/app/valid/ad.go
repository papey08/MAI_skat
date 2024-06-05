package valid

import (
	"brm-ads/internal/model"
	"net/url"
)

func CreateAd(ad model.Ad) bool {
	isValidTitle := len([]rune(ad.Title)) <= 200
	isValidText := len([]rune(ad.Text)) <= 1000

	validImageUrl := true
	if ad.ImageURL != "" {
		_, err := url.ParseRequestURI(ad.ImageURL)
		validImageUrl = err == nil && len(ad.ImageURL) <= 200
	}
	return isValidTitle &&
		isValidText &&
		validImageUrl
}

func UpdateAd(upd model.UpdateAd) bool {
	isValidTitle := len([]rune(upd.Title)) <= 200
	isValidText := len([]rune(upd.Text)) <= 1000

	validImageUrl := true
	if upd.ImageURL != "" {
		_, err := url.ParseRequestURI(upd.ImageURL)
		validImageUrl = err == nil && len(upd.ImageURL) <= 200
	}

	return isValidTitle &&
		isValidText &&
		validImageUrl
}
