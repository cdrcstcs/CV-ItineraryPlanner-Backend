package service
import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/rs/zerolog/log"
	"itineraryplanner/common/custom_errs"
	dal_inf "itineraryplanner/dal/inf"
	"itineraryplanner/models"
	"itineraryplanner/service/inf"
)
func NewEventService(dal dal_inf.EventDal) inf.EventService {
	return &EventService{
		Dal: dal,
	}
}
type EventService struct {
	Dal dal_inf.EventDal
}
func (e *EventService) CreateEvent(ctx context.Context, req *models.CreateEventReq) (*models.CreateEventResp, error) {
	event := &models.Event{}
	err := copier.Copy(event, req)
	if err != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, custom_errs.InvalidInput
	}
	event, err = e.Dal.CreateEvent(ctx, event)
	if err != nil {
		return nil, err
	}
	dto, err := e.ConvertDBOToDTOEvent(ctx, event)
	if err != nil {
		return nil, err
	}
	return &models.CreateEventResp{Event: dto}, nil
}
func (a *EventService) ConvertDBOToDTOEvent(ctx context.Context, att *models.Event) (*models.EventDTO, error) {
	log.Info().Msg("ConvertDBOToDTOEvent is called")
	if att == nil {
		return nil, custom_errs.ServerError
	}
	ret := &models.EventDTO{}
	err := copier.Copy(ret,att)
	if err != nil {
		return nil, custom_errs.InvalidInput
	}
	return ret, nil
}
func (e *EventService) GetEventById(ctx context.Context, req *models.GetEventByIdReq) (*models.GetEventByIdResp, error) {
	Event, err := e.Dal.GetEventById(ctx, req.Id)
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}
	Event1 := &models.Event{}
	err = copier.Copy(Event1, Event)
	if err != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, custom_errs.InvalidInput
	}
	dto, err := e.ConvertDBOToDTOEvent(ctx, Event1)
	if err != nil {
		return nil, err
	}
	return &models.GetEventByIdResp{Event: dto}, nil
}
func (e *EventService) GetEvent(ctx context.Context, req *models.GetEventReq) (*models.GetEventResp, error) {
	Events, err := e.Dal.GetEvent(ctx)
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}
	Event1 := []models.Event{}
	err = copier.Copy(Event1, Events)
	if err != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, custom_errs.InvalidInput
	}
	dtos := make([]*models.EventDTO, 0)
	for _, v := range Event1 {
		dto, err := e.ConvertDBOToDTOEvent(ctx, &v)
		if err != nil {
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
		return nil, custom_errs.InvalidInput
	}
	event, err = e.Dal.UpdateEvent(ctx, event)
	if err != nil {
		return nil, err
	}
	dto, err := e.ConvertDBOToDTOEvent(ctx, event)
	if err != nil {
		return nil, err
	}
	return &models.UpdateEventResp{Event: dto}, nil
}
func (e *EventService) DeleteEvent(ctx context.Context, req *models.DeleteEventReq) (*models.DeleteEventResp, error) {
	event, err := e.Dal.DeleteEvent(ctx, req.Id)
	if err != nil {
		return nil, custom_errs.DBErrDeleteWithID
	}
	event1 := &models.Event{}
	ok := copier.Copy(event1, event)
	if ok != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", ok.Error())
		return nil, custom_errs.InvalidInput
	}
	dto, err := e.ConvertDBOToDTOEvent(ctx, event1)
	if err != nil {
		return nil, err
	}
	return &models.DeleteEventResp{Event: dto}, nil
}