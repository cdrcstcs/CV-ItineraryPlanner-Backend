package controllers

import (
	"context"

	"itineraryplanner/controllers/inf"
	"itineraryplanner/models"
	service_inf "itineraryplanner/service/inf"
)

func NewRecommendItineraryController(serR service_inf.RecommendItineraryService) inf.RecommendItineraryController {
	return &AlgoController{
		serR: serR,
	}
}
func NewBuildItineraryController(serB service_inf.BuildItineraryService) inf.BuildItineraryController {
	return &AlgoController{
		serB: serB,
	}
}


type AlgoController struct {
	serR service_inf.RecommendItineraryService
	serB service_inf.BuildItineraryService
}

func (a *AlgoController) RecommendItinerary(ctx context.Context, req *models.RecommendItineraryReq) (*models.RecommendItineraryResp, error) {
	// we usually do request checking here at this layer, or can even do permission checking.
	// For now, I will leave it empty and just call service
	return a.serR.RecommendItinerary(ctx, req)
}

func (a *AlgoController) BuildItinerary(ctx context.Context, req *models.BuildItineraryReq) (*models.BuildItineraryResp, error) {
	// we usually do request checking here at this layer, or can even do permission checking.
	// For now, I will leave it empty and just call service
	return a.serB.BuildItinerary(ctx, req)
}
