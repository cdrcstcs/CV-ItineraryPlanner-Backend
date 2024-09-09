package test
import (
	"context"
	"itineraryplanner/common/custom_errs"
	"itineraryplanner/controllers"
	"itineraryplanner/models"
	"github.com/google/uuid" 
	"itineraryplanner/service/mock"
	"testing"
	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"time"
)
func TestCreateUser_Validation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := mock.NewMockUserService(ctrl)
	controller := &controllers.UserController{
		Ser:       mockService,
		Validator: validator.New(),
	}
	tests := []struct {
		name    string
		req     *models.CreateUserReq
		wantErr error
	}{
		{
			name: "valid request",
			req: &models.CreateUserReq{
				Name: "test", 
				Password: "test",    
				Email: "test@gmail.com",       
				EmailPassword: "test",
				Phone: "test",         
				Token: "test",         
				UserType: "USER",    
				RefreshToken: "test",
				CreatedAt: time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),   
				UpdatedAt: time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),    
			},
			wantErr: nil,
		},
		{
			name: "invalid request (missing token)",
			req: &models.CreateUserReq{
				Name: "test", 
				Password: "test",    
				Email: "test@gmail.com",       
				EmailPassword: "test",
				Phone: "test",         
				UserType: "USER",    
				RefreshToken: "test",
				CreatedAt: time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),   
				UpdatedAt: time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
			},
			wantErr: custom_errs.ErrServiceError, 
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr == nil {
				mockService.EXPECT().CreateUser(gomock.Any(), tt.req).Return(&models.CreateUserResp{}, nil).MaxTimes(100)
			} else {
				mockService.EXPECT().CreateUser(gomock.Any(), tt.req).Return(nil, custom_errs.ErrServiceError).MaxTimes(100)
			}
			_, err := controller.CreateUser(context.Background(), tt.req)
			if (err == nil && tt.wantErr != nil) || (err != nil && tt.wantErr == nil) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func TestGetUserById_Validation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := mock.NewMockUserService(ctrl)
	validUUID := uuid.New().String()
	controller := &controllers.UserController{
		Ser:       mockService,
		Validator: validator.New(),
	}
	tests := []struct {
		name    string
		req     *models.GetUserByIdReq
		wantErr error
	}{
		{
			name: "valid request",
			req: &models.GetUserByIdReq{
				Id: validUUID,
			},
			wantErr: nil,
		},
		{
			name: "invalid request (empty Id)",
			req: &models.GetUserByIdReq{
				Id: "",
			},
			wantErr: custom_errs.ErrServiceError, 
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr == nil {
				mockService.EXPECT().GetUserById(gomock.Any(), tt.req).Return(&models.GetUserByIdResp{}, nil).MaxTimes(100)
			} else {
				mockService.EXPECT().GetUserById(gomock.Any(), tt.req).Return(nil, custom_errs.ErrServiceError).MaxTimes(100)
			}
			_, err := controller.GetUserById(context.Background(), tt.req)
			if (err == nil && tt.wantErr != nil) || (err != nil && tt.wantErr == nil) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func TestUpdateUser_Validation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := mock.NewMockUserService(ctrl)
	validUUID := uuid.New().String()
	controller := &controllers.UserController{
		Ser:       mockService,
		Validator: validator.New(),
	}
	tests := []struct {
		name    string
		req     *models.UpdateUserReq
		wantErr error
	}{
		{
			name: "valid request",
			req: &models.UpdateUserReq{
				Id: validUUID,
				Name: "test", 
				Password: "test",    
				Email: "test@gmail.com",       
				EmailPassword: "test",
				Phone: "test",         
				Token: "test",         
				UserType: "USER",    
				RefreshToken: "test",
				CreatedAt: time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),   
				UpdatedAt: time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
			},
			wantErr: nil,
		},
		{
			name: "invalid request (empty Id)",
			req: &models.UpdateUserReq{
				Id: "",
				Name: "test", 
				Password: "test",    
				Email: "test@gmail.com",       
				EmailPassword: "test",
				Phone: "test",         
				Token: "test",         
				UserType: "USER",    
				RefreshToken: "test",
				CreatedAt: time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),   
				UpdatedAt: time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
			},
			wantErr: custom_errs.ErrServiceError, 
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr == nil {
				mockService.EXPECT().UpdateUser(gomock.Any(), tt.req).Return(&models.UpdateUserResp{}, nil).MaxTimes(100)
			} else {
				mockService.EXPECT().UpdateUser(gomock.Any(), tt.req).Return(nil, custom_errs.ErrServiceError).MaxTimes(100)
			}
			_, err := controller.UpdateUser(context.Background(), tt.req)
			if (err == nil && tt.wantErr != nil) || (err != nil && tt.wantErr == nil) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func TestDeleteUser_Validation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := mock.NewMockUserService(ctrl)
	validUUID := uuid.New().String()
	controller := &controllers.UserController{
		Ser:       mockService,
		Validator: validator.New(),
	}
	tests := []struct {
		name    string
		req     *models.DeleteUserReq
		wantErr error
	}{
		{
			name: "valid request",
			req: &models.DeleteUserReq{
				Id: validUUID,
			},
			wantErr: nil,
		},
		{
			name: "invalid request (empty Id)",
			req: &models.DeleteUserReq{
				Id: "",
			},
			wantErr: custom_errs.ErrServiceError, 
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr == nil {
				mockService.EXPECT().DeleteUser(gomock.Any(), tt.req).Return(&models.DeleteUserResp{}, nil).MaxTimes(100)
			} else {
				mockService.EXPECT().DeleteUser(gomock.Any(), tt.req).Return(nil, custom_errs.ErrServiceError).MaxTimes(100)
			}
			_, err := controller.DeleteUser(context.Background(), tt.req)
			if (err == nil && tt.wantErr != nil) || (err != nil && tt.wantErr == nil) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}