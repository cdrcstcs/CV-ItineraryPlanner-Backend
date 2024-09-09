package controllers
import (
	"context"
	"errors"
	"itineraryplanner/controllers/inf"
	"itineraryplanner/models"
	service_inf "itineraryplanner/service/inf"
	"github.com/go-playground/validator/v10"
)
type UserController struct {
	Ser       service_inf.UserService
	Validator *validator.Validate
}
func NewUserController(ser service_inf.UserService) inf.UserController {
	return &UserController{
		Ser:       ser,
		Validator: validator.New(),
	}
}
func (u *UserController) validateRequest(req interface{}) error {
	err := u.Validator.Struct(req)
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
func (u *UserController) CreateUser(ctx context.Context, req *models.CreateUserReq) (*models.CreateUserResp, error) {
	if err := u.validateRequest(req); err != nil {
		return nil, err
	}
	return u.Ser.CreateUser(ctx, req)
}
func (u *UserController) GetUser(ctx context.Context, req *models.GetUserReq) (*models.GetUserResp, error) {
	return u.Ser.GetUser(ctx, req)
}
func (u *UserController) GetUserById(ctx context.Context, req *models.GetUserByIdReq) (*models.GetUserByIdResp, error) {
	if err := u.validateRequest(req); err != nil {
		return nil, err
	}
	return u.Ser.GetUserById(ctx, req)
}
func (u *UserController) UpdateUser(ctx context.Context, req *models.UpdateUserReq) (*models.UpdateUserResp, error) {
	if err := u.validateRequest(req); err != nil {
		return nil, err
	}
	return u.Ser.UpdateUser(ctx, req)
}
func (u *UserController) DeleteUser(ctx context.Context, req *models.DeleteUserReq) (*models.DeleteUserResp, error) {
	if err := u.validateRequest(req); err != nil {
		return nil, err
	}
	return u.Ser.DeleteUser(ctx, req)
}