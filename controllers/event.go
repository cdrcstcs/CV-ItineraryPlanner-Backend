package controllers
import (
	"context"
	"errors"
	"itineraryplanner/controllers/inf"
	"itineraryplanner/models"
	service_inf "itineraryplanner/service/inf"
	"github.com/go-playground/validator/v10"
)
type EventController struct {
	Ser       service_inf.EventService
	Validator *validator.Validate
}
func NewEventController(ser service_inf.EventService) inf.EventController {
	return &EventController{
		Ser:       ser,
		Validator: validator.New(),
	}
}
func (e *EventController) validateRequest(req interface{}) error {
	err := e.Validator.Struct(req)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}
		var errMsg string
		for _, err := range err.(validator.ValidationErrors) {
			errMsg += err.Field() + " " + err.Tag() + ", "
		}
		if len(errMsg) > 0 {
			return errors.New(errMsg)
		}
	}
	return nil
}
func (e *EventController) CreateEvent(ctx context.Context, req *models.CreateEventReq) (*models.CreateEventResp, error) {
	if err := e.validateRequest(req); err != nil {
		return nil, err
	}
	return e.Ser.CreateEvent(ctx, req)
}
func (e *EventController) GetEvent(ctx context.Context, req *models.GetEventReq) (*models.GetEventResp, error) {
	return e.Ser.GetEvent(ctx, req)
}
func (e *EventController) GetEventById(ctx context.Context, req *models.GetEventByIdReq) (*models.GetEventByIdResp, error) {
	if err := e.validateRequest(req); err != nil {
		return nil, err
	}
	return e.Ser.GetEventById(ctx, req)
}
func (e *EventController) UpdateEvent(ctx context.Context, req *models.UpdateEventReq) (*models.UpdateEventResp, error) {
	if err := e.validateRequest(req); err != nil {
		return nil, err
	}
	return e.Ser.UpdateEvent(ctx, req)
}
func (e *EventController) DeleteEvent(ctx context.Context, req *models.DeleteEventReq) (*models.DeleteEventResp, error) {
	if err := e.validateRequest(req); err != nil {
		return nil, err
	}
	return e.Ser.DeleteEvent(ctx, req)
}