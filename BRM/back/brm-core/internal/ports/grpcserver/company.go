package grpcserver

import (
	"brm-core/internal/model"
	"brm-core/internal/ports/grpcserver/pb"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

func companyToModelCompany(company *pb.Company) model.Company {
	if company == nil {
		return model.Company{}
	}
	return model.Company{
		Id:           company.Id,
		Name:         company.Name,
		Description:  company.Description,
		Industry:     company.Industry,
		OwnerId:      company.OwnerId,
		Rating:       company.Rating,
		CreationDate: time.Unix(company.CreationDate, 0),
		IsDeleted:    company.IsDeleted,
	}
}

func modelCompanyToCompany(company model.Company) *pb.Company {
	if company.Id == 0 {
		return nil
	}
	return &pb.Company{
		Id:           company.Id,
		Name:         company.Name,
		Description:  company.Description,
		Industry:     company.Industry,
		OwnerId:      company.OwnerId,
		Rating:       company.Rating,
		CreationDate: company.CreationDate.UTC().Unix(),
		IsDeleted:    company.IsDeleted,
	}
}

func (s *Server) GetCompany(ctx context.Context, req *pb.GetCompanyRequest) (*pb.GetCompanyResponse, error) {
	company, err := s.App.GetCompany(ctx, req.Id)
	if err != nil {
		return nil, mapErrors(err)
	}
	return &pb.GetCompanyResponse{
		Company: modelCompanyToCompany(company),
	}, nil
}

func (s *Server) CreateCompanyAndOwner(ctx context.Context, req *pb.CreateCompanyAndOwnerRequest) (*pb.CreateCompanyAndOwnerResponse, error) {
	company, owner, err := s.App.CreateCompanyAndOwner(ctx,
		companyToModelCompany(req.Company),
		employeeToModelEmployee(req.Owner),
	)
	if err != nil {
		return nil, mapErrors(err)
	}
	return &pb.CreateCompanyAndOwnerResponse{
		Company: modelCompanyToCompany(company),
		Owner:   modelEmployeeToEmployee(owner),
	}, nil
}

func (s *Server) UpdateCompany(ctx context.Context, req *pb.UpdateCompanyRequest) (*pb.UpdateCompanyResponse, error) {
	company, err := s.App.UpdateCompany(ctx,
		req.CompanyId,
		req.OwnerId,
		model.UpdateCompany{
			Name:        req.Upd.Name,
			Description: req.Upd.Description,
			Industry:    req.Upd.Industry,
			OwnerId:     req.Upd.OwnerId,
		},
	)
	if err != nil {
		return nil, mapErrors(err)
	}
	return &pb.UpdateCompanyResponse{
		Company: modelCompanyToCompany(company),
	}, nil
}

func (s *Server) GetIndustries(ctx context.Context, _ *emptypb.Empty) (*pb.GetIndustriesResponse, error) {
	if industries, err := s.App.GetIndustries(ctx); err != nil {
		return nil, mapErrors(err)
	} else {
		return &pb.GetIndustriesResponse{Data: industries}, nil
	}
}
