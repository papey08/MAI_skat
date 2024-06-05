package model

import (
	"notifications/internal/model/notifications"
	"time"
)

type Notification struct {
	Id        uint64
	CompanyId uint64
	Type      NotificationType
	Date      time.Time
	Viewed    bool

	*notifications.NewLead
	*notifications.ClosedLead
}

type NotificationType string

const (
	Unknown    NotificationType = ""
	NewLead    NotificationType = "new_lead"
	ClosedLead NotificationType = "closed_lead"
)
