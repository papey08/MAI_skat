package repo

import (
	"brm-core/internal/model"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

const (
	createContactQuery = `
		INSERT INTO "contacts" ("owner_id", "employee_id", "notes", "creation_date", "is_deleted") 
		VALUES ($1, $2, $3, $4, $5)
		RETURNING "id";`

	updateContactQuery = `
		UPDATE "contacts"
		SET "notes" = $2
		WHERE "id" = $1 AND (NOT "is_deleted");`

	deleteContactQuery = `
		UPDATE "contacts"
		SET "is_deleted" = true
		WHERE "id" = $1 AND (NOT "is_deleted");`

	getContactsQuery = `
		SELECT * FROM "contacts"
		INNER JOIN "employees" ON "employee_id" = "employees"."id"
		WHERE "owner_id" = $1 AND (NOT "contacts"."is_deleted")
		LIMIT $2 OFFSET $3;`

	getContactsAmountQuery = `
		SELECT COUNT(*) FROM "contacts"
		WHERE "owner_id" = $1 AND (NOT "is_deleted");`

	getContactByIdQuery = `
		SELECT * FROM "contacts"
		INNER JOIN "employees" ON "employee_id" = "employees"."id"
		WHERE "contacts"."id" = $1 AND (NOT "contacts"."is_deleted");`
)

func (c *coreRepoImpl) CreateContact(ctx context.Context, contact model.Contact) (model.Contact, error) {
	var contactId uint64
	var pgErr *pgconn.PgError
	if err := c.QueryRow(ctx, createContactQuery,
		contact.OwnerId,
		contact.EmployeeId,
		contact.Notes,
		contact.CreationDate,
		contact.IsDeleted,
	).Scan(&contactId); errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505": // duplicate primary key error
			return model.Contact{}, model.ErrContactExist
		default:
			return model.Contact{}, model.ErrServiceError
		}
	} else if err != nil {
		return model.Contact{}, errors.Join(model.ErrDatabaseError, err)
	}

	contact.Id = contactId

	if employee, err := c.GetEmployeeById(ctx, contact.EmployeeId); err != nil {
		return model.Contact{}, err
	} else {
		contact.Empl = employee
	}

	return contact, nil
}

func (c *coreRepoImpl) UpdateContact(ctx context.Context, ownerId uint64, contactId uint64, upd model.UpdateContact) (model.Contact, error) {
	if e, err := c.Exec(ctx, updateContactQuery,
		contactId,
		upd.Notes,
	); err != nil {
		return model.Contact{}, errors.Join(model.ErrDatabaseError, err)
	} else if e.RowsAffected() == 0 {
		return model.Contact{}, model.ErrContactNotExists
	}

	return c.GetContactById(ctx, ownerId, contactId)
}

func (c *coreRepoImpl) DeleteContact(ctx context.Context, _ uint64, contactId uint64) error {
	if e, err := c.Exec(ctx, deleteContactQuery,
		contactId,
	); err != nil {
		return errors.Join(model.ErrDatabaseError, err)
	} else if e.RowsAffected() == 0 {
		return model.ErrContactNotExists
	} else {
		return nil
	}
}

func (c *coreRepoImpl) GetContacts(ctx context.Context, ownerId uint64, pagination model.GetContacts) ([]model.Contact, error) {
	rows, err := c.Query(ctx, getContactsQuery,
		ownerId,
		pagination.Limit,
		pagination.Offset,
	)
	if err != nil {
		return []model.Contact{}, errors.Join(model.ErrDatabaseError, err)
	}
	defer rows.Close()

	contacts := make([]model.Contact, 0)
	for rows.Next() {
		var contact model.Contact
		_ = rows.Scan(
			&contact.Id,
			&contact.OwnerId,
			&contact.EmployeeId,
			&contact.Notes,
			&contact.CreationDate,
			&contact.IsDeleted,
			&contact.Empl.Id,
			&contact.Empl.CompanyId,
			&contact.Empl.FirstName,
			&contact.Empl.SecondName,
			&contact.Empl.Email,
			&contact.Empl.JobTitle,
			&contact.Empl.Department,
			&contact.Empl.ImageURL,
			&contact.Empl.CreationDate,
			&contact.Empl.IsDeleted,
		)

		contacts = append(contacts, contact)
	}
	return contacts, nil
}

func (c *coreRepoImpl) GetContactById(ctx context.Context, _ uint64, contactId uint64) (model.Contact, error) {
	row := c.QueryRow(ctx, getContactByIdQuery,
		contactId,
	)
	var contact model.Contact
	if err := row.Scan(
		&contact.Id,
		&contact.OwnerId,
		&contact.EmployeeId,
		&contact.Notes,
		&contact.CreationDate,
		&contact.IsDeleted,
		&contact.Empl.Id,
		&contact.Empl.CompanyId,
		&contact.Empl.FirstName,
		&contact.Empl.SecondName,
		&contact.Empl.Email,
		&contact.Empl.JobTitle,
		&contact.Empl.Department,
		&contact.Empl.ImageURL,
		&contact.Empl.CreationDate,
		&contact.Empl.IsDeleted,
	); errors.Is(err, pgx.ErrNoRows) {
		return model.Contact{}, model.ErrContactNotExists
	} else if err != nil {
		return model.Contact{}, errors.Join(model.ErrDatabaseError, err)
	}

	return contact, nil
}

func (c *coreRepoImpl) GetContactsAmount(ctx context.Context, ownerId uint64) (uint, error) {
	var amount uint
	if err := c.QueryRow(ctx, getContactsAmountQuery, ownerId).Scan(&amount); err != nil {
		return 0, errors.Join(model.ErrDatabaseError, err)
	}
	return amount, nil
}
