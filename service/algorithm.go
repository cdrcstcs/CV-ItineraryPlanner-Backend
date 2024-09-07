package service

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"itineraryplanner/common/custom_errs"
	dal_inf "itineraryplanner/dal/inf"
	"itineraryplanner/models"
	"itineraryplanner/service/inf"

	"itineraryplanner/constant"
	"itineraryplanner/dal"
	"itineraryplanner/dal/db"
)


func NewRecommendItineraryService(rdal dal_inf.RecommendItineraryDal) inf.RecommendItineraryService {
	return &AlgoService{
		RDal: rdal,
	}
}
func NewBuildItineraryService(bdal dal_inf.BuildItineraryDal) inf.BuildItineraryService {
	return &AlgoService{
		BDal: bdal,
	}
}

func NewAlgoDTOService() inf.AlgoItineraryDTOService {
	return &AlgoService{}
}
type AlgoService struct {
	RDal dal_inf.RecommendItineraryDal
	BDal dal_inf.BuildItineraryDal
}

func (a *AlgoService) RecommendItinerary(ctx context.Context, req *models.RecommendItineraryReq) (*models.RecommendItineraryResp, error) {
	coordinate := &models.Coordinate{}
	err := copier.Copy(coordinate, req)
	if err != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}

	itineraries, err := a.RDal.RecommendItinerary(ctx, coordinate)
	if err != nil {
		// TODO logging
		return nil, err
	}
	dtos := []*models.ItineraryDTO{}
	for _,v :=range itineraries {
		dto, err := a.AlgoConvertDBOToDTOItinerary(ctx, v)
		if err != nil {
			return nil, err
		}
		dtos = append(dtos, dto)
	}

	return &models.RecommendItineraryResp{Itineraries: dtos}, nil
}

func (a *AlgoService) AlgoConvertDBOToDTOItinerary(ctx context.Context, iti *models.Itinerary) (*models.ItineraryDTO, error) {
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
			return nil, err
		}
	
		itinerary.Rating = resp3.GRRB.Rating
	}

	return itinerary, nil
}

func (a *AlgoService) BuildItinerary(ctx context.Context, req *models.BuildItineraryReq) (*models.BuildItineraryResp, error) {
	coordinate := &models.Coordinate{}
	err := copier.Copy(coordinate, req)
	if err != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}
	itinerary, err := a.BDal.BuildItinerary(ctx, coordinate)
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}

	itinerary1 := &models.Itinerary{}
	ok := copier.Copy(itinerary1, itinerary)
	if ok != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}
	dto, err := a.AlgoConvertDBOToDTOItinerary(ctx, itinerary1)
	if err != nil {
		return nil, err
	}
	return &models.BuildItineraryResp{Itinerary: dto}, nil
}