package controllers

import (
	"context"

	"itineraryplanner/controllers/inf"
	"itineraryplanner/models"
	service_inf "itineraryplanner/service/inf"
)

func NewAttractionController(ser service_inf.AttractionService) inf.AttractionController {
	return &AttractionController{
		ser: ser,
	}
}

type AttractionController struct {
	ser service_inf.AttractionService
}

func (a *AttractionController) CreateAttraction(ctx context.Context, req *models.CreateAttractionReq) (*models.CreateAttractionResp, error) {
	// we usually do request checking here at this layer, or can even do permission checking.
	// For now, I will leave it empty and just call service
	return a.ser.CreateAttraction(ctx, req)
}

func (a *AttractionController) GetAttraction(ctx context.Context, req *models.GetAttractionReq) (*models.GetAttractionResp, error) {
	return a.ser.GetAttraction(ctx, req)
}
func (a *AttractionController) GetAttractionById(ctx context.Context, req *models.GetAttractionByIdReq) (*models.GetAttractionByIdResp, error) {
	return a.ser.GetAttractionById(ctx, req)
}

func (a *AttractionController) UpdateAttraction(ctx context.Context, req *models.UpdateAttractionReq) (*models.UpdateAttractionResp, error) {
	// we usually do request checking here at this layer, or can even do permission checking.
	// For now, I will leave it empty and just call service
	return a.ser.UpdateAttraction(ctx, req)
}

func (a *AttractionController) DeleteAttraction(ctx context.Context, req *models.DeleteAttractionReq) (*models.DeleteAttractionResp, error) {
	return a.ser.DeleteAttraction(ctx, req)
}

