package controllers

import (
	"context"

	"itineraryplanner/controllers/inf"
	"itineraryplanner/models"
	service_inf "itineraryplanner/service/inf"
)

func NewItineraryController(ser service_inf.ItineraryService) inf.ItineraryController {
	return &ItineraryController{
		ser: ser,
	}
}
type ItineraryController struct {
	ser service_inf.ItineraryService
}

func (i *ItineraryController) CreateItinerary(ctx context.Context, req *models.CreateItineraryReq) (*models.CreateItineraryResp, error) {
	// we usually do request checking here at this layer, or can even do permission checking.
	// For now, I will leave it empty and just call service
	return i.ser.CreateItinerary(ctx, req)
}

func (i *ItineraryController) GetItinerary(ctx context.Context, req *models.GetItineraryReq) (*models.GetItineraryResp, error) {
	return i.ser.GetItinerary(ctx, req)
}
func (i *ItineraryController) GetItineraryById(ctx context.Context, req *models.GetItineraryByIdReq) (*models.GetItineraryByIdResp, error) {
	return i.ser.GetItineraryById(ctx, req)
}

func (i *ItineraryController) UpdateItinerary(ctx context.Context, req *models.UpdateItineraryReq) (*models.UpdateItineraryResp, error) {
	// we usually do request checking here at this layer, or can even do permission checking.
	// For now, I will leave it empty and just call service
	return i.ser.UpdateItinerary(ctx, req)
}

func (i *ItineraryController) DeleteItinerary(ctx context.Context, req *models.DeleteItineraryReq) (*models.DeleteItineraryResp, error) {
	return i.ser.DeleteItinerary(ctx, req)
}

