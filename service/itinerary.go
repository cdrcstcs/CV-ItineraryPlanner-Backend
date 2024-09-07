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

func NewCreateItineraryService(cdal dal_inf.CreateItineraryDal) inf.CreateItineraryService {
	return &ItineraryService{
		CDal: cdal,
	}
}
func NewGetItineraryByIdService(bdal dal_inf.GetItineraryByIdDal) inf.GetItineraryByIdService {
	return &ItineraryService{
		BDal: bdal,
	}
}
func NewGetItineraryService(gdal dal_inf.GetItineraryDal) inf.GetItineraryService {
	return &ItineraryService{
		GDal: gdal,
	}
}
func NewUpdateItineraryService(udal dal_inf.UpdateItineraryDal) inf.UpdateItineraryService {
	return &ItineraryService{
		UDal: udal,
	}
}
func NewDeleteItineraryService(ddal dal_inf.DeleteItineraryDal) inf.DeleteItineraryService {
	return &ItineraryService{
		DDal: ddal,
	}
}
func NewItineraryDTOService() inf.ItineraryDTOService {
	return &ItineraryService{}
}

type ItineraryService struct {
	CDal dal_inf.CreateItineraryDal
	BDal dal_inf.GetItineraryByIdDal
	GDal dal_inf.GetItineraryDal
	UDal dal_inf.UpdateItineraryDal
	DDal dal_inf.DeleteItineraryDal
}

func (i *ItineraryService) CreateItinerary(ctx context.Context, req *models.CreateItineraryReq) (*models.CreateItineraryResp, error) {
	itinerary := &models.Itinerary{}
	err := copier.Copy(itinerary, req)
	if err != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}

	itinerary, err = i.CDal.CreateItinerary(ctx, itinerary)
	if err != nil {
		// TODO logging
		return nil, err
	}


	dto, err := i.ConvertDBOToDTOItinerary(ctx, itinerary)
	if err != nil {
		// TODO logging
		return nil, err
	}

	return &models.CreateItineraryResp{Itinerary: dto}, nil
}

func (i *ItineraryService) ConvertDBOToDTOItinerary(ctx context.Context, iti *models.Itinerary) (*models.ItineraryDTO, error) {
	if iti == nil {
		return nil, custom_errs.ServerError
	}
	facade := &Facade{
		Service: &FacadeField{
			US: &UserService{
				GDal: &dal.UserDal{
					MainDB: db.GetMemoMongo(constant.MainMongoDB),
				},
			},
			ES: &EventService{
				GDal: &dal.EventDal{
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

	itinerary := &models.ItineraryDTO{}
	err := copier.Copy(itinerary, iti)
	if err != nil {
		// TODO logging
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}

	if iti.UserId != ""{
		req1 := &models.GetUserByIdReq{
			Id: iti.UserId,
		}
		resp1, err := facade.Execute(ctx, &models.ReqFacade{
			GURB: req1,
		}, "GUSB")
		if err != nil {
			// TODO logging
			return nil, err
		}
		itinerary.User = resp1.GURB.User
	}

	if iti.EventIds != nil {
		eventsDTO := []*models.EventDTO{}
	
		for _,v :=range iti.EventIds{
			req2 := &models.GetEventByIdReq{
				Id: v,
			}
			resp2, err := facade.Execute(ctx, &models.ReqFacade{
				GERB: req2,
			}, "GESB")
			if err != nil {
				// TODO logging
				return nil, err
			}
			eventsDTO = append(eventsDTO, resp2.GERB.Event)
		}
		itinerary.Events = eventsDTO
	}

	if iti.RatingId != ""{
		req3 := &models.GetRatingByIdReq{
			Id: iti.RatingId,
		}
		resp3, err := facade.Execute(ctx, &models.ReqFacade{
			GRRB: req3,
		}, "GRSB")
		if err != nil {
			// TODO logging
			return nil, err
		}
	
		itinerary.Rating = resp3.GRRB.Rating
	}

	return itinerary, nil
}

func (i *ItineraryService) GetItineraryById(ctx context.Context, req *models.GetItineraryByIdReq) (*models.GetItineraryByIdResp, error) {
	Itinerary, err := i.BDal.GetItineraryById(ctx, req.Id)
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}

	Itinerary1 := &models.Itinerary{}
	ok := copier.Copy(Itinerary1, Itinerary)
	if ok != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}
	dto, err := i.ConvertDBOToDTOItinerary(ctx, Itinerary1)
	if err != nil {
		// TODO logging
		return nil, err
	}
	return &models.GetItineraryByIdResp{Itinerary: dto}, nil
}

func (i *ItineraryService) GetItinerary(ctx context.Context, req *models.GetItineraryReq) (*models.GetItineraryResp, error) {
	Itineraries, err := i.GDal.GetItinerary(ctx)
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}

	Itinerary1 := []models.Itinerary{}
	ok := copier.Copy(Itinerary1, Itineraries)
	if ok != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}
	dtos := make([]*models.ItineraryDTO, 0)
	for _, v := range Itinerary1 {
		dto, err := i.ConvertDBOToDTOItinerary(ctx, &v)
		if err != nil {
			// TODO logging
			return nil, err
		}
		dtos = append(dtos, dto)
	}
	return &models.GetItineraryResp{Itineraries: dtos}, nil
}


func (i *ItineraryService) UpdateItinerary(ctx context.Context, req *models.UpdateItineraryReq) (*models.UpdateItineraryResp, error) {
	itinerary := &models.Itinerary{}
	err := copier.Copy(itinerary, req)
	if err != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}

	itinerary, err = i.UDal.UpdateItinerary(ctx, itinerary)
	if err != nil {
		// TODO logging
		return nil, err
	}
	dto, err := i.ConvertDBOToDTOItinerary(ctx, itinerary)
	if err != nil {
		// TODO logging
		return nil, err
	}

	return &models.UpdateItineraryResp{Itinerary: dto}, nil
}

func (i *ItineraryService) DeleteItinerary(ctx context.Context, req *models.DeleteItineraryReq) (*models.DeleteItineraryResp, error) {
	itinerary, err := i.DDal.DeleteItinerary(ctx, req.Id)
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}

	itinerary1 := &models.Itinerary{}
	ok := copier.Copy(itinerary1, itinerary)
	if ok != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}
	dto, err := i.ConvertDBOToDTOItinerary(ctx, itinerary1)
	if err != nil {
		// TODO logging
		return nil, err
	}
	return &models.DeleteItineraryResp{Itinerary: dto}, nil
}