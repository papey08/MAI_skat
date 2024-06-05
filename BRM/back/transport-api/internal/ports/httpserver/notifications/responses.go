package notifications

import "transport-api/internal/model/notifications"

func errorResponse(err error) notificationResponse {
	if err == nil {
		return notificationResponse{}
	}
	errStr := err.Error()
	return notificationResponse{
		Data: nil,
		Err:  &errStr,
	}
}

func notificationToNotificationData(notification notifications.Notification) notificationData {
	var res notificationData
	res.Id = notification.Id
	res.CompanyId = notification.CompanyId
	res.Date = notification.Date
	res.Type = notification.Type
	res.Viewed = notification.Viewed
	if notification.NewLead != nil {
		res.NewLeadInfo = new(newLeadData)
		res.NewLeadInfo.LeadId = notification.NewLead.LeadId
		res.NewLeadInfo.ClientCompany = notification.NewLead.ClientCompany
	}
	if notification.ClosedLead != nil {
		res.ClosedLeadInfo = new(closedLeadData)
		res.ClosedLeadInfo.AdId = notification.ClosedLead.AdId
		res.ClosedLeadInfo.ProducerCompany = notification.ClosedLead.ProducerCompany
		res.ClosedLeadInfo.Answered = notification.ClosedLead.Answered
	}
	return res
}

func notificationsToNotificationsDataList(notifications []notifications.Notification) []notificationData {
	data := make([]notificationData, len(notifications))
	for i, notification := range notifications {
		data[i] = notificationToNotificationData(notification)
	}
	return data
}

type notificationResponse struct {
	Data *notificationData `json:"data"`
	Err  *string           `json:"error"`
}

type notificationListResponse struct {
	Data *notificationListData `json:"data"`
	Err  *string               `json:"error"`
}

type notificationListData struct {
	Notifications []notificationData `json:"notifications"`
	Amount        uint               `json:"amount"`
}

type notificationData struct {
	Id        uint64 `json:"id"`
	CompanyId uint64 `json:"company_id"`
	Type      string `json:"type"`
	Date      int64  `json:"date"`
	Viewed    bool   `json:"viewed"`

	NewLeadInfo    *newLeadData    `json:"new_lead_info,omitempty"`
	ClosedLeadInfo *closedLeadData `json:"closed_lead_info,omitempty"`
}

type closedLeadData struct {
	AdId            uint64 `json:"ad_id"`
	ProducerCompany uint64 `json:"producer_company"`
	Answered        bool   `json:"answered"`
}

type newLeadData struct {
	LeadId        uint64 `json:"lead_id"`
	ClientCompany uint64 `json:"client_company"`
}
