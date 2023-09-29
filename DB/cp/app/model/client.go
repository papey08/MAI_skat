package model

// Client symbolizes client of fitness club
type Client struct {
	SubscriptionID    int
	ClientSecondName  string
	ClientName        string
	ClientThirdName   string
	Sex               string
	Birthdate         string
	Height            float64
	Weight            float64
	SubscriptionBegin string
	SubscriptionEnd   string
}
