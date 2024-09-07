package controllers

import (
	"context"

	"itineraryplanner/controllers/inf"
	"itineraryplanner/models"
	service_inf "itineraryplanner/service/inf"
)

func NewCreateRatingController(serC service_inf.CreateRatingService) inf.CreateRatingController {
	return &RatingController{
		serC: serC,
	}
}
func NewGetRatingController(serG service_inf.GetRatingService) inf.GetRatingController {
	return &RatingController{
		serG: serG,
	}
}
func NewGetRatingByIdController(serB service_inf.GetRatingByIdService) inf.GetRatingByIdController {
	return &RatingController{
		serB: serB,
	}
}
func NewUpdateRatingController(serU service_inf.UpdateRatingService) inf.UpdateRatingController {
	return &RatingController{
		serU: serU,
	}
}
func NewDeleteRatingController(serD service_inf.DeleteRatingService) inf.DeleteRatingController {
	return &RatingController{
		serD: serD,
	}
}


type RatingController struct {
	serC service_inf.CreateRatingService
	serB service_inf.GetRatingByIdService
	serG service_inf.GetRatingService
	serU service_inf.UpdateRatingService
	serD service_inf.DeleteRatingService
}

func (r *RatingController) CreateRating(ctx context.Context, req *models.CreateRatingReq) (*models.CreateRatingResp, error) {
	// we usually do request checking here at this layer, or can even do permission checking.
	// For now, I will leave it empty and just call service
	return r.serC.CreateRating(ctx, req)
}

func (r *RatingController) GetRating(ctx context.Context, req *models.GetRatingReq) (*models.GetRatingResp, error) {
	return r.serG.GetRating(ctx, req)
}
func (r *RatingController) GetRatingById(ctx context.Context, req *models.GetRatingByIdReq) (*models.GetRatingByIdResp, error) {
	return r.serB.GetRatingById(ctx, req)
}

func (r *RatingController) UpdateRating(ctx context.Context, req *models.UpdateRatingReq) (*models.UpdateRatingResp, error) {
	// we usually do request checking here at this layer, or can even do permission checking.
	// For now, I will leave it empty and just call service
	return r.serU.UpdateRating(ctx, req)
}

func (r *RatingController) DeleteRating(ctx context.Context, req *models.DeleteRatingReq) (*models.DeleteRatingResp, error) {
	return r.serD.DeleteRating(ctx, req)
}

