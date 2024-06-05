package grpcserver

import (
	"brm-leads/internal/model"
	"brm-leads/internal/ports/grpcserver/pb"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

func modelLeadToLead(lead model.Lead) *pb.Lead {
	if lead.Id == 0 {
		return nil
	}
	return &pb.Lead{
		Id:             lead.Id,
		AdId:           lead.AdId,
		Title:          lead.Title,
		Description:    lead.Description,
		Price:          uint64(lead.Price),
		Status:         lead.Status,
		Responsible:    lead.Responsible,
		CompanyId:      lead.CompanyId,
		ClientCompany:  lead.ClientCompany,
		ClientEmployee: lead.ClientEmployee,
		CreationDate:   lead.CreationDate.UTC().Unix(),
		IsDeleted:      lead.IsDeleted,
	}
}

func (s *Server) CreateLead(ctx context.Context, req *pb.CreateLeadRequest) (*pb.CreateLeadResponse, error) {
	lead, err := s.App.CreateLead(ctx, req.AdId, req.ClientCompany, req.ClientEmployee)
	if err != nil {
		return nil, mapErrors(err)
	}
	return &pb.CreateLeadResponse{
		Lead: modelLeadToLead(lead),
	}, nil
}

func (s *Server) GetLeads(ctx context.Context, req *pb.GetLeadsRequest) (*pb.GetLeadsResponse, error) {
	leads, amount, err := s.App.GetLeads(ctx, req.CompanyId, req.EmployeeId, model.Filter{
		Limit:         uint(req.Filter.Limit),
		Offset:        uint(req.Filter.Offset),
		Status:        req.Filter.Status,
		ByStatus:      req.Filter.ByStatus,
		Responsible:   req.Filter.Responsible,
		ByResponsible: req.Filter.ByResponsible,
	})
	if err != nil {
		return nil, mapErrors(err)
	}
	resp := &pb.GetLeadsResponse{
		Leads:  make([]*pb.Lead, len(leads)),
		Amount: uint64(amount),
	}
	for i, lead := range leads {
		resp.Leads[i] = modelLeadToLead(lead)
	}
	return resp, nil
}

func (s *Server) GetLeadById(ctx context.Context, req *pb.GetLeadByIdRequest) (*pb.GetLeadByIdResponse, error) {
	lead, err := s.App.GetLeadById(ctx, req.CompanyId, req.EmployeeId, req.LeadId)
	if err != nil {
		return nil, mapErrors(err)
	}
	return &pb.GetLeadByIdResponse{
		Lead: modelLeadToLead(lead),
	}, nil
}

func (s *Server) UpdateLead(ctx context.Context, req *pb.UpdateLeadRequest) (*pb.UpdateLeadResponse, error) {
	lead, err := s.App.UpdateLead(ctx, req.CompanyId, req.EmployeeId, req.Id, model.UpdateLead{
		Title:       req.Upd.Title,
		Description: req.Upd.Description,
		Price:       uint(req.Upd.Price),
		Status:      req.Upd.Status,
		Responsible: req.Upd.Responsible,
	})
	if err != nil {
		return nil, mapErrors(err)
	}
	return &pb.UpdateLeadResponse{
		Lead: modelLeadToLead(lead),
	}, nil
}

func (s *Server) GetStatuses(ctx context.Context, _ *emptypb.Empty) (*pb.GetStatusesResponse, error) {
	if statuses, err := s.App.GetStatuses(ctx); err != nil {
		return nil, mapErrors(err)
	} else {
		return &pb.GetStatusesResponse{Data: statuses}, nil
	}
}
