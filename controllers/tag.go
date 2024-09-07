package controllers

import (
	"context"

	"itineraryplanner/controllers/inf"
	"itineraryplanner/models"
	service_inf "itineraryplanner/service/inf"
)

func NewCreateTagController(serC service_inf.CreateTagService) inf.CreateTagController {
	return &TagController{
		serC: serC,
	}
}
func NewGetTagController(serG service_inf.GetTagService) inf.GetTagController {
	return &TagController{
		serG: serG,
	}
}
func NewGetTagByIdController(serB service_inf.GetTagByIdService) inf.GetTagByIdController {
	return &TagController{
		serB: serB,
	}
}
func NewUpdateTagController(serU service_inf.UpdateTagService) inf.UpdateTagController {
	return &TagController{
		serU: serU,
	}
}
func NewDeleteTagController(serD service_inf.DeleteTagService) inf.DeleteTagController {
	return &TagController{
		serD: serD,
	}
}


type TagController struct {
	serC service_inf.CreateTagService
	serB service_inf.GetTagByIdService
	serG service_inf.GetTagService
	serU service_inf.UpdateTagService
	serD service_inf.DeleteTagService
}

func (t *TagController) CreateTag(ctx context.Context, req *models.CreateTagReq) (*models.CreateTagResp, error) {
	// we usually do request checking here at this layer, or can even do permission checking.
	// For now, I will leave it empty and just call service
	return t.serC.CreateTag(ctx, req)
}

func (t *TagController) GetTag(ctx context.Context, req *models.GetTagReq) (*models.GetTagResp, error) {
	return t.serG.GetTag(ctx, req)
}
func (t *TagController) GetTagById(ctx context.Context, req *models.GetTagByIdReq) (*models.GetTagByIdResp, error) {
	return t.serB.GetTagById(ctx, req)
}

func (t *TagController) UpdateTag(ctx context.Context, req *models.UpdateTagReq) (*models.UpdateTagResp, error) {
	// we usually do request checking here at this layer, or can even do permission checking.
	// For now, I will leave it empty and just call service
	return t.serU.UpdateTag(ctx, req)
}

func (t *TagController) DeleteTag(ctx context.Context, req *models.DeleteTagReq) (*models.DeleteTagResp, error) {
	return t.serD.DeleteTag(ctx, req)
}

