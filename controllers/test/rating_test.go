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
)
func TestCreateRating_Validation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := mock.NewMockRatingService(ctrl)
	controller := &controllers.RatingController{
		Ser:       mockService,
		Validator: validator.New(),
	}
	tests := []struct {
		name    string
		req     *models.CreateRatingReq
		wantErr error
	}{
		{
			name: "valid request",
			req: &models.CreateRatingReq{
				Score: 1,
			},
			wantErr: nil,
		},
		{
			name: "invalid request (missing score)",
			req: &models.CreateRatingReq{
			},
			wantErr: custom_errs.ErrServiceError, 
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr == nil {
				mockService.EXPECT().CreateRating(gomock.Any(), tt.req).Return(&models.CreateRatingResp{}, nil).MaxTimes(100)
			} else {
				mockService.EXPECT().CreateRating(gomock.Any(), tt.req).Return(nil, custom_errs.ErrServiceError).MaxTimes(100)
			}
			_, err := controller.CreateRating(context.Background(), tt.req)
			if (err == nil && tt.wantErr != nil) || (err != nil && tt.wantErr == nil) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func TestGetRatingById_Validation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := mock.NewMockRatingService(ctrl)
	validUUID := uuid.New().String()
	controller := &controllers.RatingController{
		Ser:       mockService,
		Validator: validator.New(),
	}
	tests := []struct {
		name    string
		req     *models.GetRatingByIdReq
		wantErr error
	}{
		{
			name: "valid request",
			req: &models.GetRatingByIdReq{
				Id: validUUID,
			},
			wantErr: nil,
		},
		{
			name: "invalid request (empty Id)",
			req: &models.GetRatingByIdReq{
				Id: "",
			},
			wantErr: custom_errs.ErrServiceError, 
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr == nil {
				mockService.EXPECT().GetRatingById(gomock.Any(), tt.req).Return(&models.GetRatingByIdResp{}, nil).MaxTimes(100)
			} else {
				mockService.EXPECT().GetRatingById(gomock.Any(), tt.req).Return(nil, custom_errs.ErrServiceError).MaxTimes(100)
			}
			_, err := controller.GetRatingById(context.Background(), tt.req)
			if (err == nil && tt.wantErr != nil) || (err != nil && tt.wantErr == nil) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func TestUpdateRating_Validation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := mock.NewMockRatingService(ctrl)
	validUUID := uuid.New().String()
	controller := &controllers.RatingController{
		Ser:       mockService,
		Validator: validator.New(),
	}
	tests := []struct {
		name    string
		req     *models.UpdateRatingReq
		wantErr error
	}{
		{
			name: "valid request",
			req: &models.UpdateRatingReq{
				Id: validUUID,
				Score: 1,
			},
			wantErr: nil,
		},
		{
			name: "invalid request (empty Id)",
			req: &models.UpdateRatingReq{
				Id: "",
				Score: 1,
			},
			wantErr: custom_errs.ErrServiceError, 
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr == nil {
				mockService.EXPECT().UpdateRating(gomock.Any(), tt.req).Return(&models.UpdateRatingResp{}, nil).MaxTimes(100)
			} else {
				mockService.EXPECT().UpdateRating(gomock.Any(), tt.req).Return(nil, custom_errs.ErrServiceError).MaxTimes(100)
			}
			_, err := controller.UpdateRating(context.Background(), tt.req)
			if (err == nil && tt.wantErr != nil) || (err != nil && tt.wantErr == nil) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func TestDeleteRating_Validation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := mock.NewMockRatingService(ctrl)
	validUUID := uuid.New().String()
	controller := &controllers.RatingController{
		Ser:       mockService,
		Validator: validator.New(),
	}
	tests := []struct {
		name    string
		req     *models.DeleteRatingReq
		wantErr error
	}{
		{
			name: "valid request",
			req: &models.DeleteRatingReq{
				Id: validUUID,
			},
			wantErr: nil,
		},
		{
			name: "invalid request (empty Id)",
			req: &models.DeleteRatingReq{
				Id: "",
			},
			wantErr: custom_errs.ErrServiceError, 
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr == nil {
				mockService.EXPECT().DeleteRating(gomock.Any(), tt.req).Return(&models.DeleteRatingResp{}, nil).MaxTimes(100)
			} else {
				mockService.EXPECT().DeleteRating(gomock.Any(), tt.req).Return(nil, custom_errs.ErrServiceError).MaxTimes(100)
			}
			_, err := controller.DeleteRating(context.Background(), tt.req)
			if (err == nil && tt.wantErr != nil) || (err != nil && tt.wantErr == nil) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}