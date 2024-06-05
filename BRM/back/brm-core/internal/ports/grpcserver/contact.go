package grpcserver

import (
	"brm-core/internal/model"
	"brm-core/internal/ports/grpcserver/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"time"
)

func contactToModelContact(contact *pb.Contact) model.Contact {
	if contact == nil {
		return model.Contact{}
	}
	return model.Contact{
		Id:           contact.Id,
		OwnerId:      contact.OwnerId,
		EmployeeId:   contact.EmployeeId,
		Notes:        contact.Notes,
		CreationDate: time.Unix(contact.CreationDate, 0),
		IsDeleted:    contact.IsDeleted,
		Empl:         employeeToModelEmployee(contact.Empl),
	}
}

func modelContactToContact(contact model.Contact) *pb.Contact {
	if contact.Id == 0 {
		return nil
	}
	return &pb.Contact{
		Id:           contact.Id,
		OwnerId:      contact.OwnerId,
		EmployeeId:   contact.EmployeeId,
		Notes:        contact.Notes,
		CreationDate: contact.CreationDate.UTC().Unix(),
		IsDeleted:    contact.IsDeleted,
		Empl:         modelEmployeeToEmployee(contact.Empl),
	}
}

func (s *Server) CreateContact(ctx context.Context, req *pb.CreateContactRequest) (*pb.CreateContactResponse, error) {
	contact, err := s.App.CreateContact(ctx,
		req.OwnerId,
		req.EmployeeId,
	)
	if err != nil {
		return nil, mapErrors(err)
	}
	return &pb.CreateContactResponse{
		Contact: modelContactToContact(contact),
	}, nil
}

func (s *Server) UpdateContact(ctx context.Context, req *pb.UpdateContactRequest) (*pb.UpdateContactResponse, error) {
	contact, err := s.App.UpdateContact(ctx,
		req.OwnerId,
		req.ContactId,
		model.UpdateContact{
			Notes: req.Upd.Notes,
		},
	)
	if err != nil {
		return nil, mapErrors(err)
	}
	return &pb.UpdateContactResponse{
		Contact: modelContactToContact(contact),
	}, nil
}

func (s *Server) DeleteContact(ctx context.Context, req *pb.DeleteContactRequest) (*empty.Empty, error) {
	if err := s.App.DeleteContact(ctx,
		req.OwnerId,
		req.ContactId,
	); err != nil {
		return nil, mapErrors(err)
	}
	return &empty.Empty{}, nil
}

func (s *Server) GetContacts(ctx context.Context, req *pb.GetContactsRequest) (*pb.GetContactsResponse, error) {
	contacts, amount, err := s.App.GetContacts(
		ctx,
		req.OwnerId,
		model.GetContacts{
			Limit:  uint(req.Pagination.Limit),
			Offset: uint(req.Pagination.Offset),
		})
	if err != nil {
		return nil, mapErrors(err)
	}

	resp := &pb.GetContactsResponse{
		List:   make([]*pb.Contact, len(contacts)),
		Amount: uint64(amount),
	}

	for i, contact := range contacts {
		resp.List[i] = modelContactToContact(contact)
	}
	return resp, nil
}

func (s *Server) GetContactById(ctx context.Context, req *pb.GetContactByIdRequest) (*pb.GetContactByIdResponse, error) {
	contact, err := s.App.GetContactById(ctx, req.OwnerId, req.ContactId)
	if err != nil {
		return nil, mapErrors(err)
	}

	return &pb.GetContactByIdResponse{
		Contact: modelContactToContact(contact),
	}, nil
}
