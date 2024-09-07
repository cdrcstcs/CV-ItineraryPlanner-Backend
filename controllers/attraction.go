package controllers

import (
	"context"

	"itineraryplanner/controllers/inf"
	"itineraryplanner/models"
	service_inf "itineraryplanner/service/inf"
)

func NewCreateAttractionController(serC service_inf.CreateAttractionService) inf.CreateAttractionController {
	return &AttractionController{
		serC: serC,
	}
}
func NewGetAttractionController(serG service_inf.GetAttractionService) inf.GetAttractionController {
	return &AttractionController{
		serG: serG,
	}
}
func NewGetAttractionByIdController(serB service_inf.GetAttractionByIdService) inf.GetAttractionByIdController {
	return &AttractionController{
		serB: serB,
	}
}
func NewUpdateAttractionController(serU service_inf.UpdateAttractionService) inf.UpdateAttractionController {
	return &AttractionController{
		serU: serU,
	}
}
func NewDeleteAttractionController(serD service_inf.DeleteAttractionService) inf.DeleteAttractionController {
	return &AttractionController{
		serD: serD,
	}
}


type AttractionController struct {
	serC service_inf.CreateAttractionService
	serB service_inf.GetAttractionByIdService
	serG service_inf.GetAttractionService
	serU service_inf.UpdateAttractionService
	serD service_inf.DeleteAttractionService
}

func (a *AttractionController) CreateAttraction(ctx context.Context, req *models.CreateAttractionReq) (*models.CreateAttractionResp, error) {
	// we usually do request checking here at this layer, or can even do permission checking.
	// For now, I will leave it empty and just call service
	return a.serC.CreateAttraction(ctx, req)
}

func (a *AttractionController) GetAttraction(ctx context.Context, req *models.GetAttractionReq) (*models.GetAttractionResp, error) {
	return a.serG.GetAttraction(ctx, req)
}
func (a *AttractionController) GetAttractionById(ctx context.Context, req *models.GetAttractionByIdReq) (*models.GetAttractionByIdResp, error) {
	return a.serB.GetAttractionById(ctx, req)
}

func (a *AttractionController) UpdateAttraction(ctx context.Context, req *models.UpdateAttractionReq) (*models.UpdateAttractionResp, error) {
	// we usually do request checking here at this layer, or can even do permission checking.
	// For now, I will leave it empty and just call service
	return a.serU.UpdateAttraction(ctx, req)
}

func (a *AttractionController) DeleteAttraction(ctx context.Context, req *models.DeleteAttractionReq) (*models.DeleteAttractionResp, error) {
	return a.serD.DeleteAttraction(ctx, req)
}

