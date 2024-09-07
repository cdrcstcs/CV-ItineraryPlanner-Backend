package controllers

import (
	"context"

	"itineraryplanner/controllers/inf"
	"itineraryplanner/models"
	service_inf "itineraryplanner/service/inf"
)

func NewTagController(ser service_inf.TagService) inf.TagController {
	return &TagController{
		ser: ser,
	}
}
type TagController struct {
	ser service_inf.TagService
}

func (t *TagController) CreateTag(ctx context.Context, req *models.CreateTagReq) (*models.CreateTagResp, error) {
	// we usually do request checking here at this layer, or can even do permission checking.
	// For now, I will leave it empty and just call service
	return t.ser.CreateTag(ctx, req)
}

func (t *TagController) GetTag(ctx context.Context, req *models.GetTagReq) (*models.GetTagResp, error) {
	return t.ser.GetTag(ctx, req)
}
func (t *TagController) GetTagById(ctx context.Context, req *models.GetTagByIdReq) (*models.GetTagByIdResp, error) {
	return t.ser.GetTagById(ctx, req)
}

func (t *TagController) UpdateTag(ctx context.Context, req *models.UpdateTagReq) (*models.UpdateTagResp, error) {
	// we usually do request checking here at this layer, or can even do permission checking.
	// For now, I will leave it empty and just call service
	return t.ser.UpdateTag(ctx, req)
}

func (t *TagController) DeleteTag(ctx context.Context, req *models.DeleteTagReq) (*models.DeleteTagResp, error) {
	return t.ser.DeleteTag(ctx, req)
}

