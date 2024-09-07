package controllers

import (
	"context"

	"itineraryplanner/controllers/inf"
	"itineraryplanner/models"
	service_inf "itineraryplanner/service/inf"
)

func NewCreateEventController(serC service_inf.CreateEventService) inf.CreateEventController {
	return &EventController{
		serC: serC,
	}
}
func NewGetEventController(serG service_inf.GetEventService) inf.GetEventController {
	return &EventController{
		serG: serG,
	}
}
func NewGetEventByIdController(serB service_inf.GetEventByIdService) inf.GetEventByIdController {
	return &EventController{
		serB: serB,
	}
}
func NewUpdateEventController(serU service_inf.UpdateEventService) inf.UpdateEventController {
	return &EventController{
		serU: serU,
	}
}
func NewDeleteEventController(serD service_inf.DeleteEventService) inf.DeleteEventController {
	return &EventController{
		serD: serD,
	}
}


type EventController struct {
	serC service_inf.CreateEventService
	serB service_inf.GetEventByIdService
	serG service_inf.GetEventService
	serU service_inf.UpdateEventService
	serD service_inf.DeleteEventService
}

func (e *EventController) CreateEvent(ctx context.Context, req *models.CreateEventReq) (*models.CreateEventResp, error) {
	// we usually do request checking here at this layer, or can even do permission checking.
	// For now, I will leave it empty and just call service
	return e.serC.CreateEvent(ctx, req)
}

func (e *EventController) GetEvent(ctx context.Context, req *models.GetEventReq) (*models.GetEventResp, error) {
	return e.serG.GetEvent(ctx, req)
}
func (e *EventController) GetEventById(ctx context.Context, req *models.GetEventByIdReq) (*models.GetEventByIdResp, error) {
	return e.serB.GetEventById(ctx, req)
}

func (e *EventController) UpdateEvent(ctx context.Context, req *models.UpdateEventReq) (*models.UpdateEventResp, error) {
	// we usually do request checking here at this layer, or can even do permission checking.
	// For now, I will leave it empty and just call service
	return e.serU.UpdateEvent(ctx, req)
}

func (e *EventController) DeleteEvent(ctx context.Context, req *models.DeleteEventReq) (*models.DeleteEventResp, error) {
	return e.serD.DeleteEvent(ctx, req)
}

