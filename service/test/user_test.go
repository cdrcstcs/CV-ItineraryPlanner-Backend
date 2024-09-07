package test

import (
	"context"
	"testing"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"itineraryplanner/dal/mock"
	"itineraryplanner/models"
	"itineraryplanner/service"
	ser_mock "itineraryplanner/service/mock"
)

func ConfigCreateUser(t *testing.T) (context.Context, *mock.MockCreateUserDal, *ser_mock.MockUserDTOService, *service.UserService) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockC := mock.NewMockCreateUserDal(ctrl)
	mockDTO := ser_mock.NewMockUserDTOService(ctrl)
	userService := &service.UserService{CDal: mockC}
	return ctx, mockC, mockDTO, userService
}

func ConfigGetUser(t *testing.T) (context.Context, *mock.MockGetUserDal, *ser_mock.MockUserDTOService, *service.UserService) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockG := mock.NewMockGetUserDal(ctrl)
	mockDTO := ser_mock.NewMockUserDTOService(ctrl)
	userService := &service.UserService{GDal: mockG}
	return ctx, mockG, mockDTO, userService
}

func ConfigGetUserById(t *testing.T) (context.Context, *mock.MockGetUserByIdDal, *ser_mock.MockUserDTOService, *service.UserService) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockB := mock.NewMockGetUserByIdDal(ctrl)
	mockDTO := ser_mock.NewMockUserDTOService(ctrl)
	userService := &service.UserService{BDal: mockB}
	return ctx, mockB, mockDTO, userService
}

func ConfigUpdateUser(t *testing.T) (context.Context, *mock.MockUpdateUserDal, *ser_mock.MockUserDTOService, *service.UserService){
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockU := mock.NewMockUpdateUserDal(ctrl)
	mockDTO := ser_mock.NewMockUserDTOService(ctrl)
	userService := &service.UserService{UDal: mockU}
	return ctx, mockU, mockDTO, userService
}
func ConfigDeleteUser(t *testing.T) (context.Context, *mock.MockDeleteUserDal, *ser_mock.MockUserDTOService, *service.UserService){
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockD := mock.NewMockDeleteUserDal(ctrl)
	mockDTO := ser_mock.NewMockUserDTOService(ctrl)
	userService := &service.UserService{DDal: mockD}
	return ctx, mockD, mockDTO, userService
}
func ConfigUserDTOService(t *testing.T) (context.Context, *service.UserService) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockC := mock.NewMockCreateUserDal(ctrl)
	mockG := mock.NewMockGetUserDal(ctrl)
	tagDTOService := &service.UserService{GDal: mockG, CDal: mockC}
	return ctx, tagDTOService
}

func TestCreateUser(t *testing.T) {
	ctx, mockC, mockDTO, userService := ConfigCreateUser(t)

	type arg struct {
		req *models.CreateUserReq
		ctx context.Context
	}

	tests := []struct {
		name     string
		arg      arg
		before   func(t *testing.T)
		wantResp *models.CreateUserResp
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				req: &models.CreateUserReq{
					Name:     "test_name",
					Email:    "test_email@example.com",
					Password: "test_password",
				},
			},
			before: func(t *testing.T) {
				mockC.EXPECT().CreateUser(
					ctx,
					gomock.Any(),
				).Return(
					&models.User{
						Id:       "test_user_id",
						Name:     "test_name",
						Email:    "test_email@example.com",
						Password: "test_password",
					},
					nil,
				)
				mockDTO.EXPECT().ConvertDBOToDTOUser(
					ctx,
					&models.User{
						Id:       "test_user_id",
						Name:     "test_name",
						Email:    "test_email@example.com",
						Password: "test_password",
					},
				).Return(
					&models.UserDTO{
						Id:       "test_user_id",
						Name:     "test_name",
						Email:    "test_email@example.com",
						Password: "test_password",
					},
					nil,
				).Do(func(ctx context.Context, user *models.User) {
					fmt.Println("ConvertDBOToDTOUser called with", user)
				})
			},
			wantResp: &models.CreateUserResp{
				User: &models.UserDTO{
					Id:       "test_user_id",
					Name:     "test_name",
					Email:    "test_email@example.com",
					Password: "test_password",
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := userService.CreateUser(ctx, tt.arg.req)
			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestGetUserById(t *testing.T) {
	ctx, mockB, mockDTO, userService := ConfigGetUserById(t)

	type arg struct {
		req *models.GetUserByIdReq
		ctx context.Context
	}

	tests := []struct {
		name     string
		arg      arg
		before   func(t *testing.T)
		wantResp *models.GetUserByIdResp
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				req: &models.GetUserByIdReq{
					Id: "test_user_id",
				},
			},
			before: func(t *testing.T) {
				mockB.EXPECT().GetUserById(
					ctx,
					"test_user_id",
				).Return(
					&models.User{
						Id:       "test_user_id",
						Name:     "test_name",
						Email:    "test_email@example.com",
						Password: "test_password",
					},
					nil,
				)
				mockDTO.EXPECT().ConvertDBOToDTOUser(
					ctx,
					gomock.Any(),
				).Return(
					&models.UserDTO{
						Id:       "test_user_id",
						Name:     "test_name",
						Email:    "test_email@example.com",
						Password: "test_password",
					},
					nil,
				).Do(func(ctx context.Context, user *models.User) {
					fmt.Println("ConvertDBOToDTOUser called with", user)
				})
			},
			wantResp: &models.GetUserByIdResp{
				User: &models.UserDTO{
					Id:       "test_user_id",
					Name:     "test_name",
					Email:    "test_email@example.com",
					Password: "test_password",
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := userService.GetUserById(ctx, tt.arg.req)
			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestGetUser(t *testing.T) {
	ctx, mockG, mockDTO, userService := ConfigGetUser(t)

	type arg struct {
		req *models.GetUserReq
		ctx context.Context
	}

	tests := []struct {
		name     string
		arg      arg
		before   func(t *testing.T)
		wantResp *models.GetUserResp
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				req: &models.GetUserReq{
				},
			},
			before: func(t *testing.T) {
				mockG.EXPECT().GetUser(
					ctx,
				).Return(
					[]*models.User{
						{
							Id:       "test_user_id",
							Name:     "test_name",
							Email:    "test_email@example.com",
							Password: "test_password",
						},
					},
					nil,
				)
				mockDTO.EXPECT().ConvertDBOToDTOUser(
					ctx,
					&models.User{
						Id:       "test_user_id",
						Name:     "test_name",
						Email:    "test_email@example.com",
						Password: "test_password",
					},
				).Return(
					&models.UserDTO{
						Id:       "test_user_id",
						Name:     "test_name",
						Email:    "test_email@example.com",
						Password: "test_password",
					},
					nil,
				).Do(func(ctx context.Context, user *models.User) {
					fmt.Println("ConvertDBOToDTOUser called with", user)
				})
			},
			wantResp: &models.GetUserResp{
				Users: []*models.UserDTO{
					{
						Id:       "test_user_id",
						Name:     "test_name",
						Email:    "test_email@example.com",
						Password: "test_password",
					},
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := userService.GetUser(ctx, tt.arg.req)
			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestUpdateUser(t *testing.T) {
	ctx, mockU, mockDTO, userService := ConfigUpdateUser(t)

	type arg struct {
		req *models.UpdateUserReq
		ctx context.Context
	}

	tests := []struct {
		name     string
		arg      arg
		before   func(t *testing.T)
		wantResp *models.UpdateUserResp
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				req: &models.UpdateUserReq{
					Name:     "test_name",
					Email:    "test_email@example.com",
					Password: "test_password",
				},
			},
			before: func(t *testing.T) {
				mockU.EXPECT().UpdateUser(
					ctx,
					gomock.Any(),
				).Return(
					&models.User{
						Id:       "test_user_id",
						Name:     "test_name",
						Email:    "test_email@example.com",
						Password: "test_password",
					},
					nil,
				)
				mockDTO.EXPECT().ConvertDBOToDTOUser(
					ctx,
					&models.User{
						Id:       "test_user_id",
						Name:     "test_name",
						Email:    "test_email@example.com",
						Password: "test_password",
					},
				).Return(
					&models.UserDTO{
						Id:       "test_user_id",
						Name:     "test_name",
						Email:    "test_email@example.com",
						Password: "test_password",
					},
					nil,
				).Do(func(ctx context.Context, user *models.User) {
					fmt.Println("ConvertDBOToDTOUser called with", user)
				})
			},
			wantResp: &models.UpdateUserResp{
				User: &models.UserDTO{
					Id:       "test_user_id",
					Name:     "test_name",
					Email:    "test_email@example.com",
					Password: "test_password",
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := userService.UpdateUser(ctx, tt.arg.req)
			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestDeleteUser(t *testing.T) {
	ctx, mockD, mockDTO, userService := ConfigDeleteUser(t)

	type arg struct {
		req *models.DeleteUserReq
		ctx context.Context
	}

	tests := []struct {
		name     string
		arg      arg
		before   func(t *testing.T)
		wantResp *models.DeleteUserResp
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				req: &models.DeleteUserReq{
					Id: "test_user_id",
				},
			},
			before: func(t *testing.T) {
				mockD.EXPECT().DeleteUser(
					ctx,
					"test_user_id",
				).Return(
					&models.User{
						Id:       "test_user_id",
						Name:     "test_name",
						Email:    "test_email@example.com",
						Password: "test_password",
					},
					nil,
				)
				mockDTO.EXPECT().ConvertDBOToDTOUser(
					ctx,
					&models.User{
						Id:       "test_user_id",
						Name:     "test_name",
						Email:    "test_email@example.com",
						Password: "test_password",
					},
				).Return(
					&models.UserDTO{
						Id:       "test_user_id",
						Name:     "test_name",
						Email:    "test_email@example.com",
						Password: "test_password",
					},
					nil,
				).Do(func(ctx context.Context, user *models.User) {
					fmt.Println("ConvertDBOToDTOUser called with", user)
				})
			},
			wantResp: &models.DeleteUserResp{
				User: &models.UserDTO{
					Id:       "test_user_id",
					Name:     "test_name",
					Email:    "test_email@example.com",
					Password: "test_password",
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := userService.DeleteUser(ctx, tt.arg.req)
			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestConvertDBOToDTOUser(t *testing.T) {
	ctx, userService := ConfigUserDTOService(t)

	type arg struct {
		user *models.User
		ctx context.Context
	}

	tests := []struct {
		name     string
		before   func(t *testing.T)
		arg      arg
		wantResp *models.UserDTO
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				user: &models.User{
					Id:       "test_user_id",
					Name:     "test_name",
					Email:    "test_email@example.com",
					Password: "test_password",
				},
				ctx: ctx,
			},
			wantResp: &models.UserDTO{
				Id:       "test_user_id",
				Name:     "test_name",
				Email:    "test_email@example.com",
				Password: "test_password",
			},
			wantErr: nil,
		},
		// Add more test cases for error scenarios, if needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := userService.ConvertDBOToDTOUser(tt.arg.ctx, tt.arg.user)
			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
