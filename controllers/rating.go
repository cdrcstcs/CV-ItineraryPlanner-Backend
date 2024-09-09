package controllers
import (
	"context"
	"errors"
	"itineraryplanner/controllers/inf"
	"itineraryplanner/models"
	service_inf "itineraryplanner/service/inf"
	"github.com/go-playground/validator/v10"
)
type RatingController struct {
	Ser       service_inf.RatingService
	Validator *validator.Validate
}
func NewRatingController(ser service_inf.RatingService) inf.RatingController {
	return &RatingController{
		Ser:       ser,
		Validator: validator.New(),
	}
}
func (r *RatingController) validateRequest(req interface{}) error {
	err := r.Validator.Struct(req)
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
func (r *RatingController) CreateRating(ctx context.Context, req *models.CreateRatingReq) (*models.CreateRatingResp, error) {
	if err := r.validateRequest(req); err != nil {
		return nil, err
	}
	return r.Ser.CreateRating(ctx, req)
}
func (r *RatingController) GetRating(ctx context.Context, req *models.GetRatingReq) (*models.GetRatingResp, error) {
	return r.Ser.GetRating(ctx, req)
}
func (r *RatingController) GetRatingById(ctx context.Context, req *models.GetRatingByIdReq) (*models.GetRatingByIdResp, error) {
	if err := r.validateRequest(req); err != nil {
		return nil, err
	}
	return r.Ser.GetRatingById(ctx, req)
}
func (r *RatingController) UpdateRating(ctx context.Context, req *models.UpdateRatingReq) (*models.UpdateRatingResp, error) {
	if err := r.validateRequest(req); err != nil {
		return nil, err
	}
	return r.Ser.UpdateRating(ctx, req)
}
func (r *RatingController) DeleteRating(ctx context.Context, req *models.DeleteRatingReq) (*models.DeleteRatingResp, error) {
	if err := r.validateRequest(req); err != nil {
		return nil, err
	}
	return r.Ser.DeleteRating(ctx, req)
}