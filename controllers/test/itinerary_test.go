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
func TestCreateItinerary_Validation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := mock.NewMockItineraryService(ctrl)
	validUUID := uuid.New().String()
	controller := &controllers.ItineraryController{
		Ser:       mockService,
		Validator: validator.New(),
	}
	tests := []struct {
		name    string
		req     *models.CreateItineraryReq
		wantErr error
	}{
		{
			name: "valid request",
			req: &models.CreateItineraryReq{
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
			req: &models.CreateItineraryReq{
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
				mockService.EXPECT().CreateItinerary(gomock.Any(), tt.req).Return(&models.CreateItineraryResp{}, nil).MaxTimes(100)
			} else {
				mockService.EXPECT().CreateItinerary(gomock.Any(), tt.req).Return(nil, custom_errs.ErrServiceError).MaxTimes(100)
			}
			_, err := controller.CreateItinerary(context.Background(), tt.req)
			if (err == nil && tt.wantErr != nil) || (err != nil && tt.wantErr == nil) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func TestGetItineraryById_Validation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := mock.NewMockItineraryService(ctrl)
	validUUID := uuid.New().String()
	controller := &controllers.ItineraryController{
		Ser:       mockService,
		Validator: validator.New(),
	}
	tests := []struct {
		name    string
		req     *models.GetItineraryByIdReq
		wantErr error
	}{
		{
			name: "valid request",
			req: &models.GetItineraryByIdReq{
				Id: validUUID,
			},
			wantErr: nil,
		},
		{
			name: "invalid request (empty Id)",
			req: &models.GetItineraryByIdReq{
				Id: "",
			},
			wantErr: custom_errs.ErrServiceError, 
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr == nil {
				mockService.EXPECT().GetItineraryById(gomock.Any(), tt.req).Return(&models.GetItineraryByIdResp{}, nil).MaxTimes(100)
			} else {
				mockService.EXPECT().GetItineraryById(gomock.Any(), tt.req).Return(nil, custom_errs.ErrServiceError).MaxTimes(100)
			}
			_, err := controller.GetItineraryById(context.Background(), tt.req)
			if (err == nil && tt.wantErr != nil) || (err != nil && tt.wantErr == nil) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func TestUpdateItinerary_Validation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := mock.NewMockItineraryService(ctrl)
	validUUID := uuid.New().String()
	controller := &controllers.ItineraryController{
		Ser:       mockService,
		Validator: validator.New(),
	}
	tests := []struct {
		name    string
		req     *models.UpdateItineraryReq
		wantErr error
	}{
		{
			name: "valid request",
			req: &models.UpdateItineraryReq{
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
			req: &models.UpdateItineraryReq{
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
				mockService.EXPECT().UpdateItinerary(gomock.Any(), tt.req).Return(&models.UpdateItineraryResp{}, nil).MaxTimes(100)
			} else {
				mockService.EXPECT().UpdateItinerary(gomock.Any(), tt.req).Return(nil, custom_errs.ErrServiceError).MaxTimes(100)
			}
			_, err := controller.UpdateItinerary(context.Background(), tt.req)
			if (err == nil && tt.wantErr != nil) || (err != nil && tt.wantErr == nil) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func TestDeleteItinerary_Validation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := mock.NewMockItineraryService(ctrl)
	validUUID := uuid.New().String()
	controller := &controllers.ItineraryController{
		Ser:       mockService,
		Validator: validator.New(),
	}
	tests := []struct {
		name    string
		req     *models.DeleteItineraryReq
		wantErr error
	}{
		{
			name: "valid request",
			req: &models.DeleteItineraryReq{
				Id: validUUID,
			},
			wantErr: nil,
		},
		{
			name: "invalid request (empty Id)",
			req: &models.DeleteItineraryReq{
				Id: "",
			},
			wantErr: custom_errs.ErrServiceError, 
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr == nil {
				mockService.EXPECT().DeleteItinerary(gomock.Any(), tt.req).Return(&models.DeleteItineraryResp{}, nil).MaxTimes(100)
			} else {
				mockService.EXPECT().DeleteItinerary(gomock.Any(), tt.req).Return(nil, custom_errs.ErrServiceError).MaxTimes(100)
			}
			_, err := controller.DeleteItinerary(context.Background(), tt.req)
			if (err == nil && tt.wantErr != nil) || (err != nil && tt.wantErr == nil) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}