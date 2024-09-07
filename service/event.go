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

func NewCreateEventService(cdal dal_inf.CreateEventDal) inf.CreateEventService {
	return &EventService{
		CDal: cdal,
	}
}
func NewGetEventByIdService(bdal dal_inf.GetEventByIdDal) inf.GetEventByIdService {
	return &EventService{
		BDal: bdal,
	}
}
func NewGetEventService(gdal dal_inf.GetEventDal) inf.GetEventService {
	return &EventService{
		GDal: gdal,
	}
}
func NewUpdateEventService(udal dal_inf.UpdateEventDal) inf.UpdateEventService {
	return &EventService{
		UDal: udal,
	}
}
func NewDeleteEventService(ddal dal_inf.DeleteEventDal) inf.DeleteEventService {
	return &EventService{
		DDal: ddal,
	}
}
func NewEventDTOService() inf.EventDTOService {
	return &EventService{}
}


type EventService struct {
	CDal dal_inf.CreateEventDal
	BDal dal_inf.GetEventByIdDal
	GDal dal_inf.GetEventDal
	UDal dal_inf.UpdateEventDal
	DDal dal_inf.DeleteEventDal
}

func (e *EventService) CreateEvent(ctx context.Context, req *models.CreateEventReq) (*models.CreateEventResp, error) {
	event := &models.Event{}
	err := copier.Copy(event, req)
	if err != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}

	event, err = e.CDal.CreateEvent(ctx, event)
	if err != nil {
		// TODO logging
		return nil, err
	}

	dto, err := e.ConvertDBOToDTOEvent(ctx, event)
	if err != nil {
		// TODO logging
		return nil, err
	}

	return &models.CreateEventResp{Event: dto}, nil
}

func (e *EventService) ConvertDBOToDTOEvent(ctx context.Context, eve *models.Event) (*models.EventDTO, error) {
	if eve == nil {
		return nil, custom_errs.ServerError
	}
	facade := &Facade{
		Service: &FacadeField{
			AS: &AttractionService{
				GDal: &dal.AttractionDal{
					MainDB: db.GetMemoMongo(constant.MainMongoDB),
				},
			},
		},
	}
	event := &models.EventDTO{}
	err:= copier.Copy(event, eve)
	if err != nil {
		// TODO logging
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}
	if eve.AttractionId != ""{
		req1 := &models.GetAttractionByIdReq{
			Id: eve.AttractionId,
		}
		resp1, err := facade.Execute(ctx, &models.ReqFacade{
			GARB: req1,
		}, "GASB")
		if err != nil {
			// TODO logging
			return nil, err
		}
	
		event.Attraction = resp1.GARB.Attraction
	}
	return event, nil
}

func (e *EventService) GetEventById(ctx context.Context, req *models.GetEventByIdReq) (*models.GetEventByIdResp, error) {
	Event, err := e.BDal.GetEventById(ctx, req.Id)
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}

	Event1 := &models.Event{}
	ok := copier.Copy(Event1, Event)
	if ok != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}
	dto, err := e.ConvertDBOToDTOEvent(ctx, Event1)
	if err != nil {
		// TODO logging
		return nil, err
	}
	return &models.GetEventByIdResp{Event: dto}, nil
}

func (e *EventService) GetEvent(ctx context.Context, req *models.GetEventReq) (*models.GetEventResp, error) {
	Events, err := e.GDal.GetEvent(ctx)
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}

	Event1 := []models.Event{}
	ok := copier.Copy(Event1, Events)
	if ok != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}
	dtos := make([]*models.EventDTO, 0)
	for _, v := range Event1 {
		dto, err := e.ConvertDBOToDTOEvent(ctx, &v)
		if err != nil {
			// TODO logging
			return nil, err
		}
		dtos = append(dtos, dto)
	}
	return &models.GetEventResp{Events: dtos}, nil
}

func (e *EventService) UpdateEvent(ctx context.Context, req *models.UpdateEventReq) (*models.UpdateEventResp, error) {
	event := &models.Event{}
	err := copier.Copy(event, req)
	if err != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}

	event, err = e.UDal.UpdateEvent(ctx, event)
	if err != nil {
		// TODO logging
		return nil, err
	}
	dto, err := e.ConvertDBOToDTOEvent(ctx, event)
	if err != nil {
		// TODO logging
		return nil, err
	}

	return &models.UpdateEventResp{Event: dto}, nil
}

func (e *EventService) DeleteEvent(ctx context.Context, req *models.DeleteEventReq) (*models.DeleteEventResp, error) {
	event, err := e.DDal.DeleteEvent(ctx, req.Id)
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}

	event1 := &models.Event{}
	ok := copier.Copy(event1, event)
	if ok != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}
	dto, err := e.ConvertDBOToDTOEvent(ctx, event1)
	if err != nil {
		// TODO logging
		return nil, err
	}
	return &models.DeleteEventResp{Event: dto}, nil
}