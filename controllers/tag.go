package controllers
import (
	"context"
	"errors"
	"itineraryplanner/controllers/inf"
	"itineraryplanner/models"
	service_inf "itineraryplanner/service/inf"
	"github.com/go-playground/validator/v10"
)
type TagController struct {
	Ser       service_inf.TagService
	Validator *validator.Validate
}
func NewTagController(ser service_inf.TagService) inf.TagController {
	return &TagController{
		Ser:       ser,
		Validator: validator.New(),
	}
}
func (t *TagController) validateRequest(req interface{}) error {
	err := t.Validator.Struct(req)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}
		var errMsg string
		for _, err := range err.(validator.ValidationErrors) {
			errMsg += err.Field() + " " + err.Tag() + ", "
		}
		if len(errMsg) > 0 {
			return errors.New(errMsg)
		}
	}
	return nil
}
func (t *TagController) CreateTag(ctx context.Context, req *models.CreateTagReq) (*models.CreateTagResp, error) {
	if err := t.validateRequest(req); err != nil {
		return nil, err
	}
	return t.Ser.CreateTag(ctx, req)
}
func (t *TagController) GetTag(ctx context.Context, req *models.GetTagReq) (*models.GetTagResp, error) {
	return t.Ser.GetTag(ctx, req)
}
func (t *TagController) GetTagById(ctx context.Context, req *models.GetTagByIdReq) (*models.GetTagByIdResp, error) {
	if err := t.validateRequest(req); err != nil {
		return nil, err
	}
	return t.Ser.GetTagById(ctx, req)
}
func (t *TagController) UpdateTag(ctx context.Context, req *models.UpdateTagReq) (*models.UpdateTagResp, error) {
	if err := t.validateRequest(req); err != nil {
		return nil, err
	}
	return t.Ser.UpdateTag(ctx, req)
}
func (t *TagController) DeleteTag(ctx context.Context, req *models.DeleteTagReq) (*models.DeleteTagResp, error) {
	if err := t.validateRequest(req); err != nil {
		return nil, err
	}
	return t.Ser.DeleteTag(ctx, req)
}