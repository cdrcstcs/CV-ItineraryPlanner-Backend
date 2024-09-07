package service

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"itineraryplanner/common/custom_errs"
	"itineraryplanner/constant"
	"itineraryplanner/dal"
	"itineraryplanner/dal/db"
	dal_inf "itineraryplanner/dal/inf"
	"itineraryplanner/models"
	"itineraryplanner/service/inf"
)

func NewCreateAttractionService(cdal dal_inf.CreateAttractionDal) inf.CreateAttractionService {
	return &AttractionService{
		CDal: cdal,
	}
}
func NewGetAttractionByIdService(bdal dal_inf.GetAttractionByIdDal) inf.GetAttractionByIdService {
	return &AttractionService{
		BDal: bdal,
	}
}
func NewGetAttractionService(gdal dal_inf.GetAttractionDal) inf.GetAttractionService {
	return &AttractionService{
		GDal: gdal,
	}
}
func NewUpdateAttractionService(udal dal_inf.UpdateAttractionDal) inf.UpdateAttractionService {
	return &AttractionService{
		UDal: udal,
	}
}
func NewDeleteAttractionService(ddal dal_inf.DeleteAttractionDal) inf.DeleteAttractionService {
	return &AttractionService{
		DDal: ddal,
	}
}
func NewAttractionDTOService() inf.AttractionDTOService {
	return &AttractionService{}
}

type AttractionService struct {
	CDal dal_inf.CreateAttractionDal
	BDal dal_inf.GetAttractionByIdDal
	GDal dal_inf.GetAttractionDal
	UDal dal_inf.UpdateAttractionDal
	DDal dal_inf.DeleteAttractionDal
}

func (a *AttractionService) CreateAttraction(ctx context.Context, req *models.CreateAttractionReq) (*models.CreateAttractionResp, error) {
	attraction := &models.Attraction{}
	err := copier.Copy(attraction, req)
	if err != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}

	attraction, err = a.CDal.CreateAttraction(ctx, attraction)
	if err != nil {
		// TODO logging
		return nil, err
	}
	dto, err := a.ConvertDBOToDTOAttraction(ctx, attraction)
	if err != nil {
		// TODO logging
		return nil, err
	}

	return &models.CreateAttractionResp{Attraction: dto}, nil
}

func (a *AttractionService) ConvertDBOToDTOAttraction(ctx context.Context, att *models.Attraction) (*models.AttractionDTO, error) {
	log.Info().Msg("ConvertDBOToDTOAttraction is called") // Add this line for debugging
	if att == nil {
		return nil, custom_errs.ServerError
	}
	facade := &Facade{
		Service: &FacadeField{
			TS: &TagService{
				GDal: &dal.TagDal{
					MainDB: db.GetMemoMongo(constant.MainMongoDB),
				},
			},
			CS: &CoordinateService{
				GDal: &dal.CoordinateDal{
					MainDB: db.GetMemoMongo(constant.MainMongoDB),
				},
			},
			RS: &RatingService{
				GDal: &dal.RatingDal{
					MainDB: db.GetMemoMongo(constant.MainMongoDB),
				},
			},
		},
	}

	attraction := &models.AttractionDTO{}
	err:= copier.Copy(attraction, att)
	if err != nil {
		// TODO logging
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}

	tagsDTO := []*models.TagDTO{}
	if att.TagIDs != nil {
		for _, v:= range att.TagIDs {
			req1 := &models.GetTagByIdReq{
				Id: v,
			}
			resp1, err := facade.Execute(ctx, &models.ReqFacade{
				GTRB: req1,
			}, "GTSB")
			if err != nil {
				// TODO logging
				return nil, err
			}
			tagsDTO = append(tagsDTO, resp1.GTRB.Tag)
		}
		tagMap := map[string]bool{}
		for _, v := range tagsDTO {
			tagMap[v.Id] = true
		}
		for _, v := range att.TagIDs {
			if !tagMap[v] {
				return nil, errors.New("invalid tag id")
			}
		}
		attraction.Tags = tagsDTO
	}

	if att.CoordinateId != ""{
		req2 := &models.GetCoordinateByIdReq{
			Id: att.CoordinateId,
		}
		resp2, err := facade.Execute(ctx, &models.ReqFacade{
			GCRB: req2,
		}, "GCSB")
		if err != nil {
			// TODO logging
			return nil, err
		}
		attraction.Coordinate = resp2.GCRB.Coordinate
	}
	if att.RatingId != ""{
		req3 := &models.GetRatingByIdReq{
			Id: att.RatingId,
		}
		resp3, err := facade.Execute(ctx, &models.ReqFacade{
			GRRB: req3,
		}, "GRSB")
		if err != nil {
			// TODO logging
			return nil, err
		}
		attraction.Rating = resp3.GRRB.Rating
	}
	return attraction, nil
}

func (a *AttractionService) GetAttractionById(ctx context.Context, req *models.GetAttractionByIdReq) (*models.GetAttractionByIdResp, error) {
	attraction, err := a.BDal.GetAttractionById(ctx, req.Id)
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}

	attraction1 := &models.Attraction{}
	ok := copier.Copy(attraction1, attraction)
	if ok != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}
	dto, err := a.ConvertDBOToDTOAttraction(ctx, attraction1)
	if err != nil {
		// TODO logging
		return nil, err
	}
	return &models.GetAttractionByIdResp{Attraction: dto}, nil
}

func (a *AttractionService) GetAttraction(ctx context.Context, req *models.GetAttractionReq) (*models.GetAttractionResp, error) {
	attractions, err := a.GDal.GetAttraction(ctx)
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}

	attraction1 := []models.Attraction{}
	ok := copier.Copy(attraction1, attractions)
	if ok != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}
	dtos := make([]*models.AttractionDTO, 0)
	for _, v := range attraction1 {
		dto, err := a.ConvertDBOToDTOAttraction(ctx, &v)
		if err != nil {
			// TODO logging
			return nil, err
		}
		dtos = append(dtos, dto)
	}
	return &models.GetAttractionResp{Attractions: dtos}, nil
}

func (a *AttractionService) UpdateAttraction(ctx context.Context, req *models.UpdateAttractionReq) (*models.UpdateAttractionResp, error) {
	attraction := &models.Attraction{}
	err := copier.Copy(attraction, req)
	if err != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}

	attraction, err = a.UDal.UpdateAttraction(ctx, attraction)
	if err != nil {
		// TODO logging
		return nil, err
	}
	dto, err := a.ConvertDBOToDTOAttraction(ctx, attraction)
	if err != nil {
		// TODO logging
		return nil, err
	}

	return &models.UpdateAttractionResp{Attraction: dto}, nil
}

func (a *AttractionService) DeleteAttraction(ctx context.Context, req *models.DeleteAttractionReq) (*models.DeleteAttractionResp, error) {
	attraction, err := a.DDal.DeleteAttraction(ctx, req.Id)
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}

	attraction1 := &models.Attraction{}
	ok := copier.Copy(attraction1, attraction)
	if ok != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}
	dto, err := a.ConvertDBOToDTOAttraction(ctx, attraction1)
	if err != nil {
		// TODO logging
		return nil, err
	}
	return &models.DeleteAttractionResp{Attraction: dto}, nil
}
