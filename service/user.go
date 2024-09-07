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

func NewCreateUserService(cdal dal_inf.CreateUserDal) inf.CreateUserService {
	return &UserService{
		CDal: cdal,
	}
}
func NewGetUserByIdService(bdal dal_inf.GetUserByIdDal) inf.GetUserByIdService {
	return &UserService{
		BDal: bdal,
	}
}
func NewGetUserService(gdal dal_inf.GetUserDal) inf.GetUserService {
	return &UserService{
		GDal: gdal,
	}
}
func NewUpdateUserService(udal dal_inf.UpdateUserDal) inf.UpdateUserService {
	return &UserService{
		UDal: udal,
	}
}
func NewDeleteUserService(ddal dal_inf.DeleteUserDal) inf.DeleteUserService {
	return &UserService{
		DDal: ddal,
	}
}
func NewLoginUserService(ldal dal_inf.LoginUserDal) inf.LoginUserService {
	return &UserService{
		LDal: ldal,
	}
}
func NewUserDTOService() inf.UserDTOService {
	return &UserService{}
}
type UserService struct {
	CDal dal_inf.CreateUserDal
	BDal dal_inf.GetUserByIdDal
	GDal dal_inf.GetUserDal
	UDal dal_inf.UpdateUserDal
	DDal dal_inf.DeleteUserDal
	LDal dal_inf.LoginUserDal
}

func (a *UserService) CreateUser(ctx context.Context, req *models.CreateUserReq) (*models.CreateUserResp, error) {
	user := &models.User{}
	err := copier.Copy(user, req)
	if err != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}

	user, err = a.CDal.CreateUser(ctx, user)
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

func (a *UserService) LoginUser(ctx context.Context, req *models.LoginUserReq) (*models.LoginUserResp, error) {
	user := &models.User{}
	err := copier.Copy(user, req)
	if err != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}
	check, err := a.LDal.LoginUser(ctx, user)
	if err != nil {
		// TODO logging
		return nil, err
	}
	return &models.LoginUserResp{Check: check}, nil
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
	User, err := u.BDal.GetUserById(ctx, req.Id)
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}

	User1 := &models.User{}
	ok := copier.Copy(User1, User)
	if ok != nil {
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
	Users, err := u.GDal.GetUser(ctx)
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}

	User1 := []models.User{}
	ok := copier.Copy(User1, Users)
	if ok != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
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

	user, err = u.UDal.UpdateUser(ctx, user)
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
	user, err := u.DDal.DeleteUser(ctx, req.Id)
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}

	user1 := &models.User{}
	ok := copier.Copy(user1, user)
	if ok != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}
	dto, err := u.ConvertDBOToDTOUser(ctx, user1)
	if err != nil {
		// TODO logging
		return nil, err
	}
	return &models.DeleteUserResp{User: dto}, nil
}