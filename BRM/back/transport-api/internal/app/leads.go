package app

import (
	"context"
	"transport-api/internal/model/leads"
)

func (a *appImpl) GetLeads(ctx context.Context, companyId uint64, employeeId uint64, filter leads.Filter) ([]leads.Lead, uint, error) {
	return a.leads.GetLeads(ctx, companyId, employeeId, filter)
}

func (a *appImpl) GetLeadById(ctx context.Context, companyId uint64, employeeId uint64, leadId uint64) (leads.Lead, error) {
	return a.leads.GetLeadById(ctx, companyId, employeeId, leadId)
}

func (a *appImpl) UpdateLead(ctx context.Context, companyId uint64, employeeId uint64, id uint64, upd leads.UpdateLead) (leads.Lead, error) {
	return a.leads.UpdateLead(ctx, companyId, employeeId, id, upd)
}

func (a *appImpl) GetStatuses(ctx context.Context) (map[string]uint64, error) {
	return a.leads.GetStatuses(ctx)
}
