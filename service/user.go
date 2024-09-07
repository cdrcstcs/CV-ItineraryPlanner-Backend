package service

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"itineraryplanner/common/custom_errs"
	dal_inf "itineraryplanner/dal/inf"
	"itineraryplanner/models"
	"itineraryplanner/service/inf"
)

func NewUserService(dal dal_inf.UserDal) inf.UserService {
	return &UserService{
		Dal: dal,
	}
}
type UserService struct {
	Dal dal_inf.UserDal
}

func (a *UserService) CreateUser(ctx context.Context, req *models.CreateUserReq) (*models.CreateUserResp, error) {
	user := &models.User{}
	err := copier.Copy(user, req)
	if err != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}

	user, err = a.Dal.CreateUser(ctx, user)
	if err != nil {
		// TODO logging
		return nil, err
	}

	dto, err := a.ConvertDBOToDTOUser(ctx, user)
	if err != nil {
		// TODO logging
		return nil, err
	}

	return &models.CreateUserResp{User: dto}, nil
}

func (u *UserService) ConvertDBOToDTOUser(ctx context.Context, use *models.User) (*models.UserDTO, error) {
	if use == nil {
		return nil, custom_errs.ServerError
	}
	user := &models.UserDTO{}
	err:= copier.Copy(user, use)
	if err != nil {
		// TODO logging
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}
	return user, nil
}

func (u *UserService) GetUserById(ctx context.Context, req *models.GetUserByIdReq) (*models.GetUserByIdResp, error) {
	User, err := u.Dal.GetUserById(ctx, req.Id)
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}

	User1 := &models.User{}
	err = copier.Copy(User1, User)
	if err != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}
	dto, err := u.ConvertDBOToDTOUser(ctx, User1)
	if err != nil {
		// TODO logging
		return nil, err
	}
	return &models.GetUserByIdResp{User: dto}, nil
}

func (u *UserService) GetUser(ctx context.Context, req *models.GetUserReq) (*models.GetUserResp, error) {
	Users, err := u.Dal.GetUser(ctx)
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}

	User1 := []models.User{}
	ok := copier.Copy(User1, Users)
	if ok != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", ok.Error())
		return nil, errors.Wrap(custom_errs.ServerError, ok.Error())
	}
	dtos := make([]*models.UserDTO, 0)
	for _, v := range User1 {
		dto, err := u.ConvertDBOToDTOUser(ctx, &v)
		if err != nil {
			// TODO logging
			return nil, err
		}
		dtos = append(dtos, dto)
	}
	return &models.GetUserResp{Users: dtos}, nil
}
func (u *UserService) UpdateUser(ctx context.Context, req *models.UpdateUserReq) (*models.UpdateUserResp, error) {
	user := &models.User{}
	err := copier.Copy(user, req)
	if err != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}

	user, err = u.Dal.UpdateUser(ctx, user)
	if err != nil {
		// TODO logging
		return nil, err
	}
	dto, err := u.ConvertDBOToDTOUser(ctx, user)
	if err != nil {
		// TODO logging
		return nil, err
	}

	return &models.UpdateUserResp{User: dto}, nil
}

func (u *UserService) DeleteUser(ctx context.Context, req *models.DeleteUserReq) (*models.DeleteUserResp, error) {
	user, err := u.Dal.DeleteUser(ctx, req.Id)
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}

	user1 := &models.User{}
	ok := copier.Copy(user1, user)
	if ok != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", ok.Error())
		return nil, errors.Wrap(custom_errs.ServerError, ok.Error())
	}
	dto, err := u.ConvertDBOToDTOUser(ctx, user1)
	if err != nil {
		// TODO logging
		return nil, err
	}
	return &models.DeleteUserResp{User: dto}, nil
}