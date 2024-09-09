package controllers
import (
	"context"
	"errors"
	"itineraryplanner/controllers/inf"
	"itineraryplanner/models"
	service_inf "itineraryplanner/service/inf"
	"github.com/go-playground/validator/v10"
)
type ItineraryController struct {
	Ser       service_inf.ItineraryService
	Validator *validator.Validate
}
func NewItineraryController(ser service_inf.ItineraryService) inf.ItineraryController {
	return &ItineraryController{
		Ser:       ser,
		Validator: validator.New(),
	}
}
func (i *ItineraryController) validateRequest(req interface{}) error {
	err := i.Validator.Struct(req)
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
func (i *ItineraryController) CreateItinerary(ctx context.Context, req *models.CreateItineraryReq) (*models.CreateItineraryResp, error) {
	if err := i.validateRequest(req); err != nil {
		return nil, err
	}
	return i.Ser.CreateItinerary(ctx, req)
}
func (i *ItineraryController) GetItinerary(ctx context.Context, req *models.GetItineraryReq) (*models.GetItineraryResp, error) {
	return i.Ser.GetItinerary(ctx, req)
}
func (i *ItineraryController) GetItineraryById(ctx context.Context, req *models.GetItineraryByIdReq) (*models.GetItineraryByIdResp, error) {
	if err := i.validateRequest(req); err != nil {
		return nil, err
	}
	return i.Ser.GetItineraryById(ctx, req)
}
func (i *ItineraryController) UpdateItinerary(ctx context.Context, req *models.UpdateItineraryReq) (*models.UpdateItineraryResp, error) {
	if err := i.validateRequest(req); err != nil {
		return nil, err
	}
	return i.Ser.UpdateItinerary(ctx, req)
}
func (i *ItineraryController) DeleteItinerary(ctx context.Context, req *models.DeleteItineraryReq) (*models.DeleteItineraryResp, error) {
	if err := i.validateRequest(req); err != nil {
		return nil, err
	}
	return i.Ser.DeleteItinerary(ctx, req)
}