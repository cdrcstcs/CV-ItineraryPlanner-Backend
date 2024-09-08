package test
import (
	"context"
	"testing"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"itineraryplanner/dal/mock"
	"itineraryplanner/models"
	"itineraryplanner/service"
)
func ConfigUser(t *testing.T) (context.Context, *mock.MockUserDal, *service.UserService) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mock := mock.NewMockUserDal(ctrl)
	userService := &service.UserService{Dal: mock}
	return ctx, mock, userService
}
func TestCreateUser(t *testing.T) {
	ctx, mock, userService := ConfigUser(t)
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
				mock.EXPECT().CreateUser(
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
	ctx, mock, userService := ConfigUser(t)
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
				mock.EXPECT().GetUserById(
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
	ctx, mock, userService := ConfigUser(t)
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
				mock.EXPECT().GetUser(
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
	ctx, mock, userService := ConfigUser(t)
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
				mock.EXPECT().UpdateUser(
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
	ctx, mock, userService := ConfigUser(t)
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
				mock.EXPECT().DeleteUser(
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