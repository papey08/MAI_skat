package repo

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"notifications/internal/model"
	"notifications/internal/model/notifications"
)

const (
	createNotificationQuery = `
		INSERT INTO "notifications" ("company_id", "date", "viewed", "type") 
		VALUES ($1, $2, $3, $4)
		RETURNING "id";`

	createNewLeadNotificationQuery = `
		INSERT INTO "new_lead_notifications" ("id", "lead_id", "client_company") 
		VALUES ($1, $2, $3);`

	createClosedLeadNotificationQuery = `
		INSERT INTO "closed_lead_notifications" ("id", "ad_id", "producer_company", "answered") 
		VALUES ($1, $2, $3, FALSE);`

	deleteNotificationQuery = `
		DELETE FROM "notifications"
		WHERE "id" = $1;`

	getNotificationsQuery = `
		SELECT * FROM "notifications"
		WHERE "company_id" = $1 AND ((NOT $2) OR (NOT "viewed"))
		ORDER BY "date" DESC
		LIMIT $3 OFFSET $4;`

	getNotificationsAmountQuery = `
		SELECT COUNT(*) FROM "notifications"
		WHERE "company_id" = $1 AND ((NOT $2) OR (NOT "viewed"));`

	getNotificationQuery = `
		SELECT * FROM "notifications"
		WHERE "id" = $1;`

	markNotificationViewed = `
		UPDATE "notifications"
		SET "viewed" = TRUE
		WHERE "id" = $1;`

	getNewLeadNotificationQuery = `
		SELECT "lead_id", "client_company" FROM "new_lead_notifications"
		WHERE "id" = $1;`

	getClosedLeadNotificationQuery = `
		SELECT "ad_id", "producer_company", "answered" FROM "closed_lead_notifications"
		WHERE "id" = $1;`

	markClosedLeadNotificationAnswered = `
		UPDATE "closed_lead_notifications"
		SET "answered" = TRUE
		WHERE "id" = $1;`
)

func (n *notificationsRepoImpl) CreateNotification(ctx context.Context, notification model.Notification) error {
	var notificationId uint64
	if err := n.QueryRow(ctx, createNotificationQuery,
		notification.CompanyId,
		notification.Date,
		notification.Viewed,
		notification.Type,
	).Scan(&notificationId); err != nil {
		return errors.Join(model.ErrDatabaseError, err)
	}

	var err error
	switch notification.Type {
	case model.NewLead:
		_, err = n.Exec(ctx, createNewLeadNotificationQuery,
			notificationId,
			notification.NewLead.LeadId,
			notification.NewLead.ClientCompany,
		)
	case model.ClosedLead:
		_, err = n.Exec(ctx, createClosedLeadNotificationQuery,
			notificationId,
			notification.ClosedLead.AdId,
			notification.ClosedLead.ProducerCompany,
		)
	default:
		_, _ = n.Exec(ctx, deleteNotificationQuery, notificationId)
		return model.ErrServiceError
	}
	if err != nil {
		return errors.Join(model.ErrDatabaseError, err)
	}
	return nil
}

func (n *notificationsRepoImpl) GetNotifications(ctx context.Context, companyId uint64, limit uint, offset uint, onlyNotViewed bool) ([]model.Notification, uint, error) {
	var amount uint
	if err := n.QueryRow(ctx, getNotificationsAmountQuery,
		companyId,
		onlyNotViewed,
	).Scan(&amount); err != nil {
		return []model.Notification{}, 0, errors.Join(model.ErrDatabaseError, err)
	} else if amount == 0 {
		return []model.Notification{}, 0, nil
	}

	rows, err := n.Query(ctx, getNotificationsQuery,
		companyId,
		onlyNotViewed,
		limit,
		offset,
	)
	if err != nil {
		return []model.Notification{}, 0, errors.Join(model.ErrDatabaseError, err)
	}
	defer rows.Close()

	notificationsList := make([]model.Notification, 0)
	for rows.Next() {
		var notification model.Notification
		_ = rows.Scan(
			&notification.Id,
			&notification.CompanyId,
			&notification.Date,
			&notification.Viewed,
			&notification.Type,
		)
		notificationsList = append(notificationsList, notification)
	}
	return notificationsList, amount, nil
}

func (n *notificationsRepoImpl) GetNotification(ctx context.Context, id uint64) (model.Notification, error) {
	var notification model.Notification
	if err := n.QueryRow(ctx, getNotificationQuery, id).Scan(
		&notification.Id,
		&notification.CompanyId,
		&notification.Date,
		&notification.Viewed,
		&notification.Type,
	); errors.Is(err, pgx.ErrNoRows) {
		return model.Notification{}, model.ErrNotificationNotFound
	} else if err != nil {
		return model.Notification{}, errors.Join(model.ErrDatabaseError, err)
	}

	var err error
	switch notification.Type {
	case model.NewLead:
		notification.NewLead = new(notifications.NewLead)
		err = n.QueryRow(ctx, getNewLeadNotificationQuery, id).Scan(
			&notification.NewLead.LeadId,
			&notification.NewLead.ClientCompany,
		)
	case model.ClosedLead:
		notification.ClosedLead = new(notifications.ClosedLead)
		err = n.QueryRow(ctx, getClosedLeadNotificationQuery, id).Scan(
			&notification.ClosedLead.AdId,
			&notification.ClosedLead.ProducerCompany,
			&notification.ClosedLead.Answered,
		)
	default:
		return model.Notification{}, model.ErrServiceError
	}
	if err != nil {
		return model.Notification{}, errors.Join(model.ErrDatabaseError, err)
	}

	_, _ = n.Exec(ctx, markNotificationViewed, id)

	return notification, nil
}

func (n *notificationsRepoImpl) MarkClosedLeadNotificationAnswered(ctx context.Context, notificationId uint64) error {
	if e, err := n.Exec(ctx, markClosedLeadNotificationAnswered, notificationId); err != nil {
		return errors.Join(model.ErrDatabaseError, err)
	} else if e.RowsAffected() == 0 {
		return model.ErrNotificationNotFound
	}
	return nil
}
