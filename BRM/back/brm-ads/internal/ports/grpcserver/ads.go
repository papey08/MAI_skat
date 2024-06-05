package grpcserver

import (
	"brm-ads/internal/model"
	"brm-ads/internal/ports/grpcserver/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

func adToModelAd(ad *pb.Ad) model.Ad {
	if ad == nil {
		return model.Ad{}
	}
	return model.Ad{
		Id:           ad.Id,
		CompanyId:    ad.CompanyId,
		Title:        ad.Title,
		Text:         ad.Text,
		Industry:     ad.Industry,
		Price:        uint(ad.Price),
		ImageURL:     ad.ImageUrl,
		CreationDate: time.Unix(ad.CreationDate, 0),
		CreatedBy:    ad.CreatedBy,
		Responsible:  ad.Responsible,
		IsDeleted:    ad.IsDeleted,
	}
}

func modelAdToAd(ad model.Ad) *pb.Ad {
	if ad.Id == 0 {
		return nil
	}
	return &pb.Ad{
		Id:           ad.Id,
		CompanyId:    ad.CompanyId,
		Title:        ad.Title,
		Text:         ad.Text,
		Industry:     ad.Industry,
		Price:        uint64(ad.Price),
		ImageUrl:     ad.ImageURL,
		CreationDate: ad.CreationDate.UTC().Unix(),
		CreatedBy:    ad.CreatedBy,
		Responsible:  ad.Responsible,
		IsDeleted:    ad.IsDeleted,
	}
}

func (s *Server) GetAdById(ctx context.Context, req *pb.GetAdByIdRequest) (*pb.GetAdByIdResponse, error) {
	ad, err := s.App.GetAdById(ctx, req.Id)
	if err != nil {
		return nil, mapErrors(err)
	}
	return &pb.GetAdByIdResponse{
		Ad: modelAdToAd(ad),
	}, nil
}

func (s *Server) GetAdsList(ctx context.Context, req *pb.GetAdsListRequest) (*pb.GetAdsListResponse, error) {
	var params model.AdsListParams
	if req.Params.Search != nil {
		params.Search = &model.AdSearcher{Pattern: req.Params.Search.Pattern}
	}
	if req.Params.Filter != nil {
		params.Filter = &model.AdFilter{
			ByCompany:  req.Params.Filter.ByCompany,
			CompanyId:  req.Params.Filter.CompanyId,
			ByIndustry: req.Params.Filter.ByIndustry,
			Industry:   req.Params.Filter.Industry,
		}
	}
	if req.Params.Sort != nil {
		params.Sort = &model.AdSorter{
			ByPriceAsc:  req.Params.Sort.ByPriceAsc,
			ByPriceDesc: req.Params.Sort.ByPriceDesc,
			ByDateAsc:   req.Params.Sort.ByDateAsc,
			ByDateDesc:  req.Params.Sort.ByDateDesc,
		}
	}
	params.Limit = uint(req.Params.Limit)
	params.Offset = uint(req.Params.Offset)

	ads, amount, err := s.App.GetAdsList(ctx, params)
	if err != nil {
		return nil, mapErrors(err)
	}

	resp := &pb.GetAdsListResponse{
		List:   make([]*pb.Ad, len(ads)),
		Amount: uint64(amount),
	}
	for i, ad := range ads {
		resp.List[i] = modelAdToAd(ad)
	}
	return resp, nil
}

func (s *Server) CreateAd(ctx context.Context, req *pb.CreateAdRequest) (*pb.CreateAdResponse, error) {
	ad, err := s.App.CreateAd(ctx, req.CompanyId, req.EmployeeId, adToModelAd(req.Ad))
	if err != nil {
		return nil, mapErrors(err)
	}

	return &pb.CreateAdResponse{
		Ad: modelAdToAd(ad),
	}, nil
}

func (s *Server) UpdateAd(ctx context.Context, req *pb.UpdateAdRequest) (*pb.UpdateAdResponse, error) {
	ad, err := s.App.UpdateAd(ctx,
		req.CompanyId,
		req.EmployeeId,
		req.AdId,
		model.UpdateAd{
			Title:       req.Upd.Title,
			Text:        req.Upd.Text,
			Industry:    req.Upd.Industry,
			Price:       uint(req.Upd.Price),
			ImageURL:    req.Upd.ImageUrl,
			Responsible: req.Upd.Responsible,
		})
	if err != nil {
		return nil, mapErrors(err)
	}

	return &pb.UpdateAdResponse{
		Ad: modelAdToAd(ad),
	}, nil
}

func (s *Server) DeleteAd(ctx context.Context, req *pb.DeleteAdRequest) (*empty.Empty, error) {
	if err := s.App.DeleteAd(ctx,
		req.CompanyId,
		req.EmployeeId,
		req.AdId,
	); err != nil {
		return nil, mapErrors(err)
	}
	return &empty.Empty{}, nil
}

func (s *Server) GetIndustries(ctx context.Context, _ *emptypb.Empty) (*pb.GetIndustriesResponse, error) {
	if industries, err := s.App.GetIndustries(ctx); err != nil {
		return nil, mapErrors(err)
	} else {
		return &pb.GetIndustriesResponse{Data: industries}, nil
	}
}
