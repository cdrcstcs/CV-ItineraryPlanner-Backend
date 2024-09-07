package controllers

import (
	"context"

	"itineraryplanner/controllers/inf"
	"itineraryplanner/models"
	service_inf "itineraryplanner/service/inf"
)

func NewCreateItineraryController(serC service_inf.CreateItineraryService) inf.CreateItineraryController {
	return &ItineraryController{
		serC: serC,
	}
}
func NewGetItineraryController(serG service_inf.GetItineraryService) inf.GetItineraryController {
	return &ItineraryController{
		serG: serG,
	}
}
func NewGetItineraryByIdController(serB service_inf.GetItineraryByIdService) inf.GetItineraryByIdController {
	return &ItineraryController{
		serB: serB,
	}
}
func NewUpdateItineraryController(serU service_inf.UpdateItineraryService) inf.UpdateItineraryController {
	return &ItineraryController{
		serU: serU,
	}
}
func NewDeleteItineraryController(serD service_inf.DeleteItineraryService) inf.DeleteItineraryController {
	return &ItineraryController{
		serD: serD,
	}
}


type ItineraryController struct {
	serC service_inf.CreateItineraryService
	serB service_inf.GetItineraryByIdService
	serG service_inf.GetItineraryService
	serU service_inf.UpdateItineraryService
	serD service_inf.DeleteItineraryService
}

func (i *ItineraryController) CreateItinerary(ctx context.Context, req *models.CreateItineraryReq) (*models.CreateItineraryResp, error) {
	// we usually do request checking here at this layer, or can even do permission checking.
	// For now, I will leave it empty and just call service
	return i.serC.CreateItinerary(ctx, req)
}

func (i *ItineraryController) GetItinerary(ctx context.Context, req *models.GetItineraryReq) (*models.GetItineraryResp, error) {
	return i.serG.GetItinerary(ctx, req)
}
func (i *ItineraryController) GetItineraryById(ctx context.Context, req *models.GetItineraryByIdReq) (*models.GetItineraryByIdResp, error) {
	return i.serB.GetItineraryById(ctx, req)
}

func (i *ItineraryController) UpdateItinerary(ctx context.Context, req *models.UpdateItineraryReq) (*models.UpdateItineraryResp, error) {
	// we usually do request checking here at this layer, or can even do permission checking.
	// For now, I will leave it empty and just call service
	return i.serU.UpdateItinerary(ctx, req)
}

func (i *ItineraryController) DeleteItinerary(ctx context.Context, req *models.DeleteItineraryReq) (*models.DeleteItineraryResp, error) {
	return i.serD.DeleteItinerary(ctx, req)
}

