package controllers

import (
	"context"

	"itineraryplanner/controllers/inf"
	"itineraryplanner/models"
	service_inf "itineraryplanner/service/inf"
)

func NewUserController(ser service_inf.UserService) inf.UserController {
	return &UserController{
		ser: ser,
	}
}
type UserController struct {
	ser service_inf.UserService
}
func (u *UserController) CreateUser(ctx context.Context, req *models.CreateUserReq) (*models.CreateUserResp, error) {
	// we usually do request checking here at this layer, or can even do permission checking.
	// For now, I will leave it empty and just call service
	return u.ser.CreateUser(ctx, req)
}

func (u *UserController) GetUser(ctx context.Context, req *models.GetUserReq) (*models.GetUserResp, error) {
	return u.ser.GetUser(ctx, req)
}
func (u *UserController) GetUserById(ctx context.Context, req *models.GetUserByIdReq) (*models.GetUserByIdResp, error) {
	return u.ser.GetUserById(ctx, req)
}

func (u *UserController) UpdateUser(ctx context.Context, req *models.UpdateUserReq) (*models.UpdateUserResp, error) {
	// we usually do request checking here at this layer, or can even do permission checking.
	// For now, I will leave it empty and just call service
	return u.ser.UpdateUser(ctx, req)
}

func (u *UserController) DeleteUser(ctx context.Context, req *models.DeleteUserReq) (*models.DeleteUserResp, error) {
	return u.ser.DeleteUser(ctx, req)
}
