package app

import (
	"context"
	"transport-api/internal/model/core"
)

func (a *appImpl) CreateContact(ctx context.Context, ownerId uint64, employeeId uint64) (core.Contact, error) {
	return a.core.CreateContact(ctx, ownerId, employeeId)
}

func (a *appImpl) UpdateContact(ctx context.Context, ownerId uint64, contactId uint64, upd core.UpdateContact) (core.Contact, error) {
	return a.core.UpdateContact(ctx, ownerId, contactId, upd)
}

func (a *appImpl) DeleteContact(ctx context.Context, ownerId uint64, contactId uint64) error {
	return a.core.DeleteContact(ctx, ownerId, contactId)
}

func (a *appImpl) GetContacts(ctx context.Context, ownerId uint64, pagination core.GetContacts) ([]core.Contact, uint, error) {
	return a.core.GetContacts(ctx, ownerId, pagination)
}

func (a *appImpl) GetContactById(ctx context.Context, ownerId uint64, contactId uint64) (core.Contact, error) {
	return a.core.GetContactById(ctx, ownerId, contactId)
}
