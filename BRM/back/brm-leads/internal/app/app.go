package app

import (
	"brm-leads/internal/adapters/grpcads"
	"brm-leads/internal/adapters/grpccore"
	"brm-leads/internal/adapters/grpcnotifications"
	"brm-leads/internal/app/valid"
	"brm-leads/internal/model"
	"brm-leads/internal/repo"
	"brm-leads/pkg/logger"
	"context"
	"errors"
	"time"
)

type appImpl struct {
	leadsRepo repo.LeadsRepo

	core          grpccore.CoreClient
	ads           grpcads.AdsClient
	notifications grpcnotifications.NotificationsClient

	newLeadDefaultStatus string
	leadDoneStatus       string

	logs logger.Logger
}

func (a *appImpl) CreateLead(ctx context.Context, adId uint64, clientCompany uint64, clientEmployee uint64) (lead model.Lead, err error) {
	defer func() {
		a.writeLog(logger.Fields{
			"ad_id":  adId,
			"Method": "CreateLead",
		}, err)
	}()

	data, err := a.ads.GetAdData(ctx, adId)
	if err != nil {
		return model.Lead{}, err
	}

	title, err := a.core.GetCompanyName(ctx, clientCompany)
	if err != nil {
		return model.Lead{}, err
	}

	lead = model.Lead{
		AdId:           adId,
		Title:          title,
		Description:    "",
		Price:          data.Price,
		Status:         a.newLeadDefaultStatus,
		Responsible:    data.Responsible,
		CompanyId:      data.CompanyId,
		ClientCompany:  clientCompany,
		ClientEmployee: clientEmployee,
		CreationDate:   time.Now().UTC(),
		IsDeleted:      false,
	}

	lead, err = a.leadsRepo.CreateLead(ctx, lead)
	if err == nil {
		err = a.notifications.CreateNewLeadNotification(ctx, lead.Id, lead.CompanyId, lead.ClientCompany)
	}
	return lead, err
}

func (a *appImpl) GetLeads(ctx context.Context, companyId uint64, _ uint64, filter model.Filter) (lead []model.Lead, amount uint, err error) {
	defer func() {
		a.writeLog(logger.Fields{
			"company_id": companyId,
			"Method":     "GetLeads",
		}, err)
	}()

	return a.leadsRepo.GetLeads(ctx, companyId, filter)
}

func (a *appImpl) GetLeadById(ctx context.Context, companyId uint64, _ uint64, leadId uint64) (lead model.Lead, err error) {
	defer func() {
		a.writeLog(logger.Fields{
			"lead_id": leadId,
			"Method":  "GetLeadById",
		}, err)
	}()

	lead, err = a.leadsRepo.GetLeadById(ctx, leadId)
	if err != nil {
		return model.Lead{}, err
	} else if lead.CompanyId != companyId {
		return model.Lead{}, model.ErrAuthorization
	} else {
		return lead, nil
	}
}

func (a *appImpl) UpdateLead(ctx context.Context, companyId uint64, employeeId uint64, id uint64, upd model.UpdateLead) (lead model.Lead, err error) {
	defer func() {
		a.writeLog(logger.Fields{
			"lead_id": id,
			"Method":  "UpdateLead",
		}, err)
	}()

	lead, err = a.leadsRepo.GetLeadById(ctx, id)
	if err != nil {
		return model.Lead{}, err
	} else if lead.CompanyId != companyId {
		return model.Lead{}, model.ErrAuthorization
	}

	if employeeId != lead.Responsible {
		return model.Lead{}, model.ErrAuthorization
	}

	var newResponsibleCompanyId uint64
	if newResponsibleCompanyId, _, err = a.core.GetEmployeeById(ctx, companyId, employeeId, upd.Responsible); err != nil {
		return model.Lead{}, err
	} else if newResponsibleCompanyId != lead.CompanyId {
		return model.Lead{}, model.ErrAuthorization
	}

	if upd.Status == a.leadDoneStatus && lead.Status != a.leadDoneStatus {
		err = a.notifications.CreateCloseLeadNotification(ctx, lead.AdId, lead.CompanyId, lead.ClientCompany)
		if err != nil {
			return model.Lead{}, err
		}
	}
	if !valid.UpdateLead(upd) {
		return model.Lead{}, model.ErrValidationError
	}

	return a.leadsRepo.UpdateLead(ctx, id, upd)
}

func (a *appImpl) GetStatuses(ctx context.Context) (map[string]uint64, error) {
	var err error
	defer func() {
		a.writeLog(logger.Fields{
			"Method": "GetStatuses",
		}, err)
	}()

	statuses, err := a.leadsRepo.GetStatuses(ctx)
	return statuses, err
}

func (a *appImpl) writeLog(fields logger.Fields, err error) {
	if errors.Is(err, model.ErrDatabaseError) || errors.Is(err, model.ErrAdsError) {
		a.logs.Error(fields, err.Error())
	} else if err != nil {
		a.logs.Info(fields, err.Error())
	} else {
		a.logs.Info(fields, "ok")
	}
}
