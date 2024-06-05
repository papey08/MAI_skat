package grpcads

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
	"transport-api/internal/adapters/grpcads/pb"
	"transport-api/internal/model"
	"transport-api/internal/model/ads"
)

func respToAd(ad *pb.Ad) ads.Ad {
	if ad == nil {
		return ads.Ad{}
	}
	return ads.Ad{
		Id:           ad.Id,
		CompanyId:    ad.CompanyId,
		Title:        ad.Title,
		Text:         ad.Text,
		Industry:     ad.Industry,
		Price:        uint(ad.Price),
		ImageURL:     ad.ImageUrl,
		CreationDate: ad.CreationDate,
		CreatedBy:    ad.CreatedBy,
		Responsible:  ad.Responsible,
		IsDeleted:    ad.IsDeleted,
	}
}

func paramsToResp(params ads.ListParams) *pb.AdsListParams {
	res := &pb.AdsListParams{}
	res.Limit = uint64(params.Limit)
	res.Offset = uint64(params.Offset)
	if params.Search != nil {
		res.Search = &pb.AdSearcher{Pattern: params.Search.Pattern}
	}
	if params.Filter != nil {
		res.Filter = &pb.AdFilter{
			ByCompany:  params.Filter.ByCompany,
			CompanyId:  params.Filter.CompanyId,
			ByIndustry: params.Filter.ByIndustry,
			Industry:   params.Filter.Industry,
		}
	}
	if params.Sort != nil {
		res.Sort = &pb.AdSorter{
			ByPriceAsc:  params.Sort.ByPriceAsc,
			ByPriceDesc: params.Sort.ByPriceDesc,
			ByDateAsc:   params.Sort.ByDateAsc,
			ByDateDesc:  params.Sort.ByDateDesc,
		}
	}
	return res
}

func (a *adsClientImpl) GetAdById(ctx context.Context, id uint64) (ads.Ad, error) {
	resp, err := a.cli.GetAdById(ctx, &pb.GetAdByIdRequest{
		Id: id,
	})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.NotFound:
			return ads.Ad{}, model.ErrAdNotExists
		case codes.ResourceExhausted:
			return ads.Ad{}, model.ErrAdsError
		default:
			return ads.Ad{}, model.ErrAdsUnknown
		}
	}
	return respToAd(resp.Ad), nil
}

func (a *adsClientImpl) GetAdsList(ctx context.Context, params ads.ListParams) ([]ads.Ad, uint, error) {
	resp, err := a.cli.GetAdsList(ctx, &pb.GetAdsListRequest{
		Params: paramsToResp(params),
	})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.ResourceExhausted:
			return []ads.Ad{}, 0, model.ErrAdsError
		default:
			return []ads.Ad{}, 0, model.ErrAdsUnknown
		}
	}
	adsList := make([]ads.Ad, len(resp.List))
	for i, ad := range resp.List {
		adsList[i] = respToAd(ad)
	}
	return adsList, uint(resp.Amount), nil
}

func (a *adsClientImpl) CreateAd(ctx context.Context, companyId uint64, employeeId uint64, ad ads.Ad) (ads.Ad, error) {
	resp, err := a.cli.CreateAd(ctx, &pb.CreateAdRequest{
		CompanyId:  companyId,
		EmployeeId: employeeId,
		Ad: &pb.Ad{
			Id:           ad.Id,
			CompanyId:    ad.CompanyId,
			Title:        ad.Title,
			Text:         ad.Text,
			Industry:     ad.Industry,
			Price:        uint64(ad.Price),
			ImageUrl:     ad.ImageURL,
			CreationDate: ad.CreationDate,
			CreatedBy:    ad.CreatedBy,
			Responsible:  ad.Responsible,
			IsDeleted:    ad.IsDeleted,
		},
	})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.NotFound:
			return ads.Ad{}, model.ErrIndustryNotExists
		case codes.FailedPrecondition:
			return ads.Ad{}, model.ErrInvalidInput
		case codes.ResourceExhausted:
			return ads.Ad{}, model.ErrAdsError
		default:
			return ads.Ad{}, model.ErrAdsUnknown
		}
	}
	return respToAd(resp.Ad), nil
}

func (a *adsClientImpl) UpdateAd(ctx context.Context, companyId uint64, employeeId uint64, adId uint64, upd ads.UpdateAd) (ads.Ad, error) {
	resp, err := a.cli.UpdateAd(ctx, &pb.UpdateAdRequest{
		CompanyId:  companyId,
		EmployeeId: employeeId,
		AdId:       adId,
		Upd: &pb.UpdateAdFields{
			Title:       upd.Title,
			Text:        upd.Text,
			Industry:    upd.Industry,
			Price:       uint64(upd.Price),
			ImageUrl:    upd.ImageURL,
			Responsible: upd.Responsible,
		},
	})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.PermissionDenied:
			return ads.Ad{}, model.ErrPermissionDenied
		case codes.NotFound:
			// костыль, ну а чё ещё поделать
			if strings.Contains(err.Error(), "industry") {
				return ads.Ad{}, model.ErrIndustryNotExists
			} else {
				return ads.Ad{}, model.ErrAdNotExists
			}
		case codes.FailedPrecondition:
			return ads.Ad{}, model.ErrInvalidInput
		case codes.ResourceExhausted:
			return ads.Ad{}, model.ErrAdsError
		case codes.Unknown:
			return ads.Ad{}, model.ErrAdsUnknown
		}
	}
	return respToAd(resp.Ad), nil
}

func (a *adsClientImpl) DeleteAd(ctx context.Context, companyId uint64, employeeId uint64, adId uint64) error {
	_, err := a.cli.DeleteAd(ctx, &pb.DeleteAdRequest{
		CompanyId:  companyId,
		EmployeeId: employeeId,
		AdId:       adId,
	})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.PermissionDenied:
			return model.ErrPermissionDenied
		case codes.NotFound:
			return model.ErrAdNotExists
		case codes.ResourceExhausted:
			return model.ErrAdsError
		case codes.Unknown:
			return model.ErrAdsUnknown
		}
	}
	return nil
}

func (a *adsClientImpl) GetIndustries(ctx context.Context) (map[string]uint64, error) {
	resp, err := a.cli.GetIndustries(ctx, &empty.Empty{})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.ResourceExhausted:
			return map[string]uint64{}, model.ErrCoreError
		default:
			return map[string]uint64{}, model.ErrCoreUnknown
		}
	}
	return resp.Data, nil
}
