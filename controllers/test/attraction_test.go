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
func TestCreateAttraction_Validation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := mock.NewMockAttractionService(ctrl)
	validUUID := uuid.New().String()
	controller := &controllers.AttractionController{
		Ser:       mockService,
		Validator: validator.New(),
	}
	tests := []struct {
		name    string
		req     *models.CreateAttractionReq
		wantErr error
	}{
		{
			name: "valid request",
			req: &models.CreateAttractionReq{
				Name:       "test",
				Address:    "test",
				X:          1,
				Y:          1,
				TagIDs:     []string{validUUID, validUUID},
				RatingId:   validUUID,
				City:       "test",
			},
			wantErr: nil,
		},
		{
			name: "invalid request (missing Name)",
			req: &models.CreateAttractionReq{
				Address:    "test",
				X:          1,
				Y:          1,
				TagIDs:     []string{validUUID, validUUID},
				RatingId:   validUUID,
				City:       "test",
			},
			wantErr: custom_errs.ErrServiceError, 
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr == nil {
				mockService.EXPECT().CreateAttraction(gomock.Any(), tt.req).Return(&models.CreateAttractionResp{}, nil).MaxTimes(100)
			} else {
				mockService.EXPECT().CreateAttraction(gomock.Any(), tt.req).Return(nil, custom_errs.ErrServiceError).MaxTimes(100)
			}
			_, err := controller.CreateAttraction(context.Background(), tt.req)
			if (err == nil && tt.wantErr != nil) || (err != nil && tt.wantErr == nil) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func TestGetAttractionById_Validation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := mock.NewMockAttractionService(ctrl)
	validUUID := uuid.New().String()
	controller := &controllers.AttractionController{
		Ser:       mockService,
		Validator: validator.New(),
	}
	tests := []struct {
		name    string
		req     *models.GetAttractionByIdReq
		wantErr error
	}{
		{
			name: "valid request",
			req: &models.GetAttractionByIdReq{
				Id: validUUID,
			},
			wantErr: nil,
		},
		{
			name: "invalid request (empty Id)",
			req: &models.GetAttractionByIdReq{
				Id: "",
			},
			wantErr: custom_errs.ErrServiceError, 
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr == nil {
				mockService.EXPECT().GetAttractionById(gomock.Any(), tt.req).Return(&models.GetAttractionByIdResp{}, nil).MaxTimes(100)
			} else {
				mockService.EXPECT().GetAttractionById(gomock.Any(), tt.req).Return(nil, custom_errs.ErrServiceError).MaxTimes(100)
			}
			_, err := controller.GetAttractionById(context.Background(), tt.req)
			if (err == nil && tt.wantErr != nil) || (err != nil && tt.wantErr == nil) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func TestUpdateAttraction_Validation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := mock.NewMockAttractionService(ctrl)
	validUUID := uuid.New().String()
	controller := &controllers.AttractionController{
		Ser:       mockService,
		Validator: validator.New(),
	}
	tests := []struct {
		name    string
		req     *models.UpdateAttractionReq
		wantErr error
	}{
		{
			name: "valid request",
			req: &models.UpdateAttractionReq{
				Id: validUUID,
				Name: "test",
				Address: "test",
				X: 1,
				Y: 1,
				TagIDs: []string{validUUID, validUUID},
				RatingId: validUUID,
				City: "test",
			},
			wantErr: nil,
		},
		{
			name: "invalid request (empty Id)",
			req: &models.UpdateAttractionReq{
				Id: "",
				Name: "test",
				Address: "test",
				X: 1,
				Y: 1,
				TagIDs: []string{validUUID, validUUID},
				RatingId: validUUID,
				City: "test",
			},
			wantErr: custom_errs.ErrServiceError, 
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr == nil {
				mockService.EXPECT().UpdateAttraction(gomock.Any(), tt.req).Return(&models.UpdateAttractionResp{}, nil).MaxTimes(100)
			} else {
				mockService.EXPECT().UpdateAttraction(gomock.Any(), tt.req).Return(nil, custom_errs.ErrServiceError).MaxTimes(100)
			}
			_, err := controller.UpdateAttraction(context.Background(), tt.req)
			if (err == nil && tt.wantErr != nil) || (err != nil && tt.wantErr == nil) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func TestDeleteAttraction_Validation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := mock.NewMockAttractionService(ctrl)
	validUUID := uuid.New().String()
	controller := &controllers.AttractionController{
		Ser:       mockService,
		Validator: validator.New(),
	}
	tests := []struct {
		name    string
		req     *models.DeleteAttractionReq
		wantErr error
	}{
		{
			name: "valid request",
			req: &models.DeleteAttractionReq{
				Id: validUUID,
			},
			wantErr: nil,
		},
		{
			name: "invalid request (empty Id)",
			req: &models.DeleteAttractionReq{
				Id: "",
			},
			wantErr: custom_errs.ErrServiceError, 
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr == nil {
				mockService.EXPECT().DeleteAttraction(gomock.Any(), tt.req).Return(&models.DeleteAttractionResp{}, nil).MaxTimes(100)
			} else {
				mockService.EXPECT().DeleteAttraction(gomock.Any(), tt.req).Return(nil, custom_errs.ErrServiceError).MaxTimes(100)
			}
			_, err := controller.DeleteAttraction(context.Background(), tt.req)
			if (err == nil && tt.wantErr != nil) || (err != nil && tt.wantErr == nil) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}