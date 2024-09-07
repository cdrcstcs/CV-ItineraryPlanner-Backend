package controllers

import (
	"context"

	"itineraryplanner/controllers/inf"
	"itineraryplanner/models"
	service_inf "itineraryplanner/service/inf"
)

func NewCreateUserController(serC service_inf.CreateUserService) inf.CreateUserController {
	return &UserController{
		serC: serC,
	}
}
func NewGetUserController(serG service_inf.GetUserService) inf.GetUserController {
	return &UserController{
		serG: serG,
	}
}
func NewGetUserByIdController(serB service_inf.GetUserByIdService) inf.GetUserByIdController {
	return &UserController{
		serB: serB,
	}
}
func NewUpdateUserController(serU service_inf.UpdateUserService) inf.UpdateUserController {
	return &UserController{
		serU: serU,
	}
}
func NewDeleteUserController(serD service_inf.DeleteUserService) inf.DeleteUserController {
	return &UserController{
		serD: serD,
	}
}
func NewLoginUserController(serL service_inf.LoginUserService) inf.LoginUserController {
	return &UserController{
		serL: serL,
	}
}


type UserController struct {
	serC service_inf.CreateUserService
	serB service_inf.GetUserByIdService
	serG service_inf.GetUserService
	serU service_inf.UpdateUserService
	serD service_inf.DeleteUserService
	serL service_inf.LoginUserService
}

func (u *UserController) CreateUser(ctx context.Context, req *models.CreateUserReq) (*models.CreateUserResp, error) {
	// we usually do request checking here at this layer, or can even do permission checking.
	// For now, I will leave it empty and just call service
	return u.serC.CreateUser(ctx, req)
}

func (u *UserController) GetUser(ctx context.Context, req *models.GetUserReq) (*models.GetUserResp, error) {
	return u.serG.GetUser(ctx, req)
}
func (u *UserController) GetUserById(ctx context.Context, req *models.GetUserByIdReq) (*models.GetUserByIdResp, error) {
	return u.serB.GetUserById(ctx, req)
}

func (u *UserController) UpdateUser(ctx context.Context, req *models.UpdateUserReq) (*models.UpdateUserResp, error) {
	// we usually do request checking here at this layer, or can even do permission checking.
	// For now, I will leave it empty and just call service
	return u.serU.UpdateUser(ctx, req)
}

func (u *UserController) DeleteUser(ctx context.Context, req *models.DeleteUserReq) (*models.DeleteUserResp, error) {
	return u.serD.DeleteUser(ctx, req)
}

func (u *UserController) LoginUser(ctx context.Context, req *models.LoginUserReq) (*models.LoginUserResp, error) {
	return u.serL.LoginUser(ctx, req)
}

