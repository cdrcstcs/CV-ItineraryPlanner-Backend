package controllers

import (
	"context"

	"itineraryplanner/controllers/inf"
	"itineraryplanner/models"
	service_inf "itineraryplanner/service/inf"
)

func NewCreateCoordinateController(serC service_inf.CreateCoordinateService) inf.CreateCoordinateController {
	return &CoordinateController{
		serC: serC,
	}
}
func NewGetCoordinateController(serG service_inf.GetCoordinateService) inf.GetCoordinateController {
	return &CoordinateController{
		serG: serG,
	}
}
func NewGetCoordinateByIdController(serB service_inf.GetCoordinateByIdService) inf.GetCoordinateByIdController {
	return &CoordinateController{
		serB: serB,
	}
}
func NewUpdateCoordinateController(serU service_inf.UpdateCoordinateService) inf.UpdateCoordinateController {
	return &CoordinateController{
		serU: serU,
	}
}
func NewDeleteCoordinateController(serD service_inf.DeleteCoordinateService) inf.DeleteCoordinateController {
	return &CoordinateController{
		serD: serD,
	}
}


type CoordinateController struct {
	serC service_inf.CreateCoordinateService
	serB service_inf.GetCoordinateByIdService
	serG service_inf.GetCoordinateService
	serU service_inf.UpdateCoordinateService
	serD service_inf.DeleteCoordinateService
}

func (c *CoordinateController) CreateCoordinate(ctx context.Context, req *models.CreateCoordinateReq) (*models.CreateCoordinateResp, error) {
	// we usually do request checking here at this layer, or can even do permission checking.
	// For now, I will leave it empty and just call service
	return c.serC.CreateCoordinate(ctx, req)
}

func (c *CoordinateController) GetCoordinate(ctx context.Context, req *models.GetCoordinateReq) (*models.GetCoordinateResp, error) {
	return c.serG.GetCoordinate(ctx, req)
}
func (c *CoordinateController) GetCoordinateById(ctx context.Context, req *models.GetCoordinateByIdReq) (*models.GetCoordinateByIdResp, error) {
	return c.serB.GetCoordinateById(ctx, req)
}

func (c *CoordinateController) UpdateCoordinate(ctx context.Context, req *models.UpdateCoordinateReq) (*models.UpdateCoordinateResp, error) {
	// we usually do request checking here at this layer, or can even do permission checking.
	// For now, I will leave it empty and just call service
	return c.serU.UpdateCoordinate(ctx, req)
}

func (c *CoordinateController) DeleteCoordinate(ctx context.Context, req *models.DeleteCoordinateReq) (*models.DeleteCoordinateResp, error) {
	return c.serD.DeleteCoordinate(ctx, req)
}

