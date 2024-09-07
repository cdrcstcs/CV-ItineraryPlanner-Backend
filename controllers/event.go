package controllers

import (
	"context"

	"itineraryplanner/controllers/inf"
	"itineraryplanner/models"
	service_inf "itineraryplanner/service/inf"
)

func NewEventController(ser service_inf.EventService) inf.EventController {
	return &EventController{
		ser: ser,
	}
}

type EventController struct {
	ser service_inf.EventService
}

func (e *EventController) CreateEvent(ctx context.Context, req *models.CreateEventReq) (*models.CreateEventResp, error) {
	// we usually do request checking here at this layer, or can even do permission checking.
	// For now, I will leave it empty and just call service
	return e.ser.CreateEvent(ctx, req)
}

func (e *EventController) GetEvent(ctx context.Context, req *models.GetEventReq) (*models.GetEventResp, error) {
	return e.ser.GetEvent(ctx, req)
}
func (e *EventController) GetEventById(ctx context.Context, req *models.GetEventByIdReq) (*models.GetEventByIdResp, error) {
	return e.ser.GetEventById(ctx, req)
}

func (e *EventController) UpdateEvent(ctx context.Context, req *models.UpdateEventReq) (*models.UpdateEventResp, error) {
	// we usually do request checking here at this layer, or can even do permission checking.
	// For now, I will leave it empty and just call service
	return e.ser.UpdateEvent(ctx, req)
}

func (e *EventController) DeleteEvent(ctx context.Context, req *models.DeleteEventReq) (*models.DeleteEventResp, error) {
	return e.ser.DeleteEvent(ctx, req)
}

