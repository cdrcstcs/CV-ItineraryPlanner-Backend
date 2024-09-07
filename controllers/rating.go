package controllers

import (
	"context"

	"itineraryplanner/controllers/inf"
	"itineraryplanner/models"
	service_inf "itineraryplanner/service/inf"
)

func NewRatingController(ser service_inf.RatingService) inf.RatingController {
	return &RatingController{
		ser: ser,
	}
}


type RatingController struct {
	ser service_inf.RatingService
}

func (r *RatingController) CreateRating(ctx context.Context, req *models.CreateRatingReq) (*models.CreateRatingResp, error) {
	// we usually do request checking here at this layer, or can even do permission checking.
	// For now, I will leave it empty and just call service
	return r.ser.CreateRating(ctx, req)
}

func (r *RatingController) GetRating(ctx context.Context, req *models.GetRatingReq) (*models.GetRatingResp, error) {
	return r.ser.GetRating(ctx, req)
}
func (r *RatingController) GetRatingById(ctx context.Context, req *models.GetRatingByIdReq) (*models.GetRatingByIdResp, error) {
	return r.ser.GetRatingById(ctx, req)
}

func (r *RatingController) UpdateRating(ctx context.Context, req *models.UpdateRatingReq) (*models.UpdateRatingResp, error) {
	// we usually do request checking here at this layer, or can even do permission checking.
	// For now, I will leave it empty and just call service
	return r.ser.UpdateRating(ctx, req)
}

func (r *RatingController) DeleteRating(ctx context.Context, req *models.DeleteRatingReq) (*models.DeleteRatingResp, error) {
	return r.ser.DeleteRating(ctx, req)
}

