package grpcleads

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"strings"
	"transport-api/internal/adapters/grpcleads/pb"
	"transport-api/internal/model"
	"transport-api/internal/model/leads"
)

func respToLead(lead *pb.Lead) leads.Lead {
	if lead == nil {
		return leads.Lead{}
	}
	return leads.Lead{
		Id:             lead.Id,
		AdId:           lead.AdId,
		Title:          lead.Title,
		Description:    lead.Description,
		Price:          uint(lead.Price),
		Status:         lead.Status,
		Responsible:    lead.Responsible,
		CompanyId:      lead.CompanyId,
		ClientCompany:  lead.ClientCompany,
		ClientEmployee: lead.ClientEmployee,
		CreationDate:   lead.CreationDate,
		IsDeleted:      lead.IsDeleted,
	}
}

func (l *leadsClientImpl) GetLeads(ctx context.Context, companyId uint64, employeeId uint64, filter leads.Filter) ([]leads.Lead, uint, error) {
	resp, err := l.cli.GetLeads(ctx, &pb.GetLeadsRequest{
		CompanyId:  companyId,
		EmployeeId: employeeId,
		Filter: &pb.Filter{
			Limit:         uint64(filter.Limit),
			Offset:        uint64(filter.Offset),
			Status:        filter.Status,
			ByStatus:      filter.ByStatus,
			Responsible:   filter.Responsible,
			ByResponsible: filter.ByResponsible,
		},
	})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.PermissionDenied:
			return []leads.Lead{}, 0, model.ErrPermissionDenied
		case codes.ResourceExhausted:
			return []leads.Lead{}, 0, model.ErrLeadsError
		default:
			return []leads.Lead{}, 0, model.ErrLeadsError
		}
	}

	leadsList := make([]leads.Lead, len(resp.Leads))
	for i, lead := range resp.Leads {
		leadsList[i] = respToLead(lead)
	}
	return leadsList, uint(resp.Amount), nil
}

func (l *leadsClientImpl) GetLeadById(ctx context.Context, companyId uint64, employeeId uint64, leadId uint64) (leads.Lead, error) {
	resp, err := l.cli.GetLeadById(ctx, &pb.GetLeadByIdRequest{
		CompanyId:  companyId,
		EmployeeId: employeeId,
		LeadId:     leadId,
	})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.NotFound:
			return leads.Lead{}, model.ErrLeadNotExists
		case codes.PermissionDenied:
			return leads.Lead{}, model.ErrPermissionDenied
		case codes.ResourceExhausted:
			return leads.Lead{}, model.ErrLeadsError
		default:
			return leads.Lead{}, model.ErrLeadsError
		}
	}
	return respToLead(resp.Lead), nil
}

func (l *leadsClientImpl) UpdateLead(ctx context.Context, companyId uint64, employeeId uint64, id uint64, upd leads.UpdateLead) (leads.Lead, error) {
	resp, err := l.cli.UpdateLead(ctx, &pb.UpdateLeadRequest{
		CompanyId:  companyId,
		EmployeeId: employeeId,
		Id:         id,
		Upd: &pb.UpdateLead{
			Title:       upd.Title,
			Description: upd.Description,
			Price:       uint64(upd.Price),
			Status:      upd.Status,
			Responsible: upd.Responsible,
		},
	})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.NotFound:
			// костыль, ну а чё поделать
			switch {
			case strings.Contains(err.Error(), "lead"):
				return leads.Lead{}, model.ErrLeadNotExists
			case strings.Contains(err.Error(), "employee"):
				return leads.Lead{}, model.ErrEmployeeNotExists
			case strings.Contains(err.Error(), "status"):
				return leads.Lead{}, model.ErrStatusNotExists
			}
			return leads.Lead{}, model.ErrLeadNotExists
		case codes.PermissionDenied:
			return leads.Lead{}, model.ErrPermissionDenied
		case codes.FailedPrecondition:
			return leads.Lead{}, model.ErrInvalidInput
		case codes.ResourceExhausted:
			return leads.Lead{}, model.ErrLeadsError
		default:
			return leads.Lead{}, model.ErrLeadsError
		}
	}
	return respToLead(resp.Lead), nil
}

func (l *leadsClientImpl) GetStatuses(ctx context.Context) (map[string]uint64, error) {
	resp, err := l.cli.GetStatuses(ctx, &emptypb.Empty{})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.ResourceExhausted:
			return map[string]uint64{}, model.ErrLeadsError
		default:
			return map[string]uint64{}, model.ErrLeadsError
		}
	}
	return resp.Data, nil
}
