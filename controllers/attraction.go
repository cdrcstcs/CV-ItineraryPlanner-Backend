package controllers
import (
	"context"
	"errors"
	"itineraryplanner/controllers/inf"
	"itineraryplanner/models"
	service_inf "itineraryplanner/service/inf"
	"github.com/go-playground/validator/v10"
)
type AttractionController struct {
	Ser       service_inf.AttractionService
	Validator *validator.Validate
}
func NewAttractionController(ser service_inf.AttractionService) inf.AttractionController {
	return &AttractionController{
		Ser:       ser,
		Validator: validator.New(),
	}
}
func (a *AttractionController) validateRequest(req interface{}) error {
	err := a.Validator.Struct(req)
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
func (a *AttractionController) CreateAttraction(ctx context.Context, req *models.CreateAttractionReq) (*models.CreateAttractionResp, error) {
	if err := a.validateRequest(req); err != nil {
		return nil, err
	}
	return a.Ser.CreateAttraction(ctx, req)
}
func (a *AttractionController) GetAttraction(ctx context.Context, req *models.GetAttractionReq) (*models.GetAttractionResp, error) {
	return a.Ser.GetAttraction(ctx, req)
}
func (a *AttractionController) GetAttractionById(ctx context.Context, req *models.GetAttractionByIdReq) (*models.GetAttractionByIdResp, error) {
	if err := a.validateRequest(req); err != nil {
		return nil, err
	}
	return a.Ser.GetAttractionById(ctx, req)
}
func (a *AttractionController) UpdateAttraction(ctx context.Context, req *models.UpdateAttractionReq) (*models.UpdateAttractionResp, error) {
	if err := a.validateRequest(req); err != nil {
		return nil, err
	}
	return a.Ser.UpdateAttraction(ctx, req)
}
func (a *AttractionController) DeleteAttraction(ctx context.Context, req *models.DeleteAttractionReq) (*models.DeleteAttractionResp, error) {
	if err := a.validateRequest(req); err != nil {
		return nil, err
	}
	return a.Ser.DeleteAttraction(ctx, req)
}